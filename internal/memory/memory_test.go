package memory

import (
	"testing"
)

func TestReadByte(t *testing.T) {
	mem := NewMemory()
	address := uint16(0x200)
	expectedValue := byte(0xAB)
	mem.WriteByte(address, expectedValue)

	value, err := mem.ReadByte(address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != expectedValue {
		t.Errorf("expected %v, got %v", expectedValue, value)
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for out of bounds address")
		}
	}()
	mem.ReadByte(MemorySize)
}

func TestWriteByte(t *testing.T) {
	mem := NewMemory()
	address := uint16(0x200)
	value := byte(0xAB)

	mem.WriteByte(address, value)
	readValue, err := mem.ReadByte(address)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if readValue != value {
		t.Errorf("expected %v, got %v", value, readValue)
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for out of bounds address")
		}
	}()
	mem.WriteByte(MemorySize, value)
}

func TestLoadROM(t *testing.T) {
	mem := NewMemory()
	rom := []byte{0x01, 0x02, 0x03, 0x04}

	mem.LoadROM(rom)

	for i, b := range rom {
		address := uint16(i + 0x200)
		value, err := mem.ReadByte(address)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if value != b {
			t.Errorf("at address %v, expected %v, got %v", address, b, value)
		}
	}

	// Test ROM too large
	largeROM := make([]byte, MemorySize-0x200+1)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for ROM too large")
		}
	}()
	mem.LoadROM(largeROM)
}
