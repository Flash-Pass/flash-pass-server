package res

import (
	"bytes"
	"encoding/json"
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

func ResponseSuccessBody(data any) *bytes.Buffer {
	response := &BaseResponse{
		Code:    fpstatus.Success.ErrCode,
		Message: fpstatus.Success.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	byteData, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(byteData)
	return buffer
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

func RespondWithErrorBody(statusCode int, err *fpstatus.ErrNo, data any) *bytes.Buffer {
	response := &BaseResponse{
		Code:    err.ErrCode,
		Message: err.ErrMsg,
	}

	if data != nil {
		response.Data = data
	}

	byteData, errs := json.Marshal(response)
	if errs != nil {
		panic(errs)
	}

	buffer := bytes.NewBuffer(byteData)
	return buffer
}
