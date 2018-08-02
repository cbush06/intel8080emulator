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

// DoubleAdd implements the DAD instruction. Specifically, the content of the register pair rp is added to the
// content of the register pair Hand L. The result is placed in the register pair Hand L. Note: Only the
// CY flag is affected. It is set if there is a carry out of the double precision add; otherwise it is reset.
func (cpu *CPU) DoubleAdd(r *memory.RegisterPair) {
	var hlValue uint16
	var rpValue uint16

	cpu.HL.Read16(&hlValue)
	r.Read16(&rpValue)

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
func (cpu *CPU) IncrementRegisterPair(r *memory.RegisterPair) {
	var input uint16
	r.Read16(&input)
	r.Write16(cpu.ALU.IncrementDouble(input))
}

// IncrementMemory implements the INR M instruction. The content of the memory location whose address
// is contained in the Hand L registers is incremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) IncrementMemory(r *memory.Register) {
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
func (cpu *CPU) DecrementRegisterPair(r *memory.RegisterPair) {
	var input uint16
	r.Read16(&input)
	r.Write16(cpu.ALU.DecrementDouble(input))
}

// DecrementMemory implements the DCRM instruction. The content of the memory location whose address is
// contained in the Hand L registers is decremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) DecrementMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.ALU.Decrement(cpu.Memory[memoryAddress])
}

// RotateRight implements the RRC instruction. The content of the accumu lator is rotated right one
// position. The high order bit and the CY flag are both set to the value shifted out of the low order bit
// position. Only the CY flag is affected.
func (cpu *CPU) RotateRight() {
	cpu.ALU.RotateRight()
}
