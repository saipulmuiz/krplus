package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/pkg/serror"
)

func (h *Handler) RecordTransaction(ctx *gin.Context) {
	var (
		errx serror.SError
		req  models.RecordTransactionRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, serror.NewFromError(err))
		return
	}

	errx = h.transactionUsecase.RecordTransaction(req)
	if errx != nil {
		handleError(ctx, errx.Code(), errx)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Transaction recorded successfully",
	})
}
