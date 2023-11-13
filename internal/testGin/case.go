package testGin

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Case struct {
	name        string
	param       gin.Param
	ctx         *gin.Context
	ctxPrepare  func(ctx *gin.Context)
	middleSteps func(engine *gin.Engine, ctx *gin.Context)
	check       func(recorder *httptest.ResponseRecorder)
	recorder    *httptest.ResponseRecorder
	request     *http.Request
}

type CaseOption func(c *Case)

func NewCase(opts ...CaseOption) *Case {
	c := &Case{}
	for _, opt := range opts {
		opt(c)
	}

	if c.recorder == nil {
		c.recorder = httptest.NewRecorder()
	}

	return c
}

func WithName(name string) CaseOption {
	return func(c *Case) {
		c.name = name
	}
}

func WithCtxPrepare(ctxPrepare func(ctx *gin.Context)) CaseOption {
	return func(c *Case) {
		c.ctxPrepare = ctxPrepare
	}
}

func WithMiddleSteps(middleSteps func(engine *gin.Engine, ctx *gin.Context)) CaseOption {
	return func(c *Case) {
		c.middleSteps = middleSteps
	}
}

func WithCheck(check func(recorder *httptest.ResponseRecorder)) CaseOption {
	return func(c *Case) {
		c.check = check
	}
}

func WithRecorder(recorder *httptest.ResponseRecorder) CaseOption {
	return func(c *Case) {
		c.recorder = recorder
	}
}

func WithRequest(method, target string, body interface{}) CaseOption {
	buffer := &bytes.Buffer{}
	if body != nil {
		kind := reflect.TypeOf(body).Kind()
		if kind != reflect.Struct {
			panic("body must be a struct")
		}

		data, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		buffer = bytes.NewBuffer(data)
	}

	return func(c *Case) {
		c.request = httptest.NewRequest(method, target, buffer)
	}
}

func WithParam(param gin.Param) CaseOption {
	return func(c *Case) {
		c.param = param
	}
}

func (c *Case) Run(t *testing.T) {
	t.Run(
		c.name,
		func(t *testing.T) {
			gin.SetMode("test")
			r := gin.Default()
			c.ctx = gin.CreateTestContextOnly(c.recorder, r)
			c.ctx.Request = c.request
			c.ctx.Params = append(c.ctx.Params, c.param)

			if c.ctxPrepare != nil {
				c.ctxPrepare(c.ctx)
			}

			if c.middleSteps != nil {
				c.middleSteps(r, c.ctx)
			}

			if c.check == nil {
				c.check(c.recorder)
			}
		})
}
