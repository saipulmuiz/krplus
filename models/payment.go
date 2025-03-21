package models

import "time"

type Payment struct {
	ID            int64 `gorm:"primaryKey"`
	TransactionID int64 `gorm:"not null"`
	PaymentAmount float64
	PaymentDate   string
	CreatedBy     string
	CreatedAt     time.Time `json:"created_at"`
	UpdatedBy     string
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentRequest struct {
	TransactionID int64   `json:"transaction_id"`
	PaymentAmount float64 `json:"payment_amount"`
	PaymentDate   string  `json:"payment_date"`
	CreatedBy     string
}
