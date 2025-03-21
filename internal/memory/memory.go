package memory

import "fmt"

// 4kb of RAM
const MemorySize = 4096

type Memory struct {
	size  int
	bytes [MemorySize]byte
}

func (m *Memory) ReadByte(address uint16) (byte, error) {
	if address >= MemorySize {
		panic("address out of bounds")
	}
	return m.bytes[address], nil
}

func (m *Memory) WriteByte(address uint16, value byte) {
	if address >= MemorySize {
		panic("address out of bounds")
	}
	m.bytes[address] = value
}
func (m *Memory) LoadROM(data []byte) {
	fmt.Printf("ROM size: %d\n", len(data))
	if len(data) > MemorySize-0x200 {
		panic("ROM too large")
	}
	for i, b := range data {
		m.bytes[i+0x200] = b
	}
}

func NewMemory() *Memory {
	return &Memory{
		size:  MemorySize,
		bytes: [MemorySize]byte{},
	}
}
