package cpu

import (
	"github.com/cbush06/intel8080emulator/memory"
	"testing"
)

func TestCPU_Input(t *testing.T) {
	var port uint8 = 255
	var data uint8 = 0xAB
	cpu := makeCPU(0, []uint8{uint8(CALL), 2, 1, 0, 0}, 5)
	cpu.AddressBus = *memory.NewRegisterPair(0, 0)
	cpu.DataBus = *memory.NewRegister(data)

	cpu.Input(port)

	var registerA uint8
	cpu.A.Read8(&registerA)
	if registerA != data {
		t.Errorf("Expected register A to contain %X but contained %X", data, registerA)
	}

	expectedAddressBus := uint16(port) << 8 | uint16(port)
	var addressBus uint16
	cpu.AddressBus.Read16(&addressBus)
	if expectedAddressBus != addressBus {
		t.Errorf("Expected address bus to contain %X but contained %X", expectedAddressBus, addressBus)
	}
}

func TestCPU_Output(t *testing.T) {
	var port uint8 = 255
	var data uint8 = 0xAB
	cpu := makeCPU(0, []uint8{uint8(CALL), 2, 1, 0, 0}, 5)
	cpu.AddressBus = *memory.NewRegisterPair(0, 0)
	cpu.DataBus = *memory.NewRegister(data)
	cpu.A.Write8(data)

	cpu.Output(port)

	var dataBus uint8
	cpu.DataBus.Read8(&dataBus)
	if dataBus != data {
		t.Errorf("Expected data bus to contain %X but contained %X", data, dataBus)
	}

	expectedAddressBus := uint16(port) << 8 | uint16(port)
	var addressBus uint16
	cpu.AddressBus.Read16(&addressBus)
	if expectedAddressBus != addressBus {
		t.Errorf("Expected address bus to contain %X but contained %X", expectedAddressBus, addressBus)
	}
}

func TestCPU_MoveRegister(t *testing.T) {
	var data1 uint8 = 0xAB
	var data2 uint8 = 0xCD
	register1 := memory.NewRegister(data1)
	register2 := memory.NewRegister(data2)

	cpu := &CPU{}
	cpu.MoveRegister(register1, register2)

	var r1Value uint8
	register1.Read8(&r1Value)

	var r2Value uint8
	register2.Read8(&r2Value)

	// Confirm r1 now holds the contents of r2
	if r1Value != data2 {
		t.Errorf("Expected register1 to contain %X but contained %X", data2, r1Value)
	}

	// Confirm r2 is unchanged
	if r2Value != data2 {
		t.Errorf("Expected register2 to contain %X but contained %X", data2, r2Value)
	}
}

func TestCPU_MoveFromMemory(t *testing.T) {
	var data uint8 = 0xAB

	cpu := makeCPU(0, []uint8{0, 0, data, 0}, 0)
	cpu.HL.Write16(0x0002) // Set HL to 0x0002 so MOV r, M writes to byte 2 of memory

	register := memory.NewRegister(0x00)

	cpu.MoveFromMemory(register)

	// Confirm register now holds byte 2 of memory's value
	var registerData uint8
	register.Read8(&registerData)
	if registerData != data {
		t.Errorf("Expected register to contain %X but contained %X", data, registerData)
	}
}

func TestCPU_MoveToMemory(t *testing.T) {
	var data uint8 = 0xAB

	cpu := makeCPU(0, []uint8{ 0, 0, 0, 0 }, 0)
	cpu.HL.Write16(0x0002)

	register := memory.NewRegister(data)

	cpu.MoveToMemory(register)

	// Confirm byte 2 of memory now hold's register's value
	if cpu.Memory[0x0002] != data {
		t.Errorf("Expected memory location 0x0002 to contain %X but contained %X", data, cpu.Memory[2])
	}
}

func TestCPU_MoveImmediate(t *testing.T) {
	cpu := &CPU{}

	var data uint8 = 0xAB
	register := memory.NewRegister(0x00)

	cpu.MoveImmediate(register, data)

	// Confirm register now holds the data
	var registerData uint8
	register.Read8(&registerData)
	if registerData != data {
		t.Errorf("Expected register to contain %X but contained %X", data, registerData)
	}
}

func TestCPU_MoveToMemoryImmediate(t *testing.T) {
	cpu := makeCPU(0, []uint8{ 0, 0, 0, 0 }, 0)
	cpu.HL.Write16(0x0002)

	var data uint8 = 0xAB

	cpu.MoveToMemoryImmediate(data)

	// Confirm value in register was moved to memory location 0x0002
	if cpu.Memory[2] != data {
		t.Errorf("Expected memory location 0x0002 to contain %X but contained %X", data, cpu.Memory[2])
	}
}

