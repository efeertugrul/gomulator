package chip8

import "fmt"

// Memory implements the emulator.Memory interface for CHIP-8
type Memory struct {
	RAM [4096]byte
}

// NewMemory creates a new instance of CHIP-8 memory
func NewMemory() *Memory {
	m := &Memory{}
	// Load fontset at 0x50
	copy(m.RAM[0x50:], defaultFontset[:])
	return m
}

// Read reads a byte from the specified address
func (m *Memory) Read(addr uint16) byte {
	return m.RAM[addr]
}

// Write writes a byte to the specified address
func (m *Memory) Write(addr uint16, value byte) {
	m.RAM[addr] = value
}

// Load loads data into memory at the specified address
func (m *Memory) Load(addr uint16, data []byte) error {
	if int(addr)+len(data) > len(m.RAM) {
		return fmt.Errorf("data exceeds memory bounds")
	}
	copy(m.RAM[addr:], data)
	return nil
}
