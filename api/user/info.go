package user

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
)

type getInfoRequest struct {
	openId string `json:"open_id" form:"open_id"`
	mobile string `json:"mobile" form:"mobile"`
}

func (h *Handler) getInfo(c *gin.Context) {
	userId, ok := c.Get(constants.CtxUserIdKey)
	if !ok {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage("user id not found"), nil)
		return
	}
	params := &getInfoRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	user, err := h.service.GetUser(c, params.openId, params.mobile, userId.(int64))
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, user)
}
