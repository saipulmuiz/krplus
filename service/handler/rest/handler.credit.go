package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
	"github.com/saipulmuiz/krplus/pkg/utils/utint"
)

func (h *Handler) CreateCreditLimit(ctx *gin.Context) {
	var (
		errx serror.SError
		req  models.CreateLimitRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, serror.NewFromError(err))
		return
	}

	errx = h.creditUsecase.CreateCreditLimit(req)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Credit limit successfully created",
	})
}

func (h *Handler) GetCredits(ctx *gin.Context) {
	var (
		errx serror.SError
	)

	page := utint.StringToInt(ctx.Query("page"), 1)
	limit := utint.StringToInt(ctx.Query("limit"), 10)
	userID := utint.StringToInt(ctx.Query("user_id"), 0)

	data, totalData, errx := h.creditUsecase.GetCredits(models.CreditLimitRequest{
		Page:   int(page),
		Limit:  int(limit),
		UserID: userID,
	})
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Get credits successfully",
		Data:    data,
		Meta: map[string]interface{}{
			"total_data": totalData,
		},
	})
}
