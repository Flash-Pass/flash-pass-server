package user

import (
	"github.com/Flash-Pass/flash-pass-server/db/model"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateUserRequest struct {
	Nickname string `json:"nickname" binding:"required"`
}

func (h *Handler) update(ctx *gin.Context) {
	userId, ok := ctx.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(ctx, http.StatusUnauthorized, fpstatus.ParseTokenError, nil)
		return
	}

	var params UpdateUserRequest
	if err := ctx.ShouldBind(&params); err != nil {
		paramValidator.RespondWithParamError(ctx, err)
		return
	}

	user, err := h.service.Update(ctx, &model.User{
		Base: model.Base{
			Id: userId.(string),
		},
		Nickname: params.Nickname,
	})
	if err != nil {
		res.RespondWithError(ctx, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(ctx, user)
}
