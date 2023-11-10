package card

import (
	CardHandlerMocks "github.com/Flash-Pass/flash-pass-server/api/card/mocks"
	"github.com/Flash-Pass/flash-pass-server/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	//t.Run("create success", func(t *testing.T) {
	//	//runControllerTest()
	//})
	//
	//t.Run("create with empty question", func(t *testing.T) {
	//	runControllerTest(t,
	//		WithExpectedStatusCode(http.StatusInternalServerError),
	//		WithTest(func(t *testing.T, handler IHandler) *http.Request {
	//			req := httptest.NewRequest(http.MethodPost, "/card/", nil)
	//			return req
	//		}),
	//	)
	//})
	gin.SetMode("test")

	cases := []struct {
		ctx      *gin.Context
		ctxPre   func(ctx *gin.Context)
		check    func(recorder *httptest.ResponseRecorder)
		recorder *httptest.ResponseRecorder
		request  *http.Request
		param    gin.Param
		name     string
	}{
		{
			name: "create with empty question",
			ctxPre: func(ctx *gin.Context) {
				ctx.Set(constants.CtxUserIdKey, 1)
			},
			check: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			recorder: httptest.NewRecorder(),
			request:  httptest.NewRequest(http.MethodPost, "/card/", nil),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			tt.ctx = gin.CreateTestContextOnly(tt.recorder, r)
			tt.ctxPre(tt.ctx)
			tt.ctx.Request = tt.request
			tt.ctx.Params = append(tt.ctx.Params, tt.param)
			r.POST("/card/", func(ctx *gin.Context) {
				tt.ctx.Handler()(tt.ctx)
			})
			ctrl := gomock.NewController(t)
			service := CardHandlerMocks.NewMockService(ctrl)
			handler := NewHandler(service, 1)
			handler.CreateCardController(tt.ctx)
			tt.check(tt.recorder)
		})
	}
}
