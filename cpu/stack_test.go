package cpu

import (
	"github.com/cbush06/intel8080emulator/memory"
	"testing"
)

var (
	pc uint16 = 43690 // 1010 1010 1010 1010
)

func TestCPU_Call(t *testing.T) {
	cpu := makeCPU(0, []uint8{uint8(CALL), 2, 1, 0, 0}, 5)

	cpu.Call()

	// Verify next instruction ADDRESS is stored in SP-1 and SP-2
	if cpu.Memory[4] != uint8(0x00) {
		t.Errorf("Expected Memory[SP - 1] to be 0x0 but was 0x%X", cpu.Memory[4])
	}
	if cpu.Memory[3] != uint8(0x03) {
		t.Errorf("Expected memory[SP - 2] to be 0x3 but was 0x%X", cpu.Memory[3])
	}

	// Verify SP is decremented twice
	var sp uint16
	cpu.SP.Read16(&sp)
	if sp != 3 {
		t.Errorf("Expected SP to be 3 but was %d", sp)
	}

	// Verify PC is set to CALL address
	var expectedPC uint16 = 258
	if cpu.ProgramCounter != expectedPC {
		t.Errorf("Expected PC to be %d but was %d", expectedPC, cpu.ProgramCounter)
	}
}

func TestCPU_Restart(t *testing.T) {
	var rstOpCodes = [8]OpCode{
		RST0,
		RST1,
		RST2,
		RST3,
		RST4,
		RST5,
		RST6,
		RST7,
	}

	var rstAddrs = [8]uint8{
		(uint8(RST0) & 0x38) >> 3,
		(uint8(RST1) & 0x38) >> 3,
		(uint8(RST2) & 0x38) >> 3,
		(uint8(RST3) & 0x38) >> 3,
		(uint8(RST4) & 0x38) >> 3,
		(uint8(RST5) & 0x38) >> 3,
		(uint8(RST6) & 0x38) >> 3,
		(uint8(RST7) & 0x38) >> 3,
	}

	var cpu *CPU

	for i, opcode := range rstOpCodes {
		cpu = makeCPU(0, []uint8{uint8(opcode), 0, 0, 0, 0}, 5)

		cpu.Restart(opcode)

		// Verify next instruction ADDRESS is stored in SP-1 and SP-2
		if cpu.Memory[4] != uint8(0x00) {
			t.Errorf("Expected Memory[SP - 1] to be 0x0 but was 0x%X", cpu.Memory[4])
		}
		if cpu.Memory[3] != uint8(0x01) {
			t.Errorf("Expected memory[SP - 2] to be 0x1 but was 0x%X", cpu.Memory[3])
		}

		// Verify SP is decremented twice
		var sp uint16
		cpu.SP.Read16(&sp)
		if sp != 3 {
			t.Errorf("Expected SP to be 3 but was %d", sp)
		}

		// Verify PC is set to RST address
		var expectedPC uint16 = uint16(rstAddrs[i] * 8)
		if cpu.ProgramCounter != expectedPC {
			t.Errorf("Expected PC to be %d but was %d", expectedPC, cpu.ProgramCounter)
		}
	}
}

func TestCPU_Return(t *testing.T) {
	cpu := makeCPU(0, []uint8{uint8(CALL), 2, 1, 0, 0}, 1)

	cpu.Return()

	var (
		highPC = uint8((cpu.ProgramCounter & 0xFF00) >> 8)
		lowPC  = uint8(cpu.ProgramCounter & 0xFF)
	)

	// Confirm Memory[SP+1] is in high-order 8 bits of ProgramCounter
	if highPC != 1 {
		t.Errorf("Expected high-order bits of PC to be 0x1 but was 0x%X", highPC)
	}

	// Confirm Memory[SP] is in low-order 8 bits of ProgramCounter
	if lowPC != 2 {
		t.Errorf("Expected low-order bits of PC to be 0x2 but was 0x%X", lowPC)
	}

	// Confirm SP is incremented by 2
	var sp uint16
	cpu.SP.Read16(&sp)
	if sp != 3 {
		t.Errorf("Expected stack pointer to be 3 but was %d", sp)
	}
}

func TestCPU_ExchangeStackTopWithHandL(t *testing.T) {
	cpu := makeCPU(0, []uint8{0, 0x67, 0x89, 0}, 1)
	cpu.L = memory.NewRegister(0xAB)
	cpu.H = memory.NewRegister(0xCD)

	cpu.ExchangeStackTopWithHandL()

	if cpu.Memory[1] != 0xAB || cpu.Memory[2] != 0xCD {
		t.Errorf("Expected new stack top to be 0xABCD but was 0x%X%X", cpu.Memory[1], cpu.Memory[2])
	}

	var l uint8
	var h uint8

	cpu.L.Read8(&l)
	cpu.H.Read8(&h)

	if l != 0x67 || h != 0x89 {
		t.Errorf("Expected HL register to be 0x6789 but was 0x%X%X", l, h)
	}
}

func TestCPU_EnableInterrupts(t *testing.T) {
	cpu := makeCPU(0, []uint8{0, 0}, 0)
	cpu.InterruptsEnabled = false
	cpu.EnableInterrupts()
	if !cpu.InterruptsEnabled {
		t.Error("Expected interrupts to be enabled but were not")
	}
}

func TestCPU_DisableInterrupts(t *testing.T) {
	cpu := makeCPU(0, []uint8{0, 0}, 0)
	cpu.InterruptsEnabled = true
	cpu.DisableInterrupts()
	if cpu.InterruptsEnabled {
		t.Error("Expected interrupts to be disabled but were not")
	}
}

func TestCPU_MoveHLToSP(t *testing.T) {
	var hl uint16
	var sp uint16

	cpu := &CPU{
		HL: *memory.NewRegisterPair(0xAB, 0xCD),
		SP: *memory.NewRegisterPair(0xCD, 0xDE),
	}
	cpu.MoveHLToSP()

	cpu.HL.Read16(&hl)
	cpu.SP.Read16(&sp)

	if hl != 0xCDDE || sp != 0xABCD {
		t.Errorf("Expected HL to be 0xCDDE and SP to be 0xABCD but HL was 0x%X and SP was 0x%X", hl, sp)
	}
}
