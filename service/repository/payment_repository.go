package repository

import (
	"github.com/saipulmuiz/krplus/models"
	api "github.com/saipulmuiz/krplus/service"
	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) api.PaymentRepository {
	return &paymentRepo{db}
}

func (r *paymentRepo) GetPaymentsByTransactionID(transactionId int64) (*[]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("transaction_id", transactionId).Find(&payments).Error
	return &payments, err
}

func (r *paymentRepo) RecordPayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}
