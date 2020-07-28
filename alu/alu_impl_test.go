package alu

import (
	"testing"

	"github.com/golang/mock/gomock"

	alumock "github.com/cbush06/intel8080emulator/alu/mocks"
	"github.com/cbush06/intel8080emulator/memory"
)

func TestALUImpl_GetA(t *testing.T) {
	a := memory.NewRegister(1)
	alu := &ALUImpl{
		A: a,
	}

	if alu.GetA() != a {
		t.Error("Expected same object but was not")
	}
}

func TestALUImpl_SetA(t *testing.T) {
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
	cndFlags.EXPECT().UpdateAuxiliaryCarry(original, new)
}

func expectUpdateFlagsExceptCarry(cndFlags *alumock.MockConditionFlags, value uint8) {
	cndFlags.EXPECT().ClearFlags()
	cndFlags.EXPECT().UpdateZero(value)
	cndFlags.EXPECT().UpdateSign(value)
	cndFlags.EXPECT().UpdateParity(value)
}

func TestALUImpl_UpdateFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 255, 255)

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}
	alu.UpdateFlags(255, 255)
}

func TestALUImpl_UpdateFlagsExceptCarry(t *testing.T) {
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

func TestALUImpl_AddImmediate(t *testing.T) {
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

func TestALUImpl_AddImmediateWithCarry(t *testing.T) {
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

func TestALUImpl_DoubleAdd(t *testing.T) {
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

func TestALUImpl_SubImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags, 2, 1)

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

func TestALUImpl_SubImmediateWithBorrow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlags(cndFlags,3, 1)
	cndFlags.EXPECT().IsCarry().Return(true)

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

func TestALUImpl_Increment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 2)
	cndFlags.EXPECT().UpdateAuxiliaryCarry(uint8(1), uint8(2))

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}

	if result := alu.Increment(1); result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}
}

func TestALUImpl_IncrementDouble(t *testing.T) {
	alu := new(ALUImpl)
	result := alu.IncrementDouble(256)

	if result != 257 {
		t.Errorf("Expected 257 but got %d", result)
	}
}

func TestALUImpl_DecrementDouble(t *testing.T) {
	alu := new(ALUImpl)
	result := alu.DecrementDouble(257)

	if result != 256 {
		t.Errorf("Expected 256 but got %d", result)
	}
}

func TestALUImpl_Decrement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 1)
	cndFlags.EXPECT().UpdateAuxiliaryCarry(uint8(2), uint8(1))

	alu := &ALUImpl{
		ConditionFlags: cndFlags,
	}

	if result := alu.Decrement(2); result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}
}

func TestALUImpl_RotateRight(t *testing.T) {
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

func TestALUImpl_RotateRightThroughCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("RotateRightThroughCarry with Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsCarry().Return(true)
		cndFlags.EXPECT().SetCarry()

		alu := &ALUImpl{
			A: memory.NewRegister(0x01),
			ConditionFlags: cndFlags,
		}

		alu.RotateRightThroughCarry()

		var a uint8
		alu.GetA().Read8(&a)

		if a != 0x80 {
			t.Errorf("Expected 10000000 but got %X", a)
		}
	})

	t.Run("RotateRightThroughCarry without Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsCarry().Return(false)
		cndFlags.EXPECT().ClearCarry()

		alu := &ALUImpl{
			A: memory.NewRegister(0x02),
			ConditionFlags: cndFlags,
		}

		alu.RotateRightThroughCarry()

		var a uint8
		alu.GetA().Read8(&a)

		if a != 0x01 {
			t.Errorf("Expected 00000001 but got %X", a)
		}
	})
}

func TestALUImpl_RotateLeft(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	cndFlags.EXPECT().SetCarry()

	alu := &ALUImpl{
		A:              memory.NewRegister(0xAA), // 1010 1010b
		ConditionFlags: cndFlags,
	}

	alu.RotateLeft()

	var a uint8
	alu.GetA().Read8(&a)

	if a != 0x55 { // 0101 0101b
		t.Errorf("Expected 01010101 but got %bb", a)
	}
}

func TestALUImpl_RotateLeftThroughCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("RotateLeftThroughCarry with Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsCarry().Return(true)
		cndFlags.EXPECT().SetCarry()

		alu := &ALUImpl{
			A: memory.NewRegister(0x80),
			ConditionFlags: cndFlags,
		}

		alu.RotateLeftThroughCarry()

		var a uint8
		alu.GetA().Read8(&a)

		if a != 0x01 {
			t.Errorf("Expected 00000001 but got %X", a)
		}
	})

	t.Run("RotateLeftThroughCarry without Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsCarry().Return(false)
		cndFlags.EXPECT().ClearCarry()

		alu := &ALUImpl{
			A: memory.NewRegister(0x40),
			ConditionFlags: cndFlags,
		}

		alu.RotateLeftThroughCarry()

		var a uint8
		alu.GetA().Read8(&a)

		if a != 0x80 {
			t.Errorf("Expected 8000000 but got %X", a)
		}
	})
}

func TestALUImpl_AndAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 0xA)
	cndFlags.EXPECT().UpdateAuxiliaryCarry(uint8(0xA), uint8(0xA))
	cndFlags.EXPECT().ClearCarry()

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

func TestALUImpl_XOrAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cndFlags := alumock.NewMockConditionFlags(ctrl)
	expectUpdateFlagsExceptCarry(cndFlags, 0x5)
	cndFlags.EXPECT().ClearCarry()
	cndFlags.EXPECT().ClearAuxiliaryCarry()

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

func TestALUImpl_DecimalAdjustAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Auxiliary Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsAuxiliaryCarry().Return(true)
		cndFlags.EXPECT().IsCarry().Return(false)
		expectUpdateFlags(cndFlags, 0x11, 0x17)

		alu := &ALUImpl{
			A: memory.NewRegister(0x11),
			ConditionFlags: cndFlags,
		}

		alu.DecimalAdjustAccumulator()
	})

	t.Run("Carry", func(t *testing.T) {
		cndFlags := alumock.NewMockConditionFlags(ctrl)
		cndFlags.EXPECT().IsAuxiliaryCarry().Return(true)
		cndFlags.EXPECT().IsCarry().Return(false)
		expectUpdateFlags(cndFlags, 0x11, 0x17)

		alu := &ALUImpl{
			A: memory.NewRegister(0x11),
			ConditionFlags: cndFlags,
		}

		alu.DecimalAdjustAccumulator()
	})
}

func TestALUImpl_ComplementAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	alu := &ALUImpl{
		A: memory.NewRegister(0xAA),
	}

	alu.ComplementAccumulator()

	var a uint8
	alu.A.Read8(&a)

	if a != 0x55 {
		t.Errorf("Expected %X but got %X", 0x55, a)
	}
}
