// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/card.go

// Package handler is a generated GoMock package.
package handler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/iremsha/oapicodegen-example/internal/entity"
)

// MockCardService is a mock of CardService interface.
type MockCardService struct {
	ctrl     *gomock.Controller
	recorder *MockCardServiceMockRecorder
}

// MockCardServiceMockRecorder is the mock recorder for MockCardService.
type MockCardServiceMockRecorder struct {
	mock *MockCardService
}

// NewMockCardService creates a new mock instance.
func NewMockCardService(ctrl *gomock.Controller) *MockCardService {
	mock := &MockCardService{ctrl: ctrl}
	mock.recorder = &MockCardServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardService) EXPECT() *MockCardServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCardService) Create(card *entity.Card) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", card)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCardServiceMockRecorder) Create(card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCardService)(nil).Create), card)
}

// GetByID mocks base method.
func (m *MockCardService) GetByID(id int64) (entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCardServiceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCardService)(nil).GetByID), id)
}

// GetList mocks base method.
func (m *MockCardService) GetList() ([]entity.Card, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList")
	ret0, _ := ret[0].([]entity.Card)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockCardServiceMockRecorder) GetList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockCardService)(nil).GetList))
}
