// Package emulator provides core interfaces and types for emulator implementations.
//
// These interfaces define the contract for all emulator cores, allowing for modularity and extensibility.
// Implementations for specific systems (e.g., CHIP-8) should satisfy these interfaces.
package emulator

import (
	"context"
	"errors"
)

// ErrQuit is returned to signal a request to quit the emulator.
var ErrQuit = errors.New("quit")

// Emulator defines the common interface that all emulator implementations must satisfy.
// It provides lifecycle management, ROM loading, emulation cycles, input, rendering, and cleanup.
type Emulator interface {
	// Initialize sets up the emulator's initial state.
	Initialize(ctx context.Context) error

	// LoadROM loads a ROM file into the emulator's memory.
	LoadROM(ctx context.Context, path string) error

	// Cycle executes one CPU cycle.
	Cycle(ctx context.Context) error

	// UpdateTimers updates any system timers (e.g., delay/sound timers).
	UpdateTimers()

	// HandleInput processes any pending input.
	HandleInput(ctx context.Context) error

	// Render updates the display.
	Render(ctx context.Context) error

	// Cleanup performs any necessary cleanup before shutdown.
	Cleanup()
}

// Memory defines the interface for emulator memory operations.
// Implementations should provide safe, bounds-checked access to emulated memory.
type Memory interface {
	// Read reads a byte from the specified address.
	Read(addr uint16) byte

	// Write writes a byte to the specified address.
	Write(addr uint16, value byte)

	// Load loads data into memory at the specified address.
	Load(addr uint16, data []byte) error
}

// CPU defines the interface for emulator CPU operations.
// Implementations should provide instruction stepping and reset logic.
type CPU interface {
	// Step executes one instruction.
	Step() error

	// Reset resets the CPU to its initial state.
	Reset()
}

// Display defines the interface for emulator display operations.
// Implementations should provide pixel-level access and display buffer management.
type Display interface {
	// Clear clears the display.
	Clear()

	// Update updates the display buffer (e.g., present to screen).
	Update() error

	// SetPixel sets a pixel at the specified coordinates. Returns true if the pixel was previously set.
	SetPixel(x, y uint8, value bool) bool
}
