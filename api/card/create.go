package card

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCardRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (h *Handler) CreateCardController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &CreateCardRequest{}
	if err := ctx.ShouldBind(&params); err != nil {
		res.RespondWithError(ctx, http.StatusBadRequest, fpstatus.ParseParametersError, nil)
		return
	}

	card := model.NewCard(
		h.snowflakeHandle.GetId().String(), params.Question, params.Answer, userId.(string),
	)
	if err := h.service.CreateCard(ctx, card); err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, card)
}
