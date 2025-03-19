package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
	"github.com/saipulmuiz/krplus/service/helper"
)

func (h *Handler) Register(ctx *gin.Context) {
	var (
		request models.RegisterUser
		errx    serror.SError
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errx = serror.NewFromErrori(http.StatusBadRequest, err)
		errx.AddComments("[handler][Register] while BodyJSONBind")
		handleError(ctx, errx.Code(), errx)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		validationMessages := helper.BuildAndGetValidationMessage(err)
		handleValidationError(ctx, validationMessages)

		return
	}

	_, errx = h.userUsecase.Register(&request)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Message: "User has successfully to registered",
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var (
		request models.LoginUser
		errx    serror.SError
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		errx = serror.NewFromErrori(http.StatusBadRequest, err)
		errx.AddComments("[handler][login] while BodyJSONBind")
		handleError(ctx, errx.Code(), errx)
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		validationMessages := helper.BuildAndGetValidationMessage(err)
		handleValidationError(ctx, validationMessages)

		return
	}

	res, errx := h.userUsecase.Login(&request)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "User has successfully to login",
		Data:    res,
	})
}
