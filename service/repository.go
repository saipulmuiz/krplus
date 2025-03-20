package service

import (
	"github.com/saipulmuiz/krplus/models"
)

type UserRepository interface {
	Register(user *models.User) (*models.User, error)
	GetUserByID(userID int64) (user *models.User, err error)
	GetUserByNIK(nik string) (user *models.User, err error)
	GetUserByEmail(email string) (user *models.User, err error)
}

type CreditRepository interface {
	CreateCredits(credits []models.CreditLimit) error
	GetCredits(req models.CreditLimitRequest) (*[]models.CreditLimit, int64, error)
	GetCreditByID(creditId int64) (credit *models.CreditLimit, err error)
	UpdateCredit(creditId int64, updatedData models.CreditLimit) error
}

type TransactionRepository interface {
	GetTransactionsByUserID(userId int64) (*[]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
}

type PaymentRepository interface {
	GetPaymentsByTransactionID(transactionId int64) (*[]models.Payment, error)
	RecordPayment(payment *models.Payment) error
}
