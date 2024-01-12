package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
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

func (h *Handler) DeleteBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
	}

	params := &DeleteBookRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	if err := h.service.DeleteBook(ctx, params.Id, userId.(int64)); err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}

func (h *Handler) RemoveCardFromBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &DeleteCardFromBookRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	err := h.service.DeleteCardFromBook(ctx, params.BookId, params.CardId, userId.(int64))
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}
