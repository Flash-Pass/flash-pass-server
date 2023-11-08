package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteCardRequest struct {
	Id string `json:"id"`
}

func (h *Handler) DeleteCardController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
	}

	params := &DeleteCardRequest{}
	if err := ctx.ShouldBind(&params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	if err := h.service.DeleteCard(ctx, params.Id, userId.(string)); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, nil)
}
