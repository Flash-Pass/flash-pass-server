package user

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/Flash-Pass/flash-pass-server/internal/paramValidator"
	"github.com/Flash-Pass/flash-pass-server/internal/res"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerViaMobileRequest struct {
	Mobile   string `json:"mobile" binding:"required,stringIsMobile"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) registerViaMobile(c *gin.Context) {
	params := &registerViaMobileRequest{}
	if err := c.ShouldBind(params); err != nil {
		paramValidator.RespondWithParamError(c, err)
		return
	}

	token, err := h.service.Register(c, params.Mobile, params.Password)
	if err != nil {
		res.RespondWithError(c, http.StatusInternalServerError, fpstatus.SystemError.WithMessage(err.Error()), nil)
		return
	}

	res.RespondSuccess(c, token)
}
