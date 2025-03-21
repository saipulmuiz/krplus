package models

import "time"

type Transaction struct {
	ID                int64  `gorm:"primaryKey"`
	ContractNumber    string `gorm:"unique;not null"`
	UserID            int64  `gorm:"not null"`
	OTR               float64
	Tenor             int
	AdminFee          float64
	InstallmentAmount float64
	Interest          float64
	AssetName         string
	CreatedBy         string
	CreatedAt         time.Time `json:"created_at"`
	UpdatedBy         string
	UpdatedAt         time.Time `json:"updated_at"`
}

type RecordTransactionRequest struct {
	ContractNumber string  `json:"contract_number"`
	NIK            string  `json:"nik"`
	OTR            float64 `json:"otr"`
	AdminFee       float64 `json:"admin_fee"`
	Installment    float64 `json:"installment"`
	Interest       float64 `json:"interest"`
	AssetName      string  `json:"asset_name"`
	Tenor          int     `json:"tenor"`
}
