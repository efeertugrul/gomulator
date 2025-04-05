package main

import (
	"log"
	"os"
	"time"

	"github.com/efeertugrul/gomulator/internal/chip8"
	"github.com/efeertugrul/gomulator/pkg/emulator"
)

const (
	clockSpeedHz = 500 // CHIP-8 typically runs at 500Hz
	frameRate    = 60  // 60 updates per second for delay timers
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: chip8-emulator <ROM file>")
	}
	romPath := os.Args[1]

	// Create a new CHIP-8 emulator instance
	var emu emulator.Emulator = chip8.New()
	emu.Initialize()
	defer emu.Cleanup()

	// Load the ROM file
	if err := emu.LoadROM(romPath); err != nil {
		log.Fatalf("Failed to load ROM: %v", err)
	}

	// Set up timing
	clock := time.NewTicker(time.Second / clockSpeedHz)
	defer clock.Stop()

	timerTicker := time.NewTicker(time.Second / frameRate)
	defer timerTicker.Stop()

	// Main emulation loop
	for {
		select {
		case <-clock.C:
			if err := emu.Cycle(); err != nil {
				log.Printf("Emulation cycle error: %v", err)
			}
		case <-timerTicker.C:
			emu.UpdateTimers()
			if err := emu.HandleInput(); err != nil {
				log.Printf("Input handling error: %v", err)
			}
			if err := emu.Render(); err != nil {
				log.Printf("Render error: %v", err)
			}
		}
	}
}
