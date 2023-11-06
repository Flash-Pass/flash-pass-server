package res

import (
	"github.com/Flash-Pass/flash-pass-server/internal/fpstatus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func RespondSuccess(ctx *gin.Context, data any) {
	response := &BaseResponse{
		Code:    fpstatus.Success.ErrCode,
		Message: fpstatus.Success.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	ctx.JSON(http.StatusOK, response)
}

func RespondWithError(ctx *gin.Context, statusCode int, err *fpstatus.ErrNo, data any) {
	response := &BaseResponse{
		Code:    err.ErrCode,
		Message: err.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	ctx.AbortWithStatusJSON(statusCode, response)
}
