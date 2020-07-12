package cpu

import (
	"testing"

	"github.com/golang/mock/gomock"

	alumock "github.com/cbush06/intel8080emulator/alu/mocks"
	"github.com/cbush06/intel8080emulator/memory"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := memory.NewRegister(1)

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediate(uint8(1))
	cpu := &CPU{
		ALU: mALU,
	}
	cpu.Add(r)
}

func TestAddImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediate(uint8(1))

	cpu := &CPU{
		ALU:            mALU,
		Memory:         []uint8{0, 1},
		ProgramCounter: 0,
	}
	cpu.AddImmediate()
}

func TestDoubleAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().DoubleAdd(uint16(256), uint16(1))

	hlValue := 256

	cpu := &CPU{
		ALU: mALU,
		HL:  *memory.NewRegisterPair(uint8(hlValue>>8), uint8((hlValue&0x00FF)>>8)),
	}
	cpu.DoubleAdd(memory.NewRegisterPair(0, 1))
}

func TestSub(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediate(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.Sub(r)
}

func TestAddWithCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediateWithCarry(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.AddWithCarry(r)
}

func TestSubWithBorrow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediateWithBorrow(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.SubWithBorrow(r)
}

func TestIncrementRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().Increment(uint8(1)).Return(uint8(2))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.IncrementRegister(r)

	var rValue uint8
	r.Read8(&rValue)

	if rValue != 2 {
		t.Errorf("Expected 2 but got %d", rValue)
	}
}

func TestIncrementRegisterPair(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().IncrementDouble(uint16(256)).Return(uint16(257))

	r := memory.NewRegisterPair(0, 0)
	r.Write16(uint16(256))

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.IncrementRegisterPair(r)

	var rValue uint16
	r.Read16(&rValue)

	if rValue != 257 {
		t.Errorf("Expected 257 but got %d", rValue)
	}
}

func TestIncrementMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().Increment(uint8(1)).Return(uint8(2))

	hl := memory.NewRegisterPair(0, 1)
	memory := []uint8{0, 1}

	cpu := &CPU{
		ALU:    mALU,
		HL:     *hl,
		Memory: memory,
	}

	cpu.IncrementMemory()

	if memory[1] != 2 {
		t.Errorf("Expected 2 but got %d", memory[1])
	}
}

func TestDecrementRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().Decrement(uint8(1)).Return(uint8(0))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}

	cpu.DecrementRegister(r)

	var rValue uint8
	r.Read8(&rValue)

	if rValue != 0 {
		t.Errorf("Expected 0 but got %d", rValue)
	}
}

func TestDecrementRegisterPair(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().DecrementDouble(uint16(257)).Return(uint16(256))

	r := memory.NewRegisterPair(0, 0)
	r.Write16(257)

	cpu := &CPU{
		ALU: mALU,
	}

	cpu.DecrementRegisterPair(r)

	var rValue uint16
	r.Read16(&rValue)

	if rValue != 256 {
		t.Errorf("Expected 256 but got %d", rValue)
	}
}

func TestDecrementMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().Decrement(uint8(1)).Return(uint8(0))

	hl := memory.NewRegisterPair(0, 1)
	memory := []uint8{0, 1}

	cpu := &CPU{
		ALU:    mALU,
		HL:     *hl,
		Memory: memory,
	}

	cpu.DecrementMemory()

	if memory[1] != 0 {
		t.Errorf("Expected 0 but got %d", memory[1])
	}
}

func TestRotateRight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().RotateRight()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.RotateRight()
}

func TestAndRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AndAccumulator(uint8(1))

	cpu := &CPU{
		ALU: mALU,
	}

	r := memory.NewRegister(1)
	cpu.AndRegister(r)
}

func TestAndMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AndAccumulator(uint8(1))

	hl := memory.NewRegisterPair(0, 1)
	memory := []uint8{0, 1}

	cpu := &CPU{
		HL:     *hl,
		Memory: memory,
		ALU:    mALU,
	}
	cpu.AndMemory()
}

func TestXOrRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().XOrAccumulator(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.XOrRegister(r)
}

func TestXOrMemory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().XOrAccumulator(uint8(1))

	hl := memory.NewRegisterPair(0, 1)
	memory := []uint8{0, 1}

	cpu := &CPU{
		ALU:    mALU,
		HL:     *hl,
		Memory: memory,
	}
	cpu.XOrMemory()
}
