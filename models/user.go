package models

import (
	"time"

	"github.com/saipulmuiz/krplus/service/helper"
	"gorm.io/gorm"
)

type User struct {
	UserID      int64  `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	FullName    string `gorm:"not null"`
	LegalName   string
	Email       string `gorm:"not null;" json:"email"`
	Password    string `gorm:"not null;" json:"-"`
	NIK         string `gorm:"unique;not null"`
	BirthPlace  string
	BirthDate   string
	Salary      float64
	KTPPhoto    string
	SelfiePhoto string
	CreatedBy   string
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   string
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterUser struct {
	FullName    string  `json:"full_name" validate:"required"`
	LegalName   string  `json:"legal_name"`
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=6"`
	NIK         string  `json:"nik" validate:"required"`
	BirthPlace  string  `json:"birth_place"`
	BirthDate   string  `json:"birth_date"`
	Salary      float64 `json:"salary"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserInfo struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.HashPassword(u.Password)
	return
}
