package book

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
)

type GetBookListRequest struct {
	UserId int64  `json:"user_id,string" form:"user_id,string"`
	Search string `json:"search" form:"search"`
}

func (h *Handler) GetBookListController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)
	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("parse token error"), nil)
		return
	}

	params := &GetBookListRequest{}
	if err := c.Bind(params); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	books, err := h.service.GetBookList(ctx, params.Search, userId.(int64))
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, books)
}
