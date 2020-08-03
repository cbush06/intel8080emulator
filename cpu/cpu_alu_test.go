package cpu

import (
	"testing"

	"github.com/golang/mock/gomock"

	alumock "github.com/cbush06/intel8080emulator/alu/mocks"
	"github.com/cbush06/intel8080emulator/memory"
)

func TestCPU_AddRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := memory.NewRegister(1)

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediate(uint8(1))
	cpu := &CPU{
		ALU: mALU,
	}
	cpu.AddRegister(r)
}

func TestCPU_AddImmediate(t *testing.T) {
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

func TestCPU_AddImmediateWithCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediateWithCarry(uint8(1))

	cpu := &CPU{
		ALU:            mALU,
		Memory:         []uint8{0, 1},
		ProgramCounter: 0,
	}
	cpu.AddImmediateWithCarry()
}

func TestCPU_DoubleAdd(t *testing.T) {
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

func TestCPU_SubtractRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediate(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.SubtractRegister(r)
}

func TestCPU_SubtractImmediate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediate(uint8(0x01))

	cpu := makeCPU(0, []uint8{uint8(SUI), 0x01}, 0)
	cpu.ALU = mALU
	cpu.SubtractImmediate()
}

func TestCPU_SubtractImmediateWithBorrow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediateWithBorrow(uint8(1))

	cpu := makeCPU(0, []uint8{uint8(SBI), 0x01}, 0)
	cpu.ALU = mALU
	cpu.SubtractImmediateWithBorrow()
}

func TestCPU_AddRegisterWithCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().AddImmediateWithCarry(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.AddRegisterWithCarry(r)
}

func TestCPU_SubtractRegisterWithBorrow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SubImmediateWithBorrow(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.SubtractRegisterWithBorrow(r)
}

func TestCPU_IncrementRegister(t *testing.T) {
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

func TestCPU_IncrementRegisterPair(t *testing.T) {
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

func TestCPU_IncrementMemory(t *testing.T) {
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

func TestCPU_DecrementRegister(t *testing.T) {
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

func TestCPU_DecrementRegisterPair(t *testing.T) {
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

func TestCPU_DecrementMemory(t *testing.T) {
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

func TestCPU_RotateRight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().RotateRight()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.RotateRight()
}

func TestCPU_RotateRightThroughCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().RotateRightThroughCarry()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.RotateRightThroughCarry()
}

func TestCPU_RotateLeft(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().RotateLeft()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.RotateLeft()
}

func TestCPU_RotateLeftThroughCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().RotateLeftThroughCarry()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.RotateLeftThroughCarry()
}

func TestCPU_AndRegister(t *testing.T) {
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

func TestCPU_AndMemory(t *testing.T) {
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

func TestCPU_OrRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().OrAccumulator(uint8(1))

	r := memory.NewRegister(1)

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.OrRegister(r)
}

func TestCPU_XOrRegister(t *testing.T) {
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

func TestCPU_XOrMemory(t *testing.T) {
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

func TestCPU_DecimalAccumulatorAdjust(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().DecimalAdjustAccumulator()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.DecimalAccumulatorAdjust()
}

func TestCPU_ComplementAccumulator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().ComplementAccumulator()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.ComplementAccumulator()
}

func TestCPU_SetCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mALU := alumock.NewMockALU(ctrl)
	mALU.EXPECT().SetCarry()

	cpu := &CPU{
		ALU: mALU,
	}
	cpu.SetCarry()
}

func TestCPU_ComplementCarry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Complement Set Carry Flag", func(t *testing.T) {
		mALU := alumock.NewMockALU(ctrl)
		mALU.EXPECT().IsCarry().Return(true)
		mALU.EXPECT().ClearCarry()
		cpu := &CPU{
			ALU: mALU,
		}
		cpu.ComplementCarry()
	})

	t.Run("Complement Clear Carry Flag", func(t *testing.T) {
		mALU := alumock.NewMockALU(ctrl)
		mALU.EXPECT().IsCarry().Return(false)
		mALU.EXPECT().SetCarry()
		cpu := &CPU{
			ALU: mALU,
		}
		cpu.ComplementCarry()
	})
}
