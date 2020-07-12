package main

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

// StartCPU begins the CPU fetch-exeucte cycle and loads the specified program.
func StartCPU(program []byte) *CPUInterface {
	cpuInt := &CPUInterface{
		Interrupt: make(chan uint8),
	}

	cpu := new(cpu.CPU)
	cpu.Init()

	// Where the RUBBER MEETS THE ROAD
	go func() {
		for {
			// Handle Interrupts (button presses, vertical blank interrupts from screen, etc.)
			if cpu.EnableInterrupts {
				// Check for Interrupt; if set, execute interrupt instruction cycle
				select {
				case interruptCommand := <-cpuInt.Interrupt:
					fmt.Printf("INTERRUPT 0x%2X\n", interruptCommand)
					cpu.InterruptInstructionCycle(interruptCommand)
				default:
					// No interrupt...continue on
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return cpuInt
}
