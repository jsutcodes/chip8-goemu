package emulator

import (
	"fmt"
	"os"
	"time"

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

const CLOCK_SPEED = 700

var delayTimer byte
var soundTimer byte

type Emulator struct {
	// Add fields as needed
	RAM     *memory.Memory
	CPU     *cpu.CPU
	Input   *input.Keypad
	Display *display.Display
	Timer   *timer.Timer
	running bool
}

func (emu *Emulator) Test() {
	// Test IBM Logo
	emu.CPU.Cycle(true, emu.RAM)

}

func NewEmulator() *Emulator {
	ram := memory.NewMemory()
	display := display.NewDisplay()
	keypad := input.NewKeypad()
	return &Emulator{
		RAM:     ram,
		CPU:     cpu.NewCPU(ram, display, keypad),
		Input:   keypad,
		Display: display,
		Timer:   timer.NewTimer(),
		running: false,
	}
}

func (emu *Emulator) Run() {
	// Load the fontset into memory (0 - 80)
	fmt.Println(">>> Running Emulator")
	for i, b := range chip8Fontset {
		emu.RAM.WriteByte(uint16(i), b)
	}

	emu.RAM.PrintMemoryToFile("memory.dump")

	ticker := time.NewTicker(time.Second / CLOCK_SPEED)

	for emu.running {
		select {
		case <-ticker.C:
			fmt.Printf("Running")
			emu.Step()
		}
	}
}

func (emu *Emulator) Step() {
	// Run the emulator
	emu.CPU.Cycle(false, emu.RAM)
	emu.Timer.Update()
	emu.Display.Render()
}

func (emu *Emulator) LoadROM(path string) {
	// Implement the load ROM logic
	// convert string to file and load into bytes
	// pass the bytes to the LoadROM method in memory
	fmt.Println("Loading ROM: ", path)

	data, err := os.ReadFile(path)
	//fmt.Print(data)
	if err != nil {
		fmt.Printf("Failed to read ROM: %v", err)
	}
	emu.RAM.LoadROM(data)

	emu.running = true
}
