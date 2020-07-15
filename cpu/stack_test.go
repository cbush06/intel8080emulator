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

	// Verify next instruction is stored in SP-1 and SP-2
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
