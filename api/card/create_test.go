package card

import (
	"net/http"
	"net/http/httptest"
	"testing"

	CardHandlerMocks "github.com/Flash-Pass/flash-pass-server/api/card/mocks"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/Flash-Pass/flash-pass-server/internal/testGin"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	addUserIdToContext := testGin.WithCtxPrepare(func(ctx *gin.Context) {
		ctx.Set(constants.CtxUserIdKey, 1)
	})

	middleStepsWithoutMock := testGin.WithMiddleSteps(func(engine *gin.Engine, ctx *gin.Context) {
		engine.POST("/card/", func(ctx *gin.Context) {
			ctx.Handler()(ctx)
		})

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		service := CardHandlerMocks.NewMockService(ctrl)
		handler := NewHandler(service, 1)
		handler.CreateCardController(ctx)
	})

	testGin.NewCase(
		testGin.WithName("create without user id"),
		testGin.WithRequest(http.MethodPost, "/card/", nil),
		middleStepsWithoutMock,
		testGin.WithCheck(func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusUnauthorized, recorder.Code)
		}),
	).Run(t)

	testGin.NewCase(
		testGin.WithName("create with user id but without params"),
		addUserIdToContext,
		testGin.WithRequest(http.MethodPost, "/card/", nil),
		middleStepsWithoutMock,
		testGin.WithCheck(func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusInternalServerError, recorder.Code)
		}),
	).Run(t)

	testGin.NewCase(
		testGin.WithName("create with user id but without question"),
		addUserIdToContext,
		testGin.WithRequest(http.MethodPost, "/card/", CreateCardRequest{
			Answer: "answer",
		}),
		middleStepsWithoutMock,
		testGin.WithCheck(func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusInternalServerError, recorder.Code)
		}),
	).Run(t)

	testGin.NewCase(
		testGin.WithName("create with user id but without answer"),
		addUserIdToContext,
		testGin.WithRequest(http.MethodPost, "/card/", CreateCardRequest{
			Question: "question",
		}),
		middleStepsWithoutMock,
		testGin.WithCheck(func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusInternalServerError, recorder.Code)
		}),
	).Run(t)
}
