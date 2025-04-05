// Package emulator provides core interfaces and types for emulator implementations
package emulator

import "errors"

var ErrQuit = errors.New("quit")

// Emulator defines the common interface that all emulator implementations must satisfy
type Emulator interface {
	// Initialize sets up the emulator's initial state
	Initialize() error

	// LoadROM loads a ROM file into the emulator's memory
	LoadROM(path string) error

	// Cycle executes one CPU cycle
	Cycle() error

	// UpdateTimers updates any system timers
	UpdateTimers()

	// HandleInput processes any pending input
	HandleInput() error

	// Render updates the display
	Render() error

	// Cleanup performs any necessary cleanup before shutdown
	Cleanup()
}

// Memory defines the interface for emulator memory operations
type Memory interface {
	// Read reads a byte from the specified address
	Read(addr uint16) byte

	// Write writes a byte to the specified address
	Write(addr uint16, value byte)

	// Load loads data into memory at the specified address
	Load(addr uint16, data []byte) error
}

// CPU defines the interface for emulator CPU operations
type CPU interface {
	// Step executes one instruction
	Step() error

	// Reset resets the CPU to its initial state
	Reset()
}

// Display defines the interface for emulator display operations
type Display interface {
	// Clear clears the display
	Clear()

	// Update updates the display buffer
	Update() error

	// SetPixel sets a pixel at the specified coordinates
	SetPixel(x, y uint8, value bool) bool
}
