package plan

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type getListRequest struct {
	userId int64 `json:"id"`
}

func (h *Handler) GetList(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.ParseTokenError, nil)
		return
	}

	param := &getListRequest{}
	if err := c.ShouldBind(param); err != nil {
		res.RespondWithError(c, http.StatusBadRequest, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	if param.userId != 0 {
		planList, err := h.service.GetList(ctx, param.userId)
		if err != nil {
			res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}

		res.RespondSuccess(c, planList)
	} else {
		planList, err := h.service.GetList(ctx, userId.(int64))
		if err != nil {
			res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
			return
		}

		res.RespondSuccess(c, planList)
	}
}
