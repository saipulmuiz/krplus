package usecase

import (
	"net/http"
	"sync"
	"time"

	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
	"github.com/saipulmuiz/krplus/pkg/utils/utfloat"
	api "github.com/saipulmuiz/krplus/service"
)

type TransactionUsecase struct {
	transactionRepo api.TransactionRepository
	creditRepo      api.CreditRepository
	userRepo        api.UserRepository
}

func NewTransactionUsecase(
	transactionRepo api.TransactionRepository,
	creditRepo api.CreditRepository,
	userRepo api.UserRepository,
) api.TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: transactionRepo,
		creditRepo:      creditRepo,
		userRepo:        userRepo,
	}
}

func (u *TransactionUsecase) RecordTransaction(req models.RecordTransactionRequest) (errx serror.SError) {
	var err error

	// Validate user existence
	user, err := u.userRepo.GetUserByNIK(req.NIK)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][RecordTransaction] Error retrieving user")
		return
	}

	if user.UserID == 0 {
		errx = serror.Newi(http.StatusNotFound, "User not found")
		return
	}

	credits, _, err := u.creditRepo.GetCredits(models.CreditLimitRequest{
		UserID: user.UserID,
		Tenor:  req.Tenor,
	})
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][RecordTransaction] Error retrieving credit")
		return
	}

	if len(*credits) == 0 {
		errx = serror.Newi(http.StatusNotFound, "Credit not found")
		return
	}

	credit := (*credits)[0]

	if req.OTR > credit.RemainingLimitAmount {
		errx = serror.Newi(http.StatusBadRequest, "Insufficient credit limit")
		return
	}

	// calculate installment
	req.Installment = (req.OTR + req.AdminFee) + (req.OTR+req.Interest)/float64(req.Tenor)
	if req.Installment <= 0 {
		errx = serror.Newi(http.StatusBadRequest, "Invalid installment amount")
		return
	}

	var wg sync.WaitGroup
	var transactionErr, creditErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		transaction := models.Transaction{
			ContractNumber:    req.ContractNumber,
			UserID:            user.UserID,
			OTR:               req.OTR,
			Tenor:             req.Tenor,
			AdminFee:          req.AdminFee,
			InstallmentAmount: utfloat.Round(req.Installment, 2),
			Interest:          req.Interest,
			AssetName:         req.AssetName,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}

		transactionErr = u.transactionRepo.CreateTransaction(&transaction)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		credit.RemainingLimitAmount -= req.OTR
		credit.UsedLimitAmount += req.OTR
		credit.UpdatedAt = time.Now()
		creditErr = u.creditRepo.UpdateCredit(credit.CreditID, credit)
	}()

	wg.Wait()

	if transactionErr != nil {
		errx = serror.NewFromError(transactionErr)
		errx.AddComments("[usecase][RecordTransaction] Error creating transaction")
		return
	}

	if creditErr != nil {
		errx = serror.NewFromError(creditErr)
		errx.AddComments("[usecase][RecordTransaction] Error updating credit limit")
		return
	}

	return
}
