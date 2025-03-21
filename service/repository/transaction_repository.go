package repository

import (
	"github.com/saipulmuiz/krplus/models"
	api "github.com/saipulmuiz/krplus/service"
	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) api.TransactionRepository {
	return &transactionRepo{db}
}

func (u *transactionRepo) GetTransactionByID(transactionID int64) (transaction *models.Transaction, err error) {
	return transaction, u.db.Where("id = ?", transactionID).First(&transaction).Error
}

func (r *transactionRepo) GetTransactionsByUserID(userId int64) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id", userId).Find(&transactions).Error
	return &transactions, err
}

func (r *transactionRepo) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}