func TestCPU_LoadRegisterPairImmediate(t *testing.T) {
	cpu := &CPU{}
	registerPair := memory.NewRegisterPair(0x00, 0x00)

	var dataLow uint8 = 0xAB
	var dataHigh uint8 = 0xCD

	cpu.LoadRegisterPairImmediate(registerPair, dataLow, dataHigh)

	var actualDataLow uint8
	var actualDataHigh uint8

	registerPair.ReadLow(&actualDataLow)
	registerPair.ReadHigh(&actualDataHigh)

	// Confirm the register holds the correct values in the correct places
	if actualDataLow != dataLow {
		t.Errorf("Expected the low byte of register pair to contain %X but contained %X", dataLow, actualDataLow)
	}
	if actualDataHigh != dataHigh {
		t.Errorf("Expected the high byte of register pair to contain %X but contained %X", dataHigh, actualDataHigh)
	}
}

func TestCPU_LoadAccumulatorDirect(t *testing.T) {
	cpu := makeCPU(0, []uint8{ 0, 0, 0xAB, 0 }, 0)
	cpu.A = *memory.NewRegister(0x00)

	cpu.LoadAccumulatorDirect(0x02, 0x00)

	// Confirm that accumulator (register A) contains data from memory location 0x0002
	var registerAData uint8
	cpu.A.Read8(&registerAData)
	if registerAData != cpu.Memory[2] {
		t.Errorf("Expected accumulator (register A) to contain %X but contained %X", cpu.Memory[2], registerAData)
	}
}

func TestCPU_LoadAccumulatorIndirect(t *testing.T) {
	cpu := makeCPU(0, []uint8{ 0, 0, 0xAB, 0 }, 0)
	cpu.A = *memory.NewRegister(0x00)

	sourceRegisterPair := memory.NewRegisterPair(0x00, 0x02)

	cpu.LoadAccumulatorIndirect(sourceRegisterPair)

	// Confirm that accumulator (register A) contains data from memory location 0x0002
	var registerAData uint8
	cpu.A.Read8(&registerAData)
	if registerAData != cpu.Memory[2] {
		t.Errorf("Expected accumulator (register A) to contain %X but contained %X", cpu.Memory[2], registerAData)
	}
}

func TestCPU_StoreAccumulatorDirect(t *testing.T) {
	var data uint8 = 0xAB
	cpu := makeCPU(0, []uint8{0, 0, 0, 0}, 0)
	cpu.A = *memory.NewRegister(data)

	cpu.StoreAccumulatorDirect(0x02, 0x00)

	// Confirm that memory location 0x0002 holds 0xAB
	if cpu.Memory[2] != data {
		t.Errorf("Expected memory location 0x0002 to contain %X but contained %X", data, cpu.Memory[2])
	}
}

func TestCPU_StoreAccumulatorIndirect(t *testing.T) {
	var data uint8 = 0xAB
	cpu := makeCPU(0, []uint8{0, 0, 0, 0}, 0)
	cpu.A = *memory.NewRegister(data)

	sourceRegisterPair := memory.NewRegisterPair(0x00, 0x02)

	cpu.StoreAccumulatorIndirect(sourceRegisterPair)

	// Confirm that memory location 0x0002 holds 0xAB
	if cpu.Memory[2] != data {
		t.Errorf("Expected memory location 0x0002 to contain %X but contains %X", data, cpu.Memory[2])
	}
}

func TestCPU_LoadHandLDirect(t *testing.T) {
	var expectedLData uint8 = 0xAB
	var expectedHData uint8 = 0xCD

	cpu := &CPU{
		Memory: []uint8{expectedLData, expectedHData},
		H: memory.NewRegister(0x00),
		L: memory.NewRegister(0x00),
	}

	cpu.LoadHandLDirect(0x00, 0x00)

	var actualLData uint8
	var actualHData uint8
	cpu.L.Read8(&actualLData)
	cpu.H.Read8(&actualHData)

	if actualLData != expectedLData {
		t.Errorf("Expected %X but got %X", expectedLData, actualLData)
	}
	if actualHData != expectedHData {
		t.Errorf("Expected %X but got %X", expectedHData, actualLData)
	}
}

func TestCPU_StoreHandLDirect(t *testing.T) {
	var expectedMemoryLow uint8 = 0xAB
	var expectedMemoryHigh uint8 = 0xCD

	cpu := &CPU {
		Memory: []uint8{0, 0, 0, 0},
		H: memory.NewRegister(expectedMemoryHigh),
		L: memory.NewRegister(expectedMemoryLow),
	}

	cpu.StoreHandLDirect(0x00, 0x00)

	if expectedMemoryLow != cpu.Memory[0] {
		t.Errorf("Expected %X but got %X", expectedMemoryLow, cpu.Memory[0])
	}
	if expectedMemoryHigh != cpu.Memory[1] {
		t.Errorf("Expected %X but got %X", expectedMemoryHigh, cpu.Memory[1])
	}
}