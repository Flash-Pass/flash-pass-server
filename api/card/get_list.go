package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetCardListRequest struct {
	Search string `json:"search"`
	UserId int64  `json:"id"`
}

func (h *Handler) GetCardListController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	params := &GetCardListRequest{}
	if err := c.Bind(params); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	if params.UserId == 0 {
		userId, ok := c.Get("userId")
		if !ok {
			res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("parse token error"), nil)
			return
		}
		params.UserId = userId.(int64)
	}

	cards, err := h.service.GetCardList(ctx, params.Search, params.UserId)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, cards)
}
