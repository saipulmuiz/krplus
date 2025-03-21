package service

import (
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
)

type UserUsecase interface {
	Register(request *models.RegisterUser) (user *models.User, errx serror.SError)
	Login(request *models.LoginUser) (res models.LoginResponse, errx serror.SError)
}

type CreditUsecase interface {
	CreateCreditLimit(req models.CreateLimitRequest) (errx serror.SError)
	GetCredits(req models.CreditLimitRequest) (res []models.GetCreditResponse, totalData int64, errx serror.SError)
}

type TransactionUsecase interface {
	RecordTransaction(req models.RecordTransactionRequest) (errx serror.SError)
}

type PaymentUsecase interface {
	CreatePayment(req models.PaymentRequest) (errx serror.SError)
}
