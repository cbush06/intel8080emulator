package alu

import (
	"testing"

	"github.com/golang/mock/gomock"

	alumock "github.com/cbush06/intel8080emulator/alu/mocks"
	"github.com/cbush06/intel8080emulator/memory"
)

func TestGetA(t *testing.T) {
	a := memory.NewRegister(1)
	alu := &ALUImpl{
		A: a,
	}

	if alu.GetA() != a {
		t.Error("Expected same object but was not")
	}
}

func TestSetA(t *testing.T) {
	a := memory.NewRegister(1)
	alu := &ALUImpl{
		A: memory.NewRegister(2),
	}

	alu.SetA(a)
	if alu.GetA() != a {
		t.Error("Expected same object but was not")
	}
}

func expectUpdateFlags(cndFlags *alumock.MockConditionFlags, value uint16) {
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(value)
	cndFlags.EXPECT().UpdateSign(value)
	cndFlags.EXPECT().UpdateParity(value)
	cndFlags.EXPECT().UpdateCarry(value)
}

func expectUpdateFlagsExceptCarry(cndFlags *alumock.MockConditionFlags, value uint16) {
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(value)
	cndFlags.EXPECT().UpdateSign(value)
	cndFlags.EXPECT().UpdateParity(value)
}

func TestUpdateFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 256)

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}
	alu.UpdateFlags(256)
}

func TestUpdateFlagsExceptCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(uint16(256))
	cndFlags.EXPECT().UpdateSign(uint16(256))
	cndFlags.EXPECT().UpdateParity(uint16(256))

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}
	alu.UpdateFlagsExceptCarry(256)
}

func TestAddImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 2)

	alu := &ALUImpl{
		A:              memory.NewRegister(1),
		ConditionFlags: cndFlags,
	}

	alu.AddImmediate(1)

	var aValue uint8
	alu.GetA().Read8(&aValue)

	if aValue != 2 {
		t.Errorf("Expected 2 but got %d", aValue)
	}
}

func TestAddImmediateWithCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 3)
	cndFlags.EXPECT().IsCarry().Return(true)

	alu := &ALUImpl{
		A:              memory.NewRegister(1),
		ConditionFlags: cndFlags,
	}

	alu.AddImmediateWithCarry(1)

	var aValue uint8
	alu.GetA().Read8(&aValue)

	if aValue != 3 {
		t.Errorf("Expected 3 but got %d", aValue)
	}
}

func TestDoubleAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	cndFlags.EXPECT().UpdateCarryDoublePrecision(uint32(2))

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}

	if result := alu.DoubleAdd(1, 1); result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}
}

func TestSubImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 1)
	cndFlags.EXPECT().UpdateBorrow(uint8(2), uint8(1))

	alu := &ALUImpl{
		A:              memory.NewRegister(2),
		ConditionFlags: cndFlags,
	}

	alu.SubImmediate(1)

	var result uint8
	alu.GetA().Read8(&result)

	if result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}
}

func TestSubImmediateBorrow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 1)
	cndFlags.EXPECT().IsCarry().Return(true)
	cndFlags.EXPECT().UpdateBorrow(uint8(3), uint8(2))

	alu := &ALUImpl{
		A:              memory.NewRegister(3),
		ConditionFlags: cndFlags,
	}

	alu.SubImmediateWithBorrow(1)

	var result uint8
	alu.GetA().Read8(&result)

	if result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}
}
