// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cbush06/intel8080emulator/alu (interfaces: ALU)

// Package alu is a generated GoMock package.
package alu

import (
	memory "github.com/cbush06/intel8080emulator/memory"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockALU is a mock of ALU interface
type MockALU struct {
	ctrl     *gomock.Controller
	recorder *MockALUMockRecorder
}

// MockALUMockRecorder is the mock recorder for MockALU
type MockALUMockRecorder struct {
	mock *MockALU
}

// NewMockALU creates a new mock instance
func NewMockALU(ctrl *gomock.Controller) *MockALU {
	mock := &MockALU{ctrl: ctrl}
	mock.recorder = &MockALUMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockALU) EXPECT() *MockALUMockRecorder {
	return m.recorder
}

// AddImmediate mocks base method
func (m *MockALU) AddImmediate(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddImmediate", arg0)
}

// AddImmediate indicates an expected call of AddImmediate
func (mr *MockALUMockRecorder) AddImmediate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImmediate", reflect.TypeOf((*MockALU)(nil).AddImmediate), arg0)
}

// AddImmediateWithCarry mocks base method
func (m *MockALU) AddImmediateWithCarry(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddImmediateWithCarry", arg0)
}

// AddImmediateWithCarry indicates an expected call of AddImmediateWithCarry
func (mr *MockALUMockRecorder) AddImmediateWithCarry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImmediateWithCarry", reflect.TypeOf((*MockALU)(nil).AddImmediateWithCarry), arg0)
}

// AndAccumulator mocks base method
func (m *MockALU) AndAccumulator(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AndAccumulator", arg0)
}

// AndAccumulator indicates an expected call of AndAccumulator
func (mr *MockALUMockRecorder) AndAccumulator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AndAccumulator", reflect.TypeOf((*MockALU)(nil).AndAccumulator), arg0)
}

// ApplyStatusWord mocks base method
func (m *MockALU) ApplyStatusWord(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyStatusWord", arg0)
}

// ApplyStatusWord indicates an expected call of ApplyStatusWord
func (mr *MockALUMockRecorder) ApplyStatusWord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyStatusWord", reflect.TypeOf((*MockALU)(nil).ApplyStatusWord), arg0)
}

// ClearAuxiliaryCarry mocks base method
func (m *MockALU) ClearAuxiliaryCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearAuxiliaryCarry")
}

// ClearAuxiliaryCarry indicates an expected call of ClearAuxiliaryCarry
func (mr *MockALUMockRecorder) ClearAuxiliaryCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearAuxiliaryCarry", reflect.TypeOf((*MockALU)(nil).ClearAuxiliaryCarry))
}

// ClearCarry mocks base method
func (m *MockALU) ClearCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearCarry")
}

// ClearCarry indicates an expected call of ClearCarry
func (mr *MockALUMockRecorder) ClearCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCarry", reflect.TypeOf((*MockALU)(nil).ClearCarry))
}

// ClearFlags mocks base method
func (m *MockALU) ClearFlags() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearFlags")
}

// ClearFlags indicates an expected call of ClearFlags
func (mr *MockALUMockRecorder) ClearFlags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearFlags", reflect.TypeOf((*MockALU)(nil).ClearFlags))
}

// ClearParity mocks base method
func (m *MockALU) ClearParity() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearParity")
}

// ClearParity indicates an expected call of ClearParity
func (mr *MockALUMockRecorder) ClearParity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearParity", reflect.TypeOf((*MockALU)(nil).ClearParity))
}

// ClearSign mocks base method
func (m *MockALU) ClearSign() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearSign")
}

// ClearSign indicates an expected call of ClearSign
func (mr *MockALUMockRecorder) ClearSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearSign", reflect.TypeOf((*MockALU)(nil).ClearSign))
}

// ClearZero mocks base method
func (m *MockALU) ClearZero() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearZero")
}

// ClearZero indicates an expected call of ClearZero
func (mr *MockALUMockRecorder) ClearZero() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearZero", reflect.TypeOf((*MockALU)(nil).ClearZero))
}

// CompareAccumulator mocks base method
func (m *MockALU) CompareAccumulator(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CompareAccumulator", arg0)
}

