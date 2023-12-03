package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetBookListRequest struct {
	UserId int64  `json:"user_id,string" form:"user_id,string"`
	Search string `json:"search" form:"search"`
}

func (h *Handler) GetBookListController(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("parse token error"), nil)
		return
	}

	params := &GetBookListRequest{}
	if err := ctx.Bind(params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	books, err := h.service.GetBookList(ctx, params.Search, userId.(int64))
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, books)
}
