package cpu

import "github.com/cbush06/intel8080emulator/memory"

// AddRegister implements the ADD instruction. Specifically, the content of register r is added to the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (cpu *CPU) AddRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AddImmediate(input)
	cpu.ProgramCounter += 1
}

// AddImmediate implements the ADI data instruction. The content of the second byte of the instruction is added
// to the content of the accumulator. The result is placed in the accumulator.
func (cpu *CPU) AddImmediate() {
	cpu.ALU.AddImmediate(cpu.Memory[cpu.ProgramCounter+1])
	cpu.ProgramCounter += 2
}

// AddImmediateWithCarry implements the ACI data instruction. The content of the second byte of the instruction and
// the content of the CY flag are added to the contents of the accumulator. The result is placed in the accumulator.
func (cpu *CPU) AddImmediateWithCarry() {
	cpu.ALU.AddImmediateWithCarry(cpu.Memory[cpu.ProgramCounter+1])
	cpu.ProgramCounter += 2
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
	cpu.ProgramCounter += 1
}

// SubtractRegister implements the SUB instruction. Specifically, the content of register r is subtracted from the content of the
// accumulator. The result is placed in the accumulator. The AluFlags will be updated based on this
// operation's result.
func (cpu *CPU) SubtractRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.SubImmediate(input)
	cpu.ProgramCounter += 1
}

// SubtractImmediate implements the SUI instruction. (A) <- (A) - (byte 2), The content of the second byte of the
// instruction is subtracted from the content of the accumulator. The result is placed in the accumulator.
func (cpu *CPU) SubtractImmediate() {
	cpu.ALU.SubImmediate(cpu.Memory[cpu.ProgramCounter+1])
	cpu.ProgramCounter += 2
}

// SubtractImmediateWithBorrow implements the SBI instruction. (A) <- (A) - (byte 2) - (CY). The contents of the second
// byte of the instruction and the contents of the CY flag are both subtracted from the accumulator. The result is
// placed in the accumulator.
func (cpu *CPU) SubtractImmediateWithBorrow() {
	cpu.ALU.SubImmediateWithBorrow(cpu.Memory[cpu.ProgramCounter+1])
	cpu.ProgramCounter += 2
}

// AddRegisterWithCarry implements the ADC instruction. The content of register r and the content of the carry
// bit are added to the content of the accumulator. The result is placed in the accumulator. The AluFlags
// will be updated based on this operation's result.
func (cpu *CPU) AddRegisterWithCarry(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AddImmediateWithCarry(input)
	cpu.ProgramCounter += 1
}

// SubtractRegisterWithBorrow implements the SBB instruction.
func (cpu *CPU) SubtractRegisterWithBorrow(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.SubImmediateWithBorrow(input)
	cpu.ProgramCounter += 1
}

// IncrementRegister implements the INR instruction. The content of register r is incremented by one.
// Note: All condition flags except CY are affected.
func (cpu *CPU) IncrementRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	r.Write8(cpu.ALU.Increment(input))
	cpu.ProgramCounter += 1
}

// IncrementRegisterPair implements the INX instruction. The content of the register pair is incremented by
// one. Note: No condition flags are affected.
func (cpu *CPU) IncrementRegisterPair(rp *memory.RegisterPair) {
	var input uint16
	rp.Read16(&input)
	rp.Write16(cpu.ALU.IncrementDouble(input))
	cpu.ProgramCounter += 1
}

// IncrementMemory implements the INR M instruction. The content of the memory location whose address
// is contained in the H and L registers is incremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) IncrementMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.ALU.Increment(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// DecrementRegister implements the DCR instruction. The content of register r is decremented by one.
// Note: All condition flag~ except CY are affected.
func (cpu *CPU) DecrementRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	r.Write8(cpu.ALU.Decrement(input))
	cpu.ProgramCounter += 1
}

// DecrementRegisterPair implements the DCX instruction. The content of the register pair is decremented by
// one. Note: No condition flags are affected.
func (cpu *CPU) DecrementRegisterPair(rp *memory.RegisterPair) {
	var input uint16
	rp.Read16(&input)
	rp.Write16(cpu.ALU.DecrementDouble(input))
	cpu.ProgramCounter += 1
}

// DecrementMemory implements the DCRM instruction. The content of the memory location whose address is
// contained in the H and L registers is decremented by one. Note: All condition flags except CY are affected.
func (cpu *CPU) DecrementMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.ALU.Decrement(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// RotateRight implements the RRC instruction. The content of the accumulator is rotated right one
// position. The high order bit and the CY flag are both set to the value shifted out of the low order bit
// position. Only the CY flag is affected.
func (cpu *CPU) RotateRight() {
	cpu.ALU.RotateRight()
	cpu.ProgramCounter += 1
}

// RotateRightThroughCarry implements the RAR instruction. The content of the accumulator is rotated right one
// position through the CY flag. The high order bit is set to the CY flag and the CY flag is set to the value
// shifted out of the low order bit. Only the CY flag is affected.
func (cpu *CPU) RotateRightThroughCarry() {
	cpu.ALU.RotateRightThroughCarry()
	cpu.ProgramCounter += 1
}

// RotateLeft implements the RLC instruction. The content of the accumulator is rotated left one
// position. The low order bit and the CY flag are both set to the value shifted out of the high order bit
// position. Only the CY flag is affected.
func (cpu *CPU) RotateLeft() {
	cpu.ALU.RotateLeft()
	cpu.ProgramCounter += 1
}

// RotateLeftThroughCarry implements the RAL instruction. The content of the accumulator is rotated left one
// position through the CY flag. The low order bit is set equal to the CY flag and the CY flag is set to the
// value shifted out of the high order bit. Only the CY flag is affected.
func (cpu *CPU) RotateLeftThroughCarry() {
	cpu.ALU.RotateLeftThroughCarry()
	cpu.ProgramCounter += 1
}

// AndImmediate implements the ANI data instruction. (A) <- (A) /\ (byte 2). The content of the second byte of the
// instruction is logically anded with the contents of the accumulator. The result is placed in the accumulator.
// The CY and AC flags are cleared.
func (cpu *CPU) AndImmediate() {
	operand := cpu.Memory[cpu.ProgramCounter+1]
	cpu.ALU.AndAccumulator(operand)
	cpu.ALU.ClearAuxiliaryCarry()
	cpu.ProgramCounter += 2
}

// AndRegister implements the ANA r instruction. The content of register r is logically anded with the content
// of the accumulator. The result is placed in the accumulator. The CY flag is cleared.
func (cpu *CPU) AndRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.AndAccumulator(input)
	cpu.ProgramCounter += 1
}

