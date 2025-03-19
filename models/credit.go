package models

import "time"

type CreditLimit struct {
	CreditID             int64 `gorm:"primaryKey"`
	UserID               int64 `gorm:"not null"`
	Tenor                int   `gorm:"not null"`
	InitialLimitAmount   float64
	UsedLimitAmount      float64
	RemainingLimitAmount float64
	CreatedBy            string
	CreatedAt            time.Time `json:"created_at"`
	UpdatedBy            string
	UpdatedAt            time.Time `json:"updated_at"`
}

type CreditLimitRequest struct {
	UserID int64 `json:"user_id"`
	Tenor  int   `json:"tenor"`
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
}

type GetCreditResponse struct {
	CreditID  int64        `json:"credit_id"`
	UserName  string       `json:"user_name"`
	Limits    []TenorLimit `json:"limits"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

type TenorLimit struct {
	Tenor       int     `json:"tenor"`
	LimitAmount float64 `json:"limit_amount"`
}
