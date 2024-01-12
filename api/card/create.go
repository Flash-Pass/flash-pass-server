package card

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
)

type CreateCardRequest struct {
	Question    string `json:"question" binding:"required"`
	Answer      string `json:"answer" binding:"required"`
	IsAddToBook bool   `json:"is_add_to_book"`
	BookId      int64  `json:"book_id,string"`
}

func (h *Handler) CreateCardController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get("userId")
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &CreateCardRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	card := model.NewCard(
		h.snowflakeHandle.GetId().Int64(), params.Question, params.Answer, userId.(int64),
	)

	if !params.IsAddToBook {
		if err := h.service.CreateCard(ctx, card); err != nil {
			res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	} else {
		if err := h.service.CreateCardAndAddToBook(ctx, card, params.BookId); err != nil {
			res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	}

	res.RespondSuccess(c, card)
}
