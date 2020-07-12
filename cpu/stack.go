package cpu

import "github.com/cbush06/intel8080emulator/memory"

// Call implements the CALL addr instruction. The high-order eight bits of the next instruction address
// are moved to the memory location whose address is one less than the content of register SP. The
// low-order eight bits of the next instruction address are moved to the memory location whose address
// is two less than the content of register SP. The content of register SP is decremented by 2. Control
// is transferred to the instruction whose address is specified in byte 3 and byte 2 of the current
// instruction.
func (cpu *CPU) Call() {
	var stackPointer uint16

	nextInstruction := cpu.ProgramCounter + 1
	cpu.SP.Read16(&stackPointer)

	nextHigh := uint8((nextInstruction & 0xFF00) >> 8)
	nextLow := uint8(nextInstruction & 0xFF)

	cpu.Memory[stackPointer-1] = nextHigh
	cpu.Memory[stackPointer-2] = nextLow

	cpu.SP.Write16(stackPointer - 2)

	// Subtract 1 because 1 will be added on the next execution
	cpu.ProgramCounter = cpu.getJumpAddress() - 1
}

// Restart implements the RST n instruction. The high-order eight bits of the next instruction address
// are moved to the memory location whose address is one less than the content of register SP. The
// low-order eight bits of the next instruction address are moved to the memory location whose
// address is two less than the content of register SP. The content of register SP is decremented by two.
// Control is transferred to the instruction whose address is eight times the content of NNN.
func (cpu *CPU) Restart(opcode OpCode) {
	var stackPointer uint16

	nextInstruction := cpu.ProgramCounter + 1
	cpu.SP.Read16(&stackPointer)

	nextHigh := uint8((nextInstruction & 0xFF00) >> 8)
	nextLow := uint8(nextInstruction & 0xFF)

	cpu.Memory[stackPointer-1] = nextHigh
	cpu.Memory[stackPointer-2] = nextLow

	cpu.SP.Write16(stackPointer - 2)

	// Transfer control to Interrupt Handler by masking all but bits 4, 5, and 6
	// and multiplying their value by 8
	cpu.ProgramCounter = uint16(8 * ((opcode & 0x38) >> 3))
}

// Return implements the RET instruction. The content of the memory location whose address is specified
// in register SP is moved to the low-order eight bits of register PC. The content of the memory location
// whose address is one more than the content of register SP is moved to the high-order eight bits of
// register PC. The content of register SP is incremented by 2.
func (cpu *CPU) Return() {
	var stackPointer uint16
	var newProgramCounter uint16

	cpu.SP.Read16(&stackPointer)
	newProgramCounter |= uint16(cpu.Memory[stackPointer])
	newProgramCounter |= uint16(cpu.Memory[stackPointer+1]) << 8

	cpu.SP.Write16(stackPointer + 2)

	// Subtract 1 because 1 will be added on the next execution
	cpu.ProgramCounter = newProgramCounter - 1
}

// Push implements the PUSH rp instruction. The content of the high-order register of register pair
// rp is moved to the memory location whose address is one less than the content of register SP. The
// content of the low-order register of register pair rp is moved to the memory location whose
// address is two less than the content of register SP. The cont~nt of register SP is decremented by
// 2. Note: Register pair rp = SP may not be specified.
func (cpu *CPU) Push(rp *memory.RegisterPair) {
	if rp == &cpu.SP {
		panic("PUSH rp called where rp == SP")
	}

	var stackPointer uint16
	cpu.SP.Read16(&stackPointer)
	rp.ReadHigh(&cpu.Memory[stackPointer-1])
	rp.ReadLow(&cpu.Memory[stackPointer-2])
	cpu.SP.Write16(stackPointer - 2)
}

// PushProcessorStatusWord implements the PUSH PSW instruction. The content of register A is moved to the
// memory location whose address is one less than register SP. The contents of the condition flags are
// assembled into a processor status word and the word is moved to the memory location whose address is
// two less than the content of register SP. The content of register SP is decremented by two.
func (cpu *CPU) PushProcessorStatusWord() {
	var stackPointer uint16
	cpu.SP.Read16(&stackPointer)
	cpu.A.Read8(&cpu.Memory[stackPointer-1])
	cpu.Memory[stackPointer-2] = cpu.ALU.CreateStatusWord()
	cpu.SP.Write16(stackPointer - 2)
}

// Pop implements the POP rp instruction. The content of the memory location, whose address
// is specified by the content of register SP, is moved to the low-order register of register
// pair rp. The content of the memory location, whose address is one more than the content of
// register SP, is moved to the high-order register of register pair rp. The content of
// register SP is incremented by 2. Note: Register pair rp = SP may not be specified.
func (cpu *CPU) Pop(rp *memory.RegisterPair) {
	var stackPointer uint16
	cpu.SP.Read16(&stackPointer)
	rp.WriteLow(cpu.Memory[stackPointer])
	rp.WriteHigh(cpu.Memory[stackPointer+1])
	cpu.SP.Write16(stackPointer + 2)
}

// PopProcessorStatusWord implements the POP PSW instruction. The content of the memory
// location whose address is specified by the content of register SP is used to restore the
// condition flags. The content of the memory location whose address is one more than the
// content of register SP is moved to register A. The content of register SP is incremented by 2.
func (cpu *CPU) PopProcessorStatusWord() {
	var stackPointer uint16
	cpu.SP.Read16(&stackPointer)
	cpu.ALU.ApplyStatusWord(cpu.Memory[stackPointer])
	cpu.ALU.GetA().Write8(cpu.Memory[stackPointer+1])
	cpu.SP.Write16(stackPointer + 2)
}
