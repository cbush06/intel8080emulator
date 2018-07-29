package cpu

import "github.com/cbush06/intel8080emulator/memory"

// Add implements the ADD instruction. Specifically, the content of register r is added to the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (cpu *CPU) Add(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AddImmediate(input)
}

// Sub implements the SUB instruction. Specifically, the content of register r is subtracted from the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (cpu *CPU) Sub(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.SubImmediate(input)
}

// AddWithCarry implements the ADC instruction. The content of register r and the content of the carry
// bit are added to the content of the accumulator. The result is placed in the accumulator. The AluFlags
// will be updated based on this operation's result.
func (cpu *CPU) AddWithCarry(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AddImmediateWithCarry(input)
}

// SubWithBorrow implements the SBB instruction.
func (cpu *CPU) SubWithBorrow(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.SubImmediateWithBorrow(input)
}

// IncrementRegister implements the INR instruction.
func (cpu *CPU) IncrementRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	r.Write8(cpu.ALU.Increment(input))
}
