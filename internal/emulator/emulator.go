package emulator

import (
	"log"
	"os"

	"github.com/jsutcodes/chip8-goemu/internal/cpu"
	"github.com/jsutcodes/chip8-goemu/internal/memory"
	"github.com/jsutcodes/chip8-goemu/internal/timer"

	"github.com/jsutcodes/chip8-goemu/internal/display"
	"github.com/jsutcodes/chip8-goemu/internal/input"
)

var chip8Fontset = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

var delayTimer byte
var soundTimer byte

type Emulator struct {
	// Add fields as needed
	RAM     *memory.Memory
	CPU     *cpu.CPU
	Input   *input.Keypad
	Display *display.Display
	Timer   *timer.Timer
}

func NewEmulator() *Emulator {
	ram := memory.NewMemory()
	return &Emulator{
		RAM:     ram,
		CPU:     cpu.NewCPU(ram),
		Input:   input.NewKeypad(),
		Display: display.NewDisplay(),
		Timer:   timer.NewTimer(),
	}
}

func (emu *Emulator) Run() {
	// Implement the start logic

	// Load the fontset into memory (0 -80)
	for i, b := range chip8Fontset {
		emu.RAM.WriteByte(uint16(i), b)
	}
}

func (emu *Emulator) LoadROM(path string) {
	// Implement the load ROM logic
	// convert string to file and load into bytes
	// pass the bytes to the LoadROM method in memory
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read ROM: %v", err)
	}
	emu.RAM.LoadROM(data)
}
