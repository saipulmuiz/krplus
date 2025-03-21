package usecase

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
	api "github.com/saipulmuiz/krplus/service"
	"gorm.io/gorm"
)

type PaymentUsecase struct {
	paymentRepo     api.PaymentRepository
	transactionRepo api.TransactionRepository
	creditRepo      api.CreditRepository
	userRepo        api.UserRepository
}

func NewPaymentUsecase(
	paymentRepo api.PaymentRepository,
	transactionRepo api.TransactionRepository,
	creditRepo api.CreditRepository,
	userRepo api.UserRepository,
) api.PaymentUsecase {
	return &PaymentUsecase{
		paymentRepo:     paymentRepo,
		transactionRepo: transactionRepo,
		creditRepo:      creditRepo,
		userRepo:        userRepo,
	}
}

func (u *PaymentUsecase) CreatePayment(req models.PaymentRequest) (errx serror.SError) {
	// Validate payment request
	if req.PaymentAmount <= 0 {
		return serror.Newi(http.StatusBadRequest, "invalid payment amount")
	}

	if req.TransactionID == 0 {
		return serror.Newi(http.StatusBadRequest, "transaction ID is required")
	}

	// check if transaction exists
	transaction, err := u.transactionRepo.GetTransactionByID(req.TransactionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errx = serror.Newi(http.StatusNotFound, "Transaction not found")
			return
		}

		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreatePayment] Error retrieving transaction")
		return
	}

	payments, err := u.paymentRepo.GetPaymentsByTransactionID(req.TransactionID)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreatePayment] Error retrieving payment")
		return
	}
	if len(*payments) > transaction.Tenor {
		errx = serror.Newi(http.StatusBadRequest, "Transaction already paid off")
		return
	}

	if req.PaymentAmount != transaction.InstallmentAmount {
		errx = serror.Newi(http.StatusBadRequest, "Payment amount is not equal to the installment amount")
		return
	}

	// Check if user exists
	user, err := u.userRepo.GetUserByID(transaction.UserID)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreatePayment] Error retrieving user")
		return
	}
	if user == nil {
		errx = serror.Newi(http.StatusNotFound, "User not found")
		return
	}

	payment := &models.Payment{
		TransactionID: req.TransactionID,
		PaymentAmount: req.PaymentAmount,
		PaymentDate:   req.PaymentDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     user.Email,
		UpdatedBy:     user.Email,
	}

	err = u.paymentRepo.CreatePayment(payment)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreatePayment] Error creating payment record")
		return
	}

	return nil
}
