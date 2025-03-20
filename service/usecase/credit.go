package usecase

import (
	"net/http"
	"time"

	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
	api "github.com/saipulmuiz/krplus/service"
	"github.com/saipulmuiz/krplus/service/helper"
)

type CreditUsecase struct {
	creditRepo api.CreditRepository
	userRepo   api.UserRepository
}

func NewCreditUsecase(
	creditRepo api.CreditRepository,
	userRepo api.UserRepository,
) api.CreditUsecase {
	return &CreditUsecase{
		creditRepo: creditRepo,
		userRepo:   userRepo,
	}
}

func (u *CreditUsecase) CreateCreditLimit(req models.CreateLimitRequest) (errx serror.SError) {
	var err error

	// Validate user existence
	user, err := u.userRepo.GetUserByID(req.UserID)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreateCreditLimit] Error retrieving user")
		return
	}

	if user.UserID == 0 {
		errx = serror.Newi(http.StatusNotFound, "User not found")
		return
	}

	// Check if credit already exists for the user
	existingCredits, _, err := u.creditRepo.GetCredits(models.CreditLimitRequest{
		UserID: req.UserID,
	})
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreateCreditLimit] Error checking existing credits")
		return
	}

	if len(*existingCredits) > 0 {
		errx = serror.Newi(http.StatusBadRequest, "Credit limit already exists for this user")
		return
	}

	var defaultCreditLimit = []models.TenorLimit{
		{Tenor: 1, LimitAmount: 1000000},
		{Tenor: 2, LimitAmount: 1200000},
		{Tenor: 3, LimitAmount: 1500000},
		{Tenor: 6, LimitAmount: 2000000},
	}

	var credits []models.CreditLimit
	for _, limit := range defaultCreditLimit {
		credit := models.CreditLimit{
			UserID:               req.UserID,
			Tenor:                limit.Tenor,
			InitialLimitAmount:   limit.LimitAmount,
			UsedLimitAmount:      0,
			RemainingLimitAmount: limit.LimitAmount,
			CreatedBy:            user.FullName,
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
			UpdatedBy:            user.FullName,
		}

		credits = append(credits, credit)
	}

	err = u.creditRepo.CreateCredits(credits)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][CreateCreditLimit] Error creating credit limit")
		return
	}

	return
}

func (u *CreditUsecase) GetCredits(req models.CreditLimitRequest) (res []models.GetCreditResponse, totalData int64, errx serror.SError) {
	var (
		err     error
		credits *[]models.CreditLimit
	)

	credits, totalData, err = u.creditRepo.GetCredits(req)
	if err != nil {
		errx = serror.NewFromError(err)
		errx.AddComments("[usecase][GetCredits] Error retrieving credits")
		return
	}

	var groupedCreditsMap = make(map[int64]*models.GetCreditResponse)
	for _, credit := range *credits {
		if _, exists := groupedCreditsMap[credit.UserID]; !exists {
			user, err := u.userRepo.GetUserByID(credit.UserID)
			if err != nil {
				errx = serror.NewFromError(err)
				errx.AddComments("[usecase][GetCredits] Error retrieving user")
				return
			}

			groupedCreditsMap[credit.UserID] = &models.GetCreditResponse{
				CreditID:  credit.CreditID,
				UserName:  user.FullName,
				Limits:    []models.TenorLimit{},
				CreatedAt: helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, credit.CreatedAt),
				UpdatedAt: helper.ParseDateTime(helper.DATE_FORMAT_YYYY_MM_DD_TIME, credit.UpdatedAt),
			}
		}
		groupedCreditsMap[credit.UserID].Limits = append(groupedCreditsMap[credit.UserID].Limits, models.TenorLimit{
			Tenor:       credit.Tenor,
			LimitAmount: credit.InitialLimitAmount,
		})
	}

	for _, value := range groupedCreditsMap {
		res = append(res, *value)
	}

	return
}
