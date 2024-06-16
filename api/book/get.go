package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBookRequest struct {
	Id int64 `json:"id,string" form:"id" binding:"required"`
}

func (h *Handler) GetBookController(ctx *gin.Context) {
	var params GetBookRequest
	if err := ctx.Bind(&params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}
	book, err := h.service.GetBook(ctx, params.Id)
	if err != nil {
		res.RespondWithError(ctx, http.StatusNotFound, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}
	res.RespondSuccess(ctx, book)
}
