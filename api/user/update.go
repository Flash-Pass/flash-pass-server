package user

import (
	"github.com/Flash-Pass/flash-pass-server/internal/ctxlog"
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Nickname string `json:"nickname" binding:"required"`
}

func (h *Handler) update(c *gin.Context) {
	ctx, _ := ctxlog.Export(c)

	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	var params UpdateUserRequest
	if err := c.ShouldBind(&params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	user, err := h.service.Update(ctx, &model.User{
		Base: model.Base{
			Id: userId.(int64),
		},
		Nickname: params.Nickname,
	})
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, user)
}
