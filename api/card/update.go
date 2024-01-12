package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type UpdateCardRequest struct {
	Id       int64  `json:"id" binding:"required"`
	Question string `json:"question" binding:"stringNotBothEmpty=Answer"`
	Answer   string `json:"answer"`
}

func (h *Handler) UpdateCardController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &UpdateCardRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	card, err := h.service.UpdateCard(ctx, model.NewCard(
		params.Id, params.Question, params.Answer, userId.(int64),
	))
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, card)
}
