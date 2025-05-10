// Package chip8 provides a CHIP-8 emulator implementation for the Gomulator project.
//
// This package contains the CHIP-8 CPU, memory, display, input, and main emulator logic.
package chip8

import (
	"context"
	"log/slog"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Chip8 implements the CHIP-8 emulator core, including CPU, memory, display, and input.
type Chip8 struct {
	CPU     *CPU
	Memory  *Memory
	Display *Display
	Input   *Input
}

// New returns a new, initialized CHIP-8 emulator instance.
func New() *Chip8 {
	chip8 := &Chip8{
		CPU:     NewCPU(),
		Memory:  NewMemory(),
		Display: NewDisplay(),
		Input:   &Input{},
	}

	return chip8
}

// InitDisplay initializes the SDL display for CHIP-8 output.
func (c *Chip8) InitDisplay() {
	var err error
	sdl.Init(sdl.INIT_VIDEO)

	c.Display.sdl, err = newSDLRenderer()
	if err != nil {
		slog.Default().Error("SDL Init failed", "error", err)
		os.Exit(1)
	}
}

// Initialize sets up the CHIP-8 emulator state and display.
func (c *Chip8) Initialize(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	c.InitDisplay()
	return nil
}

// UpdateTimers decrements the delay and sound timers, and plays sound if needed.
func (c *Chip8) UpdateTimers() {
	if c.CPU.DelayTimer > 0 {
		c.CPU.DelayTimer--
	}
	if c.CPU.SoundTimer > 0 {
		c.CPU.SoundTimer--
		c.PlaySound()
	}
}

// Cleanup releases all emulator resources and quits SDL.
func (c *Chip8) Cleanup() {
	if c.Display != nil {
		c.Display.Cleanup()
	}
	sdl.Quit()
}

// Render updates the CHIP-8 display.
func (c *Chip8) Render(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	c.Display.Render()
	return nil
}
