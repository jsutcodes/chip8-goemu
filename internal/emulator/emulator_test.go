package emulator

import (
	"testing"
	"time"

	"github.com/jsutcodes/chip8-goemu/internal/cpu"
	"github.com/jsutcodes/chip8-goemu/internal/display"
	"github.com/jsutcodes/chip8-goemu/internal/input"
	"github.com/jsutcodes/chip8-goemu/internal/memory"
	"github.com/jsutcodes/chip8-goemu/internal/timer"
)

func TestRun(t *testing.T) {
	ram := memory.NewMemory()
	display := display.NewDisplay()
	keypad := input.NewKeypad()
	emu := &Emulator{
		RAM:     ram,
		CPU:     cpu.NewCPU(ram, display, keypad),
		Input:   keypad,
		Display: display,
		Timer:   timer.NewTimer(),
		running: true,
	}

	go emu.Run()

	// Allow some time for the emulator to run
	time.Sleep(100 * time.Millisecond)

	// Check if the fontset is loaded into memory
	for i, b := range chip8Fontset {
		if byteRead, err := ram.ReadByte(uint16(i)); err != nil || byteRead != b {
			byteRead, _ := ram.ReadByte(uint16(i))
			t.Errorf("Expected fontset byte %x at position %d, but got %x", b, i, byteRead)
		}
	}

	// Check if the emulator is stepping
	// This is a bit tricky to test directly, so we will check if the CPU cycle count has increased
	initialCycleCount := emu.CPU.CycleCount
	time.Sleep(100 * time.Millisecond)
	if emu.CPU.CycleCount == initialCycleCount {
		t.Errorf("Expected CPU cycle count to increase, but it did not")
	}

	// Stop the emulator
	emu.running = false
}
