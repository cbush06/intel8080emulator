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

func expectUpdateFlags(cndFlags *alumock.MockConditionFlags, original uint8, new uint8) {
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(new)
	cndFlags.EXPECT().UpdateSign(new)
	cndFlags.EXPECT().UpdateParity(new)
	cndFlags.EXPECT().UpdateCarry(original, new)
}

func expectUpdateFlagsExceptCarry(cndFlags *alumock.MockConditionFlags, value uint8) {
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(value)
	cndFlags.EXPECT().UpdateSign(value)
	cndFlags.EXPECT().UpdateParity(value)
}

func TestUpdateFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 255, 255)

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}
	alu.UpdateFlags(255, 255)
}

func TestUpdateFlagsExceptCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(uint8(255))
	cndFlags.EXPECT().UpdateSign(uint8(255))
	cndFlags.EXPECT().UpdateParity(uint8(255))

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}
	alu.UpdateFlagsExceptCarry(255)
}

func TestAddImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 1, 2)

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
	expectUpdateFlags(cndFlags, 1, 3)
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
	cndFlags.EXPECT().UpdateCarryDoublePrecision(uint16(1), uint16(2))

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
	cndFlags.EXPECT().UpdateCarry(uint8(2), uint8(1))

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
	cndFlags.EXPECT().UpdateCarry(uint8(3), uint8(1))

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

func TestIncrement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 2)

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}

	if result := alu.Increment(1); result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}
}

func TestIncrementDouble(t *testing.T) {
	alu := new(ALUImpl)
	result := alu.IncrementDouble(256)

	if result != 257 {
		t.Errorf("Expected 257 but got %d", result)
	}
}

func TestDecrementDouble(t *testing.T) {
	alu := new(ALUImpl)
	result := alu.DecrementDouble(257)

	if result != 256 {
		t.Errorf("Expected 256 but got %d", result)
	}
}

func TestDecrement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 1)

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}

	if result := alu.Decrement(2); result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}
}

func TestRotateRight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	cndFlags.EXPECT().SetCarry()

	alu := &ALUImpl{
		A:              memory.NewRegister(0x55), // 0101 0101b
		ConditionFlags: cndFlags,
	}

	alu.RotateRight()

	var a uint8
	alu.GetA().Read8(&a)

	if a != 0xAA { // 1010 1010b
		t.Errorf("Expected 10101010 but got %bb", a)
	}
}

func TestAndAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 0xA)
	cndFlags.EXPECT().ClearCarry()
	cndFlags.EXPECT().ClearAuxillaryCarry()

	alu := &ALUImpl{
		A:              memory.NewRegister(0xA), // 0000 1010b
		ConditionFlags: cndFlags,
	}

	alu.AndAccumulator(0xF) // 0000 1111b

	var a uint8
	alu.GetA().Read8(&a)

	if a != 0xA { // 0000 1010b
		t.Errorf("Expected 00001010 but got %bb", a)
	}
}

func TestXOrAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 0x5)
	cndFlags.EXPECT().ClearCarry()
	cndFlags.EXPECT().ClearAuxillaryCarry()

	alu := &ALUImpl{
		A:              memory.NewRegister(0xA), // 0000 1010b
		ConditionFlags: cndFlags,
	}

	alu.XOrAccumulator(0xF) // 0000 1111b

	var a uint8
	alu.GetA().Read8(&a)

	if a != 0x5 { // 0000 0101b
		t.Errorf("Expected 00000101 but got %bb", a)
	}
}