// CompareAccumulator indicates an expected call of CompareAccumulator
func (mr *MockALUMockRecorder) CompareAccumulator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareAccumulator", reflect.TypeOf((*MockALU)(nil).CompareAccumulator), arg0)
}

// ComplementAccumulator mocks base method
func (m *MockALU) ComplementAccumulator() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ComplementAccumulator")
}

// ComplementAccumulator indicates an expected call of ComplementAccumulator
func (mr *MockALUMockRecorder) ComplementAccumulator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComplementAccumulator", reflect.TypeOf((*MockALU)(nil).ComplementAccumulator))
}

// CreateStatusWord mocks base method
func (m *MockALU) CreateStatusWord() byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStatusWord")
	ret0, _ := ret[0].(byte)
	return ret0
}

// CreateStatusWord indicates an expected call of CreateStatusWord
func (mr *MockALUMockRecorder) CreateStatusWord() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatusWord", reflect.TypeOf((*MockALU)(nil).CreateStatusWord))
}

// DecimalAdjustAccumulator mocks base method
func (m *MockALU) DecimalAdjustAccumulator() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DecimalAdjustAccumulator")
}

// DecimalAdjustAccumulator indicates an expected call of DecimalAdjustAccumulator
func (mr *MockALUMockRecorder) DecimalAdjustAccumulator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecimalAdjustAccumulator", reflect.TypeOf((*MockALU)(nil).DecimalAdjustAccumulator))
}

// Decrement mocks base method
func (m *MockALU) Decrement(arg0 byte) byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrement", arg0)
	ret0, _ := ret[0].(byte)
	return ret0
}

// Decrement indicates an expected call of Decrement
func (mr *MockALUMockRecorder) Decrement(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrement", reflect.TypeOf((*MockALU)(nil).Decrement), arg0)
}

// DecrementDouble mocks base method
func (m *MockALU) DecrementDouble(arg0 uint16) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrementDouble", arg0)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// DecrementDouble indicates an expected call of DecrementDouble
func (mr *MockALUMockRecorder) DecrementDouble(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrementDouble", reflect.TypeOf((*MockALU)(nil).DecrementDouble), arg0)
}

// DoubleAdd mocks base method
func (m *MockALU) DoubleAdd(arg0, arg1 uint16) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoubleAdd", arg0, arg1)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// DoubleAdd indicates an expected call of DoubleAdd
func (mr *MockALUMockRecorder) DoubleAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoubleAdd", reflect.TypeOf((*MockALU)(nil).DoubleAdd), arg0, arg1)
}

// GetA mocks base method
func (m *MockALU) GetA() *memory.Register {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetA")
	ret0, _ := ret[0].(*memory.Register)
	return ret0
}

// GetA indicates an expected call of GetA
func (mr *MockALUMockRecorder) GetA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetA", reflect.TypeOf((*MockALU)(nil).GetA))
}

// Increment mocks base method
func (m *MockALU) Increment(arg0 byte) byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Increment", arg0)
	ret0, _ := ret[0].(byte)
	return ret0
}

// Increment indicates an expected call of Increment
func (mr *MockALUMockRecorder) Increment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Increment", reflect.TypeOf((*MockALU)(nil).Increment), arg0)
}

// IncrementDouble mocks base method
func (m *MockALU) IncrementDouble(arg0 uint16) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementDouble", arg0)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// IncrementDouble indicates an expected call of IncrementDouble
func (mr *MockALUMockRecorder) IncrementDouble(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementDouble", reflect.TypeOf((*MockALU)(nil).IncrementDouble), arg0)
}

// IsAuxiliaryCarry mocks base method
func (m *MockALU) IsAuxiliaryCarry() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuxiliaryCarry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAuxiliaryCarry indicates an expected call of IsAuxiliaryCarry
func (mr *MockALUMockRecorder) IsAuxiliaryCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuxiliaryCarry", reflect.TypeOf((*MockALU)(nil).IsAuxiliaryCarry))
}

// IsCarry mocks base method
func (m *MockALU) IsCarry() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCarry")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCarry indicates an expected call of IsCarry
func (mr *MockALUMockRecorder) IsCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCarry", reflect.TypeOf((*MockALU)(nil).IsCarry))
}

