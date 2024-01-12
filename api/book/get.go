package book

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type GetBookRequest struct {
	Id int64 `json:"id,string" form:"id" binding:"required"`
}

func (h *Handler) GetBookController(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	var params GetBookRequest
	if err := c.Bind(&params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}
	book, err := h.service.GetBook(ctx, params.Id)
	if err != nil {
		res.RespondWithError(c, http.StatusNotFound, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}
	res.RespondSuccess(c, book)
}
