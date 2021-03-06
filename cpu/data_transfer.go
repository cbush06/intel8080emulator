package cpu

import (
	"github.com/cbush06/intel8080emulator/memory"
)

// Input moves data that was placed on the eight bit bi-directional data bus by the specified
// port to register A (the accumulator). Presumably, the system state will have switched
// the WR pin to 1 (indicating a read operation).
func (cpu *CPU) Input() {
	var incomingData uint8
	port := cpu.Memory[cpu.ProgramCounter+1]

	// Write PORT selection to both high and low byte of AddressBus
	// SEE: Wikipedia's Intel 8080 I/O Scheme: https://en.wikipedia.org/wiki/Intel_8080#Input/output_scheme
	cpu.AddressBus.Write16((uint16(port) << 8) | uint16(port))

	// SEE: "I/O Addressing" in Intel 8080 System User's Manual (page 3-9)
	// SEE: "I/O Port Decoder" and Example #1 of that section in Intel 8080 System User's Manual (page 5-149)

	// Read DataBus values into register A (the accumulator)
	cpu.DataBus.Read8(&incomingData)
	cpu.A.Write8(incomingData)
	cpu.ProgramCounter += 2
}

// Output places the content of register A (the accumulator) on the eight-bit bi-directional data bus
// for transmission to the specified port. Presumably, the system state will have switched
// the WR pin to 0 (indicating a write operation).
func (cpu *CPU) Output() {
	var outgoingData uint8
	port := cpu.Memory[cpu.ProgramCounter+1]

	// Write PORT selection to both high and low byte of AddressBus
	// SEE: Wikipedia's Intel 8080 I/O Scheme: https://en.wikipedia.org/wiki/Intel_8080#Input/output_scheme
	cpu.AddressBus.Write16((uint16(port) << 8) | uint16(port))

	// SEE: "I/O Addressing" in Intel 8080 System User's Manual (page 3-9)
	// SEE: "I/O Port Decoder" and Example #1 of that section in Intel 8080 System User's Manual (page 5-149)

	// Write register A (the accumulator) values into DataBus
	cpu.A.Read8(&outgoingData)
	cpu.DataBus.Write8(outgoingData)
	cpu.ProgramCounter += 2
}

// MoveRegister implements MOV r1, r2. The content of register r2 is moved to register r1.
func (cpu *CPU) MoveRegister(r1 *memory.Register, r2 *memory.Register) {
	var data uint8
	r2.Read8(&data)
	r1.Write8(data)
	cpu.ProgramCounter += 1
}