// IsParity mocks base method
func (m *MockALU) IsParity() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsParity")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsParity indicates an expected call of IsParity
func (mr *MockALUMockRecorder) IsParity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsParity", reflect.TypeOf((*MockALU)(nil).IsParity))
}

// IsSign mocks base method
func (m *MockALU) IsSign() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSign")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSign indicates an expected call of IsSign
func (mr *MockALUMockRecorder) IsSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSign", reflect.TypeOf((*MockALU)(nil).IsSign))
}

// IsZero mocks base method
func (m *MockALU) IsZero() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsZero")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsZero indicates an expected call of IsZero
func (mr *MockALUMockRecorder) IsZero() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsZero", reflect.TypeOf((*MockALU)(nil).IsZero))
}

// OrAccumulator mocks base method
func (m *MockALU) OrAccumulator(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OrAccumulator", arg0)
}

// OrAccumulator indicates an expected call of OrAccumulator
func (mr *MockALUMockRecorder) OrAccumulator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrAccumulator", reflect.TypeOf((*MockALU)(nil).OrAccumulator), arg0)
}

// RotateLeft mocks base method
func (m *MockALU) RotateLeft() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RotateLeft")
}

// RotateLeft indicates an expected call of RotateLeft
func (mr *MockALUMockRecorder) RotateLeft() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateLeft", reflect.TypeOf((*MockALU)(nil).RotateLeft))
}

// RotateLeftThroughCarry mocks base method
func (m *MockALU) RotateLeftThroughCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RotateLeftThroughCarry")
}

// RotateLeftThroughCarry indicates an expected call of RotateLeftThroughCarry
func (mr *MockALUMockRecorder) RotateLeftThroughCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateLeftThroughCarry", reflect.TypeOf((*MockALU)(nil).RotateLeftThroughCarry))
}

// RotateRight mocks base method
func (m *MockALU) RotateRight() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RotateRight")
}

// RotateRight indicates an expected call of RotateRight
func (mr *MockALUMockRecorder) RotateRight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateRight", reflect.TypeOf((*MockALU)(nil).RotateRight))
}

// RotateRightThroughCarry mocks base method
func (m *MockALU) RotateRightThroughCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RotateRightThroughCarry")
}

// RotateRightThroughCarry indicates an expected call of RotateRightThroughCarry
func (mr *MockALUMockRecorder) RotateRightThroughCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateRightThroughCarry", reflect.TypeOf((*MockALU)(nil).RotateRightThroughCarry))
}

// SetA mocks base method
func (m *MockALU) SetA(arg0 *memory.Register) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetA", arg0)
}

// SetA indicates an expected call of SetA
func (mr *MockALUMockRecorder) SetA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetA", reflect.TypeOf((*MockALU)(nil).SetA), arg0)
}

// SetAuxiliaryCarry mocks base method
func (m *MockALU) SetAuxiliaryCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAuxiliaryCarry")
}

// SetAuxiliaryCarry indicates an expected call of SetAuxiliaryCarry
func (mr *MockALUMockRecorder) SetAuxiliaryCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAuxiliaryCarry", reflect.TypeOf((*MockALU)(nil).SetAuxiliaryCarry))
}

// SetCarry mocks base method
func (m *MockALU) SetCarry() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCarry")
}

// SetCarry indicates an expected call of SetCarry
func (mr *MockALUMockRecorder) SetCarry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCarry", reflect.TypeOf((*MockALU)(nil).SetCarry))
}

// SetParity mocks base method
func (m *MockALU) SetParity() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParity")
}

// SetParity indicates an expected call of SetParity
func (mr *MockALUMockRecorder) SetParity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParity", reflect.TypeOf((*MockALU)(nil).SetParity))
}

// SetSign mocks base method
func (m *MockALU) SetSign() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSign")
}

// SetSign indicates an expected call of SetSign
func (mr *MockALUMockRecorder) SetSign() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSign", reflect.TypeOf((*MockALU)(nil).SetSign))
}

// SetZero mocks base method
func (m *MockALU) SetZero() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetZero")
}

// SetZero indicates an expected call of SetZero
func (mr *MockALUMockRecorder) SetZero() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetZero", reflect.TypeOf((*MockALU)(nil).SetZero))
}

