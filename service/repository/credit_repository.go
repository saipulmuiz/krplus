package repository

import (
	"github.com/saipulmuiz/krplus/models"
	api "github.com/saipulmuiz/krplus/service"
	"gorm.io/gorm"
)

type creditRepo struct {
	db *gorm.DB
}

func NewCreditRepo(db *gorm.DB) api.CreditRepository {
	return &creditRepo{db}
}

func (r *creditRepo) GetCredits(req models.CreditLimitRequest) (*[]models.CreditLimit, int64, error) {
	var (
		credits []models.CreditLimit
		count   int64
	)

	if req.Limit == 0 || req.Page == 0 {
		req.Limit = 10
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit

	query := r.db.Table("credit_limits").
		Offset(offset).
		Limit(req.Limit).
		Order("created_at DESC")

	if req.UserID != 0 {
		query = query.Where("user_id = ?", req.UserID)
	}

	if req.Tenor != 0 {
		query = query.Where("tenor = ?", req.Tenor)
	}

	err := query.Find(&credits).Error
	if err != nil {
		return nil, count, err
	}

	err = query.Count(&count).Error

	return &credits, count, err
}

func (u *creditRepo) GetCreditByID(creditId int64) (credit *models.CreditLimit, err error) {
	return credit, u.db.Where("credit_id", creditId).Take(&credit).Error
}

func (r *creditRepo) UpdateCredit(creditId int64, updatedData models.CreditLimit) error {
	return r.db.Model(&models.CreditLimit{}).Where("credit_id = ?", creditId).Updates(updatedData).Error
}

func (r *creditRepo) CreateCredits(credits []models.CreditLimit) error {
	return r.db.Create(&credits).Error
}
