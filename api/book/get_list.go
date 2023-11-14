package book

import (
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
	params := &GetBookListRequest{}
	if err := ctx.Bind(params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	books, err := h.service.GetBookList(ctx, params.Search, params.UserId)
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, books)
}
