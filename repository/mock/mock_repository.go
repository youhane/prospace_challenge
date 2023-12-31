// Code generated by MockGen. DO NOT EDIT.
// Source: prospace_challenge/repository (interfaces: InterGalacticConverter)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterGalacticConverter is a mock of InterGalacticConverter interface.
type MockInterGalacticConverter struct {
	ctrl     *gomock.Controller
	recorder *MockInterGalacticConverterMockRecorder
}

// MockInterGalacticConverterMockRecorder is the mock recorder for MockInterGalacticConverter.
type MockInterGalacticConverterMockRecorder struct {
	mock *MockInterGalacticConverter
}

// NewMockInterGalacticConverter creates a new mock instance.
func NewMockInterGalacticConverter(ctrl *gomock.Controller) *MockInterGalacticConverter {
	mock := &MockInterGalacticConverter{ctrl: ctrl}
	mock.recorder = &MockInterGalacticConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterGalacticConverter) EXPECT() *MockInterGalacticConverterMockRecorder {
	return m.recorder
}

// ConvertRomanToArabic mocks base method.
func (m *MockInterGalacticConverter) ConvertRomanToArabic(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertRomanToArabic", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// ConvertRomanToArabic indicates an expected call of ConvertRomanToArabic.
func (mr *MockInterGalacticConverterMockRecorder) ConvertRomanToArabic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertRomanToArabic", reflect.TypeOf((*MockInterGalacticConverter)(nil).ConvertRomanToArabic), arg0)
}

// ConvertToCredits mocks base method.
func (m *MockInterGalacticConverter) ConvertToCredits(arg0, arg1 string) (float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertToCredits", arg0, arg1)
	ret0, _ := ret[0].(float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertToCredits indicates an expected call of ConvertToCredits.
func (mr *MockInterGalacticConverterMockRecorder) ConvertToCredits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertToCredits", reflect.TypeOf((*MockInterGalacticConverter)(nil).ConvertToCredits), arg0, arg1)
}

// ConvertToIntergalactic mocks base method.
func (m *MockInterGalacticConverter) ConvertToIntergalactic(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertToIntergalactic", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// ConvertToIntergalactic indicates an expected call of ConvertToIntergalactic.
func (mr *MockInterGalacticConverterMockRecorder) ConvertToIntergalactic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertToIntergalactic", reflect.TypeOf((*MockInterGalacticConverter)(nil).ConvertToIntergalactic), arg0)
}

// DefineMaterial mocks base method.
func (m *MockInterGalacticConverter) DefineMaterial(arg0 []string, arg1 string, arg2 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DefineMaterial", arg0, arg1, arg2)
}

// DefineMaterial indicates an expected call of DefineMaterial.
func (mr *MockInterGalacticConverterMockRecorder) DefineMaterial(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefineMaterial", reflect.TypeOf((*MockInterGalacticConverter)(nil).DefineMaterial), arg0, arg1, arg2)
}

// DefineUnit mocks base method.
func (m *MockInterGalacticConverter) DefineUnit(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefineUnit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DefineUnit indicates an expected call of DefineUnit.
func (mr *MockInterGalacticConverterMockRecorder) DefineUnit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefineUnit", reflect.TypeOf((*MockInterGalacticConverter)(nil).DefineUnit), arg0, arg1)
}
