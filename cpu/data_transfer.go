package cpu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// Input moves data that was placed on the eight bit bi-directional data bus by the specified
// port to register A (the accumulator). Presumably, the system state will have switched
// the WR pin to 1 (indicating a read operation).
func (c *CPU) Input(port uint8) {
	var incomingData uint8

	// Write PORT selection to both high and low byte of AddressBus
	// SEE: Wikipedia's Intel 8080 I/O Scheme: https://en.wikipedia.org/wiki/Intel_8080#Input/output_scheme
	c.AddressBus.Write16((uint16(port) << 8) | uint16(port))

	// SEE: "I/O Addressing" in Intel 8080 System User's Manual (page 3-9)
	// SEE: "I/O Port Decoder" and Example #1 of that section in Intel 8080 System User's Manual (page 5-149)

	// Read DataBus values into register A (the accumulator)
	c.DataBus.Read8(&incomingData)
	c.A.Write8(incomingData)
}

// Output places the content of register A (the accumulator) on the eight-bit bi-directional data bus
// for transmission to the specified port. Presumably, the system state will have switched
// the WR pin to 0 (indicating a write operation).
func (c *CPU) Output(port uint8) {
	var outgoingData uint8

	// Write PORT selection to both high and low byte of AddressBus
	// SEE: Wikipedia's Intel 8080 I/O Scheme: https://en.wikipedia.org/wiki/Intel_8080#Input/output_scheme
	c.AddressBus.Write16((uint16(port) << 8) | uint16(port))

	// SEE: "I/O Addressing" in Intel 8080 System User's Manual (page 3-9)
	// SEE: "I/O Port Decoder" and Example #1 of that section in Intel 8080 System User's Manual (page 5-149)

	// Write register A (the accumulator) values into DataBus
	c.A.Read8(&outgoingData)
	c.DataBus.Write8(outgoingData)
}

// MoveRegister implements MOV r1, r2. The content of register r2 is moved to register r1.
func (c *CPU) MoveRegister(r1 *memory.Register, r2 *memory.Register) {
	var data uint8
	r2.Read8(&data)
	r1.Write8(data)
	c.ProgramCounter += 1
}

// MoveFromMemory implements MOV r, M. The content of the memory location, whose address is in registers H and L, is moved to register r.
func (c *CPU) MoveFromMemory(r *memory.Register) {
	var memoryAddress uint16
	c.HL.Read16(&memoryAddress)
	r.Write8(c.Memory[memoryAddress])
	c.ProgramCounter += 1
}

// MoveToMemory implements MOV M, r. The content of register r is moved to the memory location whose address is in registers H and L.
func (c *CPU) MoveToMemory(r *memory.Register) {
	var memoryAddress uint16
	c.HL.Read16(&memoryAddress)
	r.Read8(&c.Memory[memoryAddress])
	c.ProgramCounter += 1
}

// MoveImmediate implements MOV r, data. The data argument is moved to register r.
func (c *CPU) MoveImmediate(r *memory.Register, data uint8) {
	r.Write8(data)
	c.ProgramCounter += 2
}

// MoveToMemoryImmediate implements MVI M, data. The data argument is moved to
// the memory location whose address is in registers H and L.
func (c *CPU) MoveToMemoryImmediate(data uint8) {
	var memoryAddress uint16
	c.HL.Read16(&memoryAddress)
	c.Memory[memoryAddress] = data
	c.ProgramCounter += 2
}

// LoadRegisterPairImmediate implements LXI rp, data 16. Byte 3 of the instruction is moved into the high-order register (rh) of the
// register pair rp. Byte 2 of the instruction is moved into the low-order register (rl) of the register pair rp.
func (c *CPU) LoadRegisterPairImmediate(rp *memory.RegisterPair, byte2 uint8, byte3 uint8) {
	rp.Low.Write8(byte2)
	rp.High.Write8(byte3)
	c.ProgramCounter += 3
}

// LoadAccumulatorDirect implements LDA addr. The content of the memory location, whose address
// is specified in byte 2 and byte 3 of the instruction, is moved to register A.
func (c *CPU) LoadAccumulatorDirect(byte2 uint8, byte3 uint8) {
	var memoryAddress uint16
	memoryAddress = (uint16(byte3) << 8) | uint16(byte2)
	c.A.Write8(c.Memory[memoryAddress])
	c.ProgramCounter += 3
}

// StoreAccumulatorDirect implements STA addr. The content of the accumulator is moved to the
// memory location whose address is specified in byte 2 and byte 3 of the instruction.
func (c *CPU) StoreAccumulatorDirect(byte2 uint8, byte3 uint8) {
	var memoryAddress uint16
	memoryAddress = (uint16(byte3) << 8) | uint16(byte2)
	c.A.Read8(&c.Memory[memoryAddress])
	c.ProgramCounter += 3
}

// LoadHandLDirect implements LHLD addr. (L) <- ((byte 3)(byte 2)); (H) <- ((byte 3) (byte 2) + 1).
// The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is
// moved to register L. The content of the memory location at the succeeding address is moved to register H.
func (c *CPU) LoadHandLDirect(byte2 uint8, byte3 uint8) {
	var memoryAddress uint16
	memoryAddress = (uint16(byte3) << 8) | uint16(byte2)
	c.L.Write8(c.Memory[memoryAddress])
	c.H.Write8(c.Memory[memoryAddress+1])
	c.ProgramCounter += 3
}

// StoreHandLDirect implements SHLD addr. ((byte 3) (byte 2)) <- (L); ((byte 3)(byte 2) + 1) <- (H).
// The content of register L is moved to the memory location whose address is specified in byte 2 and byte
// 3. The content of register H is moved to the succeeding memory location.
func (c *CPU) StoreHandLDirect(byte2 uint8, byte3 uint8) {
	var memoryAddress uint16
	memoryAddress = (uint16(byte3) << 8) | uint16(byte2)
	c.L.Read8(&c.Memory[memoryAddress])
	c.H.Read8(&c.Memory[memoryAddress+1])
	c.ProgramCounter += 3
}

// LoadAccumulatorIndirect implements LDAX rp. The content of the memory location, whose address
// is in the register pair rp, is moved to register A. Note: only register pairs rp=B (registers B and C·) or rp=D
// (registers D and E) may be specified.
func (c *CPU) LoadAccumulatorIndirect(rp *memory.RegisterPair) {
	var memoryAddress uint16
	rp.Read16(&memoryAddress)
	c.A.Write8(c.Memory[memoryAddress])
	c.ProgramCounter += 1
}

// StoreAccumulatorIndirect implements STAX rp. The content of register A is moved to the memory location whose address is in the
// register pair rp. Note: only register pairs rp=B (registers B and C) or rp=D (registers D and E) may be specified.
func (c *CPU) StoreAccumulatorIndirect(rp *memory.RegisterPair) {
	var memoryAddress uint16
	rp.Read16(&memoryAddress)
	c.A.Read8(&c.Memory[memoryAddress])
	c.ProgramCounter += 1
}

// TODO: XCHG -- Exchange H and L with D and E