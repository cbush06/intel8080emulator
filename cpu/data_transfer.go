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

// MoveMemory implements MOV r, M. The content of the memory location, whose address is in registers Hand L, is moved to register r.
func (c *CPU) MoveMemory(r *memory.RegisterPair) {

}

// LoadRegisterPairImmediate implements LXI rp, data 16. Byte 3 of the instruction is moved into the high-order register (rh) of the
// register pair rp. Byte 2 of the in-struction is moved into the low-order register (rl) of the register pair rp.
func (c *CPU) LoadRegisterPairImmediate(r *memory.RegisterPair, byte2 uint8, byte3 uint8) {
	r.Low.Write8(byte2)
	r.High.Write8(byte3)
}
