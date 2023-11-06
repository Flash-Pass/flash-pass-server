package card

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateCardRequest struct {
	id       string `json:"id"`
	question string `json:"question"`
	answer   string `json:"answer"`
}

func (h *Handler) UpdateCardController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &UpdateCardRequest{}
	if err := ctx.ShouldBind(&params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	card, err := h.service.UpdateCard(ctx, model.NewCard(
		params.id, params.question, params.answer, userId.(string),
	))
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, card)
}