// AndMemory implements the ANA M instruction. The contents of the memory location whose address is contained
// in the Hand L registers is logically anded with the content of the accumulator. The result is placed in the
// accumulator. The CY flag is cleared.
func (cpu *CPU) AndMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.ALU.AndAccumulator(cpu.Memory[memoryAddress])
	cpu.ProgramCounter = +1
}

// OrImmediate implements the ORI data instruction. (A) <- (A) V (byte 2). The content of the second byte of the
// instruction is inclusive-OR'd with the content of the accumulator. The result is placed in the accumulator.
// The CY and AC flags are cleared.
func (cpu *CPU) OrImmediate() {
	operand := cpu.Memory[cpu.ProgramCounter+1]
	cpu.ALU.OrAccumulator(operand)
	cpu.ProgramCounter += 2
}

// OrRegister implements the ORA r instruction. (A) <- (A) V (r). The content of register r is inclusive-OR'd with the
// content of the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
func (cpu *CPU) OrRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.OrAccumulator(input)
	cpu.ProgramCounter += 1
}

// XOrImmediate implements the XRI data instruction. (A) <- (A) ^ (byte 2). The content of the second byte of the
// instruction is exclusive-O R'd with the content of the accumulator. The result is placed in the accumulator. The
// CY and AC flags are cleared.
func (cpu *CPU) XOrImmediate() {
	operand := cpu.Memory[cpu.ProgramCounter+1]
	cpu.ALU.XOrAccumulator(operand)
	cpu.ProgramCounter += 2
}

// XOrRegister implements the XRA r instruction. The content of register r is exclusive-or'd with the
// content of the accumulator. The result is placed in the accumulator. The CY and AC flags are cleared.
func (cpu *CPU) XOrRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.XOrAccumulator(input)
	cpu.ProgramCounter += 1
}

// XOrMemory implements the XRA M instruction. The content of the memory location whose address is contained
// in the Hand L registers is exclusive-OR'd with the content of the accumulator. The result is placed in the
// accumulator. The CY and AC flags are cleared.
func (cpu *CPU) XOrMemory() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.ALU.XOrAccumulator(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// DecimalAccumulatorAdjust implements the DAA instruction. The eight-bit number in the accumulator is adjusted
// to form two four-bit Binary-Coded-Decimal digits by
// the following process:
//
//		1. If the value of the least significant 4 bits of the
//		   accumulator is greater than 9 or if the AC flag
//	       is set, 6 is added to the accumulator.
//
//		2. If the value of the most significant 4 bits of the
//		   accumulator is now greater than 9, or if the CY
//		   flag is set, 6 is added to the most significant 4
//		   bits of the accumulator.
//
// NOTE: All flags are affected.
func (cpu *CPU) DecimalAccumulatorAdjust() {
	cpu.ALU.DecimalAdjustAccumulator()
	cpu.ProgramCounter += 1
}

// ComplementAccumulator implements the CMA instruction. The contents of the accumulator are complemented
// (zero bits become 1, one bits become 0). No flags are affected.
func (cpu *CPU) ComplementAccumulator() {
	cpu.ALU.ComplementAccumulator()
	cpu.ProgramCounter += 1
}

// SetCarry implements the STC instruction. (CY) <- 1. The CY flag is set to 1. No other flags are affected.
func (cpu *CPU) SetCarry() {
	cpu.ALU.SetCarry()
	cpu.ProgramCounter += 1
}

// ComplementCarry implements the CMC instruction. (CY) <- !(CY). The CY flag is complemented. No other flags are
// affected.
func (cpu *CPU) ComplementCarry() {
	if cpu.ALU.IsCarry() {
		cpu.ALU.ClearCarry()
	} else {
		cpu.ALU.SetCarry()
	}
}

// CompareRegister implements the CMP r instruction. (A) - (r). The content of register r is subtracted from the
// accumulator. The accumulator remains unchanged. The condition flags are set as a result of the subtraction.
// The Z flag is set to 1 if (A) = (r). The CY flag is set to 1 if (A) < (r).
func (cpu *CPU) CompareRegister(r *memory.Register) {
	var input uint8
	r.Read8(&input)
	cpu.ALU.CompareAccumulator(input)
	cpu.ProgramCounter += 1
}

// CompareImmediate implements the CPI data instruction. (A) - (byte 2). The content of the second byte of the
// instruction is subtracted from the accumulator. The condition flags are set by the result of the subtraction.
// The Z flag is set to 1 if (A) = (byte 2). The CY flag is set to 1 if (A) < (byte 2).
func (cpu *CPU) CompareImmediate() {
	operand := cpu.Memory[cpu.ProgramCounter+1]
	cpu.ALU.CompareAccumulator(operand)
	cpu.ProgramCounter += 2
}
