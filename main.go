package main

//go:generate mockgen -destination=alu/mocks/condition_flags_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ConditionFlags
//go:generate mockgen -destination=alu/mocks/alu_mock.go -package=alu github.com/cbush06/intel8080emulator/alu ALU

import (
	"fmt"
	"time"

	"github.com/cbush06/intel8080emulator/cpu"
)

// CPUInterface is a struct comprising multiple go channels to allow IPC between
// this CPU, the machine using it, and any peripherals.
type CPUInterface struct {
	Interrupt chan uint8
}

func main() {}

// StartCPU begins the CPU fetch-execute cycle and loads the specified program.
func StartCPU(program []byte, programCounter uint16) *CPUInterface {
	cpuInt := &CPUInterface{
		Interrupt: make(chan uint8),
	}

	var mainCpu = new(cpu.CPU)
	mainCpu.Init()

	// Set ProgramCounter to execution starting point
	mainCpu.ProgramCounter = programCounter

	// Copy program into working memory
	copy(mainCpu.Memory, program)

	// Where the RUBBER MEETS THE ROAD
	go func() {
		for {
			// Handle Interrupts (button presses, vertical blank interrupts from screen, etc.)
			if mainCpu.EnableInterrupts {
				// Check for Interrupt; if set, execute interrupt instruction cycle
				select {
				case interruptCommand := <-cpuInt.Interrupt:
					fmt.Printf("INTERRUPT 0x%2X\n", interruptCommand)
					mainCpu.InterruptInstructionCycle(interruptCommand)
				default:
					// No interrupt...continue on
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return cpuInt
}