// MoveFromMemory implements MOV r, M. The content of the memory location, whose address is in registers H and L, is moved to register r.
func (cpu *CPU) MoveFromMemory(r *memory.Register) {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	r.Write8(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// MoveToMemory implements MOV M, r. The content of register r is moved to the memory location whose address is in registers H and L.
func (cpu *CPU) MoveToMemory(r *memory.Register) {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	r.Read8(&cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// MoveImmediate implements MOV r, data. The data argument is moved to register r.
func (cpu *CPU) MoveImmediate(r *memory.Register) {
	r.Write8(cpu.Memory[cpu.ProgramCounter+1])
	cpu.ProgramCounter += 2
}

// MoveToMemoryImmediate implements MVI M, data. The data argument is moved to
// the memory location whose address is in registers H and L.
func (cpu *CPU) MoveToMemoryImmediate() {
	var memoryAddress uint16
	cpu.HL.Read16(&memoryAddress)
	cpu.Memory[memoryAddress] = cpu.Memory[cpu.ProgramCounter+1]
	cpu.ProgramCounter += 2
}

// LoadRegisterPairImmediate implements LXI rp, data 16. Byte 3 of the instruction is moved into the high-order register (rh) of the
// register pair rp. Byte 2 of the instruction is moved into the low-order register (rl) of the register pair rp.
func (cpu *CPU) LoadRegisterPairImmediate(rp *memory.RegisterPair) {
	rp.Low.Write8(cpu.Memory[cpu.ProgramCounter+1])
	rp.High.Write8(cpu.Memory[cpu.ProgramCounter+2])
	cpu.ProgramCounter += 3
}

// LoadAccumulatorDirect implements LDA addr. The content of the memory location, whose address
// is specified in byte 2 and byte 3 of the instruction, is moved to register A.
func (cpu *CPU) LoadAccumulatorDirect() {
	var memoryAddress uint16
	memoryAddress = (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
	cpu.A.Write8(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 3
}

// StoreAccumulatorDirect implements STA addr. The content of the accumulator is moved to the
// memory location whose address is specified in byte 2 and byte 3 of the instruction.
func (cpu *CPU) StoreAccumulatorDirect() {
	var memoryAddress uint16
	memoryAddress = (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
	cpu.A.Read8(&cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 3
}

// LoadHandLDirect implements LHLD addr. (L) <- ((byte 3)(byte 2)); (H) <- ((byte 3) (byte 2) + 1).
// The content of the memory location, whose address is specified in byte 2 and byte 3 of the instruction, is
// moved to register L. The content of the memory location at the succeeding address is moved to register H.
func (cpu *CPU) LoadHandLDirect() {
	var memoryAddress uint16
	memoryAddress = (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
	cpu.L.Write8(cpu.Memory[memoryAddress])
	cpu.H.Write8(cpu.Memory[memoryAddress+1])
	cpu.ProgramCounter += 3
}

// StoreHandLDirect implements SHLD addr. ((byte 3) (byte 2)) <- (L); ((byte 3)(byte 2) + 1) <- (H).
// The content of register L is moved to the memory location whose address is specified in byte 2 and byte
// 3. The content of register H is moved to the succeeding memory location.
func (cpu *CPU) StoreHandLDirect() {
	var memoryAddress uint16
	memoryAddress = (uint16(cpu.Memory[cpu.ProgramCounter+2]) << 8) | uint16(cpu.Memory[cpu.ProgramCounter+1])
	cpu.L.Read8(&cpu.Memory[memoryAddress])
	cpu.H.Read8(&cpu.Memory[memoryAddress+1])
	cpu.ProgramCounter += 3
}

// LoadAccumulatorIndirect implements LDAX rp. The content of the memory location, whose address
// is in the register pair rp, is moved to register A. Note: only register pairs rp=B (registers B and C·) or rp=D
// (registers D and E) may be specified.
func (cpu *CPU) LoadAccumulatorIndirect(rp *memory.RegisterPair) {
	var memoryAddress uint16
	rp.Read16(&memoryAddress)
	cpu.A.Write8(cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}

// StoreAccumulatorIndirect implements STAX rp. The content of register A is moved to the memory location whose address is in the
// register pair rp. Note: only register pairs rp=B (registers B and C) or rp=D (registers D and E) may be specified.
func (cpu *CPU) StoreAccumulatorIndirect(rp *memory.RegisterPair) {
	var memoryAddress uint16
	rp.Read16(&memoryAddress)
	cpu.A.Read8(&cpu.Memory[memoryAddress])
	cpu.ProgramCounter += 1
}


// MoveHandLtoPC implements the PCHL instruction. The content of register H is moved to the high-order eight bits
// of register PC. The content of register l is moved to the low-order eight bits of register PC.
// (PCH) <- (H)
// (PCl) <- (l)
func (cpu *CPU) MoveHandLtoPC() {
	var hl uint16
	cpu.HL.Read16(&hl)
	cpu.ProgramCounter = hl
}

// ExchangeHandLWithDAndE implements the XCHG instruction.
func (cpu *CPU) ExchangeHandLWithDAndE() {
	var hl uint16
	var de uint16
	cpu.HL.Read16(&hl)
	cpu.DE.Read16(&de)

	cpu.HL.Write16(de)
	cpu.DE.Write16(hl)

	cpu.ProgramCounter += 1
}
