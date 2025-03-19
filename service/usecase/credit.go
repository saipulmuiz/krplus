package usecase

import (
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
