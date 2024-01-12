// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package CardHandlerMocks is a generated GoMock package.
package CardHandlerMocks

import (
	reflect "reflect"

	model "github.com/Flash-Pass/flash-pass-server/db/model"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIHandler is a mock of IHandler interface.
type MockIHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIHandlerMockRecorder
}

// MockIHandlerMockRecorder is the mock recorder for MockIHandler.
type MockIHandlerMockRecorder struct {
	mock *MockIHandler
}

// NewMockIHandler creates a new mock instance.
func NewMockIHandler(ctrl *gomock.Controller) *MockIHandler {
	mock := &MockIHandler{ctrl: ctrl}
	mock.recorder = &MockIHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHandler) EXPECT() *MockIHandlerMockRecorder {
	return m.recorder
}

// CreateCardController mocks base method.
func (m *MockIHandler) CreateCardController(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateCardController", ctx)
}

// CreateCardController indicates an expected call of CreateCardController.
func (mr *MockIHandlerMockRecorder) CreateCardController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCardController", reflect.TypeOf((*MockIHandler)(nil).CreateCardController), ctx)
}

// DeleteCardController mocks base method.
func (m *MockIHandler) DeleteCardController(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCardController", ctx)
}

// DeleteCardController indicates an expected call of DeleteCardController.
func (mr *MockIHandlerMockRecorder) DeleteCardController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardController", reflect.TypeOf((*MockIHandler)(nil).DeleteCardController), ctx)
}

// GetCardController mocks base method.
func (m *MockIHandler) GetCardController(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetCardController", ctx)
}

// GetCardController indicates an expected call of GetCardController.
func (mr *MockIHandlerMockRecorder) GetCardController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardController", reflect.TypeOf((*MockIHandler)(nil).GetCardController), ctx)
}

// GetCardListController mocks base method.
func (m *MockIHandler) GetCardListController(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetCardListController", ctx)
}

// GetCardListController indicates an expected call of GetCardListController.
func (mr *MockIHandlerMockRecorder) GetCardListController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardListController", reflect.TypeOf((*MockIHandler)(nil).GetCardListController), ctx)
}

// UpdateCardController mocks base method.
func (m *MockIHandler) UpdateCardController(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateCardController", ctx)
}

// UpdateCardController indicates an expected call of UpdateCardController.
func (mr *MockIHandlerMockRecorder) UpdateCardController(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCardController", reflect.TypeOf((*MockIHandler)(nil).UpdateCardController), ctx)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateCard mocks base method.
func (m *MockService) CreateCard(c *gin.Context, card *model.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCard", ctx, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCard indicates an expected call of CreateCard.
func (mr *MockServiceMockRecorder) CreateCard(ctx, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCard", reflect.TypeOf((*MockService)(nil).CreateCard), ctx, card)
}

// DeleteCard mocks base method.
func (m *MockService) DeleteCard(c *gin.Context, id, userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCard", ctx, id, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCard indicates an expected call of DeleteCard.
func (mr *MockServiceMockRecorder) DeleteCard(ctx, id, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCard", reflect.TypeOf((*MockService)(nil).DeleteCard), ctx, id, userId)
}

// GetCard mocks base method.
func (m *MockService) GetCard(c *gin.Context, id int64) (*model.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCard", ctx, id)
	ret0, _ := ret[0].(*model.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCard indicates an expected call of GetCard.
func (mr *MockServiceMockRecorder) GetCard(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCard", reflect.TypeOf((*MockService)(nil).GetCard), ctx, id)
}

// GetCardList mocks base method.
func (m *MockService) GetCardList(c *gin.Context, search string, userId int64) ([]*model.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardList", ctx, search, userId)
	ret0, _ := ret[0].([]*model.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCardList indicates an expected call of GetCardList.
func (mr *MockServiceMockRecorder) GetCardList(ctx, search, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardList", reflect.TypeOf((*MockService)(nil).GetCardList), ctx, search, userId)
}

// UpdateCard mocks base method.
func (m *MockService) UpdateCard(c *gin.Context, card *model.Card) (*model.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCard", ctx, card)
	ret0, _ := ret[0].(*model.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCard indicates an expected call of UpdateCard.
func (mr *MockServiceMockRecorder) UpdateCard(ctx, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCard", reflect.TypeOf((*MockService)(nil).UpdateCard), ctx, card)
}
