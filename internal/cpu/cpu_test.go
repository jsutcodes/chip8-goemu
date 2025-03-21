package cpu

import (
	"testing"

	"github.com/jsutcodes/chip8-goemu/internal/display"
	"github.com/jsutcodes/chip8-goemu/internal/input"
	"github.com/jsutcodes/chip8-goemu/internal/memory"
)

func setup() *CPU {
	ram := memory.NewMemory()
	display := display.NewDisplay()
	input := input.NewKeypad()
	return NewCPU(ram, display, input)
}

func TestOpcode00E0(t *testing.T) {
	cpu := setup()
	cpu.Display.SetPixel(0, 0, true)
	cpu.decodeAndExecute(0x00E0)
	if cpu.Display.IsPixelOn(0, 0) {
		t.Errorf("Expected display to be cleared")
	}
}

func TestOpcode1NNN(t *testing.T) {
	cpu := setup()
	cpu.decodeAndExecute(0x1234)
	if cpu.PC != 0x0234 {
		t.Errorf("Expected PC to be 0x0234, got 0x%X", cpu.PC)
	}
}

func TestOpcode6XNN(t *testing.T) {
	cpu := setup()
	cpu.decodeAndExecute(0x60FF)
	if cpu.V[0] != 0xFF {
		t.Errorf("Expected V0 to be 0xFF, got 0x%X", cpu.V[0])
	}
}

func TestOpcode7XNN(t *testing.T) {
	cpu := setup()
	cpu.V[0] = 0x01
	cpu.decodeAndExecute(0x7001)
	if cpu.V[0] != 0x02 {
		t.Errorf("Expected V0 to be 0x02, got 0x%X", cpu.V[0])
	}
}

func TestOpcodeANNN(t *testing.T) {
	cpu := setup()
	cpu.decodeAndExecute(0xA123)
	if cpu.I != 0x0123 {
		t.Errorf("Expected I to be 0x0123, got 0x%X", cpu.I)
	}
}

// TODO: Fix this, test fails
func TestOpcodeDXYN(t *testing.T) {
	cpu := setup()
	cpu.V[0] = 0
	cpu.V[1] = 0
	cpu.I = 0x300
	cpu.RAM.WriteByte(0x300, 0xFF)
	cpu.decodeAndExecute(0xD011)
	if !cpu.Display.IsPixelOn(0, 0) {
		t.Errorf("Expected pixel (0, 0) to be set")
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("Expected VF to be 0, got 0x%X", cpu.V[0xF])
	}
}

func TestOpcodeFX1E(t *testing.T) {
	cpu := setup()
	cpu.I = 0x100
	cpu.V[0] = 0x10
	cpu.decodeAndExecute(0xF01E)
	if cpu.I != 0x110 {
		t.Errorf("Expected I to be 0x110, got 0x%X", cpu.I)
	}
}

func TestOpcodeFX65(t *testing.T) {
	cpu := setup()
	cpu.I = 0x300
	cpu.RAM.WriteByte(0x300, 0x01)
	cpu.RAM.WriteByte(0x301, 0x02)
	cpu.RAM.WriteByte(0x302, 0x03)
	cpu.decodeAndExecute(0xF265)
	if cpu.V[0] != 0x01 || cpu.V[1] != 0x02 || cpu.V[2] != 0x03 {
		t.Errorf("Expected V0, V1, V2 to be 0x01, 0x02, 0x03, got 0x%X, 0x%X, 0x%X", cpu.V[0], cpu.V[1], cpu.V[2])
	}
}
