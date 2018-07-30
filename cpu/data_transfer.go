package cpu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// MoveRegister implements MOV r1, r2. The content of register r2 is moved to register r1.
func (c *CPU) MoveRegister(r1 *memory.RegisterPair, r2 *memory.RegisterPair) {
	var buf = make([]byte, 2)
	r2.Read(buf)
	r1.Write(buf)
}

// MoveMemory implements MOV r, M. The content of the memory location, whose address is in registers H and L, is moved to register r.
func (c *CPU) MoveMemory(r *memory.Register) {
	var memoryAddress uint16
	c.HL.Read16(&memoryAddress)
	r.Write8(c.Memory[memoryAddress])
}

// MoveImmediate implements MOV r, data. The data argument is moved to register r.
func (c *CPU) MoveImmediate(r *memory.Register, data uint8) {
	r.Write8(data)
}

// MoveToMemoryImmediate implements MVI M, data. The data argument is moved to
// the memory location whose address is in registers H and L.
func (c *CPU) MoveToMemoryImmediate(data uint8) {
	var memoryAddress uint16
	c.HL.Read16(&memoryAddress)
	c.Memory[memoryAddress] = data
}

// LoadRegisterPairImmediate implements LXI rp, data 16. Byte 3 of the instruction is moved into the high-order register (rh) of the
// register pair rp. Byte 2 of the in-struction is moved into the low-order register (rl) of the register pair rp.
func (c *CPU) LoadRegisterPairImmediate(r *memory.RegisterPair, byte2 uint8, byte3 uint8) {
	r.Low.Write8(byte2)
	r.High.Write8(byte3)
}

// StoreAccumulatorIndirect implements STAX rp. The content of register A is moved to the memory location whose address is in the
// register pair rp. Note: only register pairs rp=B (registers B and C) or rp=D (registers D and E) may be specified.
func (c *CPU) StoreAccumulatorIndirect(r *memory.RegisterPair) {
	var memoryAddress uint16
	r.Read16(&memoryAddress)
	c.A.Read8(&c.Memory[memoryAddress])
}