// SubImmediate mocks base method
func (m *MockALU) SubImmediate(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubImmediate", arg0)
}

// SubImmediate indicates an expected call of SubImmediate
func (mr *MockALUMockRecorder) SubImmediate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubImmediate", reflect.TypeOf((*MockALU)(nil).SubImmediate), arg0)
}

// SubImmediateWithBorrow mocks base method
func (m *MockALU) SubImmediateWithBorrow(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SubImmediateWithBorrow", arg0)
}

// SubImmediateWithBorrow indicates an expected call of SubImmediateWithBorrow
func (mr *MockALUMockRecorder) SubImmediateWithBorrow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubImmediateWithBorrow", reflect.TypeOf((*MockALU)(nil).SubImmediateWithBorrow), arg0)
}

// UpdateAuxiliaryCarry mocks base method
func (m *MockALU) UpdateAuxiliaryCarry(arg0, arg1 byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuxiliaryCarry", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateAuxiliaryCarry indicates an expected call of UpdateAuxiliaryCarry
func (mr *MockALUMockRecorder) UpdateAuxiliaryCarry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuxiliaryCarry", reflect.TypeOf((*MockALU)(nil).UpdateAuxiliaryCarry), arg0, arg1)
}

// UpdateCarry mocks base method
func (m *MockALU) UpdateCarry(arg0, arg1 byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarry", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateCarry indicates an expected call of UpdateCarry
func (mr *MockALUMockRecorder) UpdateCarry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarry", reflect.TypeOf((*MockALU)(nil).UpdateCarry), arg0, arg1)
}

// UpdateCarryDoublePrecision mocks base method
func (m *MockALU) UpdateCarryDoublePrecision(arg0, arg1 uint16) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCarryDoublePrecision", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateCarryDoublePrecision indicates an expected call of UpdateCarryDoublePrecision
func (mr *MockALUMockRecorder) UpdateCarryDoublePrecision(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCarryDoublePrecision", reflect.TypeOf((*MockALU)(nil).UpdateCarryDoublePrecision), arg0, arg1)
}

// UpdateFlags mocks base method
func (m *MockALU) UpdateFlags(arg0, arg1 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateFlags", arg0, arg1)
}

// UpdateFlags indicates an expected call of UpdateFlags
func (mr *MockALUMockRecorder) UpdateFlags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlags", reflect.TypeOf((*MockALU)(nil).UpdateFlags), arg0, arg1)
}

// UpdateFlagsExceptCarry mocks base method
func (m *MockALU) UpdateFlagsExceptCarry(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateFlagsExceptCarry", arg0)
}

// UpdateFlagsExceptCarry indicates an expected call of UpdateFlagsExceptCarry
func (mr *MockALUMockRecorder) UpdateFlagsExceptCarry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlagsExceptCarry", reflect.TypeOf((*MockALU)(nil).UpdateFlagsExceptCarry), arg0)
}

// UpdateParity mocks base method
func (m *MockALU) UpdateParity(arg0 byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParity", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateParity indicates an expected call of UpdateParity
func (mr *MockALUMockRecorder) UpdateParity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParity", reflect.TypeOf((*MockALU)(nil).UpdateParity), arg0)
}

// UpdateSign mocks base method
func (m *MockALU) UpdateSign(arg0 byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSign", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateSign indicates an expected call of UpdateSign
func (mr *MockALUMockRecorder) UpdateSign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSign", reflect.TypeOf((*MockALU)(nil).UpdateSign), arg0)
}

// UpdateZero mocks base method
func (m *MockALU) UpdateZero(arg0 byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateZero", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateZero indicates an expected call of UpdateZero
func (mr *MockALUMockRecorder) UpdateZero(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateZero", reflect.TypeOf((*MockALU)(nil).UpdateZero), arg0)
}

// XOrAccumulator mocks base method
func (m *MockALU) XOrAccumulator(arg0 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "XOrAccumulator", arg0)
}

// XOrAccumulator indicates an expected call of XOrAccumulator
func (mr *MockALUMockRecorder) XOrAccumulator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XOrAccumulator", reflect.TypeOf((*MockALU)(nil).XOrAccumulator), arg0)
}
