package repository

import (
	"github.com/saipulmuiz/krplus/models"
	api "github.com/saipulmuiz/krplus/service"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) api.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Register(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepo) GetUserByID(userID int64) (user *models.User, err error) {
	return user, u.db.Where("user_id = ?", userID).First(&user).Error
}

func (u *userRepo) GetUserByEmail(email string) (user *models.User, err error) {
	return user, u.db.Where("email = ?", email).First(&user).Error
}
