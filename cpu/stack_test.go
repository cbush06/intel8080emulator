package cpu

import (
	"testing"

	"github.com/cbush06/intel8080emulator/memory"
)

var (
	pc uint16 = 43690 // 1010 1010 1010 1010
)

func makeCPU(programCounter uint16, memoryBuffer []uint8, stackPointer uint16) *CPU {
	rp := memory.NewRegisterPair(0, 0)
	rp.Write16(stackPointer)

	return &CPU{
		ProgramCounter: programCounter,
		Memory:         memoryBuffer,
		SP:             *rp,
	}
}

func TestCall(t *testing.T) {
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
	var expectedPC uint16 = (258 - 1) // 0x0102 // 0000 0001 0000 0010 // subtract 1 because the PC is incremented before every execution
	if cpu.ProgramCounter != expectedPC {
		t.Errorf("Expected PC to be %d but was %d", expectedPC, cpu.ProgramCounter)
	}
}

func TestRestart(t *testing.T) {
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
