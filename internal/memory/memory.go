package memory

// 4kb of RAM
const MemorySize = 4096

type Memory struct {
	size  int
	bytes [MemorySize]byte
}

func (Memory) InitMemory() Memory {
	return Memory{size: MemorySize, bytes: [MemorySize]byte{}}
}

func (m *Memory) ReadByte(address int) byte {
	if address < 0 || address >= MemorySize {
		panic("address out of bounds")
	}
	return m.bytes[address]
}

func (m *Memory) WriteByte(address int, value byte) {
	if address < 0 || address >= MemorySize {
		panic("address out of bounds")
	}
	m.bytes[address] = value
}
