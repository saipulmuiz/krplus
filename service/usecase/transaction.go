package usecase

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
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
		Tenor:  1,
	})
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][RecordTransaction] Error retrieving credit")
		return
	}

	credit := (*credits)[0]

	if req.OTR > credit.RemainingLimitAmount {
		errx = serror.Newi(http.StatusBadRequest, "Insufficient credit limit")
		return
	}

	// Record transaction
	transaction := models.Transaction{
		ContractNumber:    req.ContractNumber,
		UserID:            user.UserID,
		OTR:               req.OTR,
		AdminFee:          req.AdminFee,
		InstallmentAmount: req.Installment,
		Interest:          req.Interest,
		AssetName:         req.AssetName,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err = u.transactionRepo.CreateTransaction(&transaction)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][RecordTransaction] Error creating transaction")
		return
	}

	// Update credit limit
	credit.RemainingLimitAmount -= req.OTR
	credit.UsedLimitAmount += req.OTR
	credit.UpdatedAt = time.Now()
	err = u.creditRepo.UpdateCredit(credit.CreditID, credit)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][RecordTransaction] Error updating credit limit")
		return
	}

	return
}
