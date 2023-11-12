package book

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/entity"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateBookRequest struct {
	Id          int64  `json:"id,string" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *Handler) UpdateBookController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &UpdateBookRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	book := model.NewBook(params.Id, params.Title, params.Description, userId.(int64))
	err := h.service.UpdateBook(ctx, book)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, entity.ConvertToBookVO(book))
}
