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

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AddCardToBookRequest struct {
	BookId int64 `json:"bookId,string" binding:"required"`
	CardId int64 `json:"cardId,string" binding:"required"`
}

func (h *Handler) CreateBookController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &CreateBookRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	// TODO model 对象的构建是否放在 service 层会更好，controller 层只负责参数校验和调用 service 层
	book := model.NewBook(
		h.snowflakeHandle.GetId().Int64(), params.Title, params.Description, userId.(int64),
	)
	if err := h.service.CreateBook(ctx, book); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, entity.ConvertToBookVO(book))
}

func (h *Handler) AddCardToBookController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &AddCardToBookRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	err := h.service.AddCardToBook(ctx, model.NewBookCard(
		h.snowflakeHandle.GetId().Int64(), params.BookId, params.CardId, userId.(int64),
	))
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
