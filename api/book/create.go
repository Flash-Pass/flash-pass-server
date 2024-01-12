package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/entity"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
)

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AddCardToBookRequest struct {
	BookId int64 `json:"book_id,string" binding:"required"`
	CardId int64 `json:"card_id,string" binding:"required"`
}

func (h *Handler) CreateBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &CreateBookRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	// TODO model 对象的构建是否放在 service 层会更好，controller 层只负责参数校验和调用 service 层
	book := model.NewBook(
		h.snowflakeHandle.GetId().Int64(), params.Title, params.Description, userId.(int64),
	)
	if err := h.service.CreateBook(ctx, book); err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, entity.ConvertToBookVO(book))
}

func (h *Handler) AddCardToBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &AddCardToBookRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	err := h.service.AddCardToBook(ctx, model.NewBookCard(
		h.snowflakeHandle.GetId().Int64(), params.BookId, params.CardId, userId.(int64),
	))
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, nil)
}
