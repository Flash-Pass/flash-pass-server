package user

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginViaWeChatRequest struct {
	code string `json:"code" binding:"required"`
}

func (h *Handler) loginViaWeChat(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	params := &loginViaWeChatRequest{}
	if err := c.ShouldBindJSON(params); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	token, err := h.service.LoginViaWeChat(ctx, params.code)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, token)
}
