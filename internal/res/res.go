package res

import (
	"net/http"

	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespondSuccess(c *gin.Context, data any) {
	response := &BaseResponse{
		Code:    fpstatus.Success.ErrCode,
		Message: fpstatus.Success.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	c.JSON(http.StatusOK, response)
}

func RespondWithError(c *gin.Context, statusCode int, err *fpstatus.ErrNo, data any) {
	response := &BaseResponse{
		Code:    err.ErrCode,
		Message: err.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	c.AbortWithStatusJSON(statusCode, response)
}
