package card

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type CreateCardRequest struct {
	Question    string `json:"question" binding:"required"`
	Answer      string `json:"answer" binding:"required"`
	IsAddToBook bool   `json:"is_add_to_book"`
	BookId      int64  `json:"book_id,string"`
}

func (h *Handler) CreateCardController(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	params := &CreateCardRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	card := model.NewCard(
		h.snowflakeHandle.GetId().Int64(), params.Question, params.Answer, userId.(int64),
	)

	if !params.IsAddToBook {
		if err := h.service.CreateCard(ctx, card); err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	} else {
		if err := h.service.CreateCardAndAddToBook(ctx, card, params.BookId); err != nil {
			res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}
	}

	res.RespondSuccess(ctx, card)
}
