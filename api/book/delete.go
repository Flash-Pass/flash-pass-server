package book

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type DeleteBookRequest struct {
	Id int64 `json:"id,string" binding:"required"`
}

type DeleteCardFromBookRequest struct {
	BookId int64 `json:"bookId,string" binding:"required"`
	CardId int64 `json:"cardId,string" binding:"required"`
}

func (h *Handler) DeleteBookController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
	}

	params := &DeleteBookRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	if err := h.service.DeleteBook(ctx, params.Id, userId.(int64)); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}

func (h *Handler) RemoveCardFromBookController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &DeleteCardFromBookRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	err := h.service.DeleteCardFromBook(ctx, params.BookId, params.CardId, userId.(int64))
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
