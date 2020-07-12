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

// AddImmediate implements the ADI data instruction. The content of the second byte of the instruction is added
// to the content of the accumulator. The result is placed in the accumulator.
func (cpu *CPU) AddImmediate() {
	cpu.ALU.AddImmediate(cpu.Memory[cpu.ProgramCounter+1])
}

// DoubleAdd implements the DAD instruction. Specifically, the content of the register pair rp is added to the
// content of the register pair Hand L. The result is placed in the register pair H and L. Note: Only the
// CY flag is affected. It is set if there is a carry out of the double precision add; otherwise it is reset.
func (cpu *CPU) DoubleAdd(rp *memory.RegisterPair) {
	var hlValue uint16
	var rpValue uint16

	cpu.HL.Read16(&hlValue)
	rp.Read16(&rpValue)

	cpu.HL.Write16(cpu.ALU.DoubleAdd(hlValue, rpValue))
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

// IncrementRegister implements the INR instruction. The content of register r is incremented by one.
// Note: All condition flags except CY are affected.
func (cpu *CPU) IncrementRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	r.Write8(cpu.ALU.Increment(input))
}

// IncrementRegisterPair implements the INX instruction. The content of the register pair is incremented by
// one. Note: No condition flags are affected.
func (cpu *CPU) IncrementRegisterPair(rp *memory.RegisterPair) {
	var input uint16
	rp.Read16(&input)
	rp.Write16(cpu.ALU.IncrementDouble(input))
}

// IncrementMemory implements the INR M instruction. The content of the memory location whose address
// is contained in the H and L registers is incremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) IncrementMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.ALU.Increment(cpu.Memory[memoryAddress])
}

// DecrementRegister implements the DCR instruction. The content of register r is decremented by one.
// Note: All condition flag~ except CY are affected.
func (cpu *CPU) DecrementRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	r.Write8(cpu.ALU.Decrement(input))
}

// DecrementRegisterPair implements the DCX instruction. The content of the register pair is decremented by
// one. Note: No condition flags are affected.
func (cpu *CPU) DecrementRegisterPair(rp *memory.RegisterPair) {
	var input uint16
	rp.Read16(&input)
	rp.Write16(cpu.ALU.DecrementDouble(input))
}

// DecrementMemory implements the DCRM instruction. The content of the memory location whose address is
// contained in the H and L registers is decremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) DecrementMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.ALU.Decrement(cpu.Memory[memoryAddress])
}

// RotateRight implements the RRC instruction. The content of the accumulator is rotated right one
// position. The high order bit and the CY flag are both set to the value shifted out of the low order bit
// position. Only the CY flag is affected.
func (cpu *CPU) RotateRight() {
	cpu.ALU.RotateRight()
}

// AndRegister implements the ANA r instruction. The content of register r is logically anded with the content
// of the accumulator. The result is placed in the accumulator. The CY flag is cleared.
func (cpu *CPU) AndRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AndAccumulator(input)
}

// AndMemory implements the ANA M instruction. The contents of the memory location whose address is contained
// in the Hand L registers is logically anded with the content of the accumulator. The result is placed in the
// accumulator. The CY flag is cleared.
func (cpu *CPU) AndMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.ALU.AndAccumulator(cpu.Memory[memoryAddress])
}

// XOrRegister implements the XRA r instruction. The content of register r is exclusive-or'd with the
// content of the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
func (cpu *CPU) XOrRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.XOrAccumulator(input)
}

// XOrMemory implements the XRA M instruction. The content of the memory location whose address is contained
// in the Hand L registers is exclusive-OR'd with the content of the accumulator. The result is placed in the
// accumulator. The CY and AC flags are cleared.
func (cpu *CPU) XOrMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.ALU.XOrAccumulator(cpu.Memory[memoryAddress])
}
