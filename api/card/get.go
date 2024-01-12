package card

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetCardRequest struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}

func (h *Handler) GetCardController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	var params GetCardRequest
	if err := c.Bind(&params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	card, err := h.service.GetCard(ctx, params.Id)
	if err != nil {
		res.RespondWithError(c, http.StatusNotFound, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, card)
}
