package cpu

import "github.com/cbush06/intel8080emulator/memory"

func makeCPU(programCounter uint16, memoryBuffer []uint8, stackPointer uint16) *CPU {
	rp := memory.NewRegisterPair(0, 0)
	rp.Write16(stackPointer)

	return &CPU{
		ProgramCounter: programCounter,
		Memory:         memoryBuffer,
		SP:             *rp,
	}
}
