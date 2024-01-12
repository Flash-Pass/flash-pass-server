package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/entity"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type UpdateBookRequest struct {
	Id          int64  `json:"id,string" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (h *Handler) UpdateBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &UpdateBookRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	book := model.NewBook(params.Id, params.Title, params.Description, userId.(int64))
	err := h.service.UpdateBook(ctx, book)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, entity.ConvertToBookVO(book))
}
