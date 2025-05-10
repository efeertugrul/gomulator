package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/efeertugrul/gomulator/internal/chip8"
	"github.com/efeertugrul/gomulator/pkg/emulator"
)

const (
	clockSpeedHz = 500 // CHIP-8 typically runs at 500Hz
	frameRate    = 60  // 60 updates per second for delay timers
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic occurred", "panic", r)
			os.Exit(2)
		}
	}()
	if len(os.Args) < 2 {
		logger.Error("Usage: chip8-emulator <ROM file>")
		os.Exit(1)
	}
	romPath := os.Args[1]

	// Set up context and signal handling for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	// Create a new CHIP-8 emulator instance
	var emu emulator.Emulator = chip8.New()
	emu.Initialize(ctx)
	defer emu.Cleanup()

	// Load the ROM file
	if err := emu.LoadROM(ctx, romPath); err != nil {
		logger.Error("Failed to load ROM", "error", err)
		os.Exit(1)
	}

	// Set up timing
	clock := time.NewTicker(time.Second / clockSpeedHz)
	defer clock.Stop()

	timerTicker := time.NewTicker(time.Second / frameRate)
	defer timerTicker.Stop()

	// Main emulation loop
	for {
		select {
		case <-ctx.Done():
			logger.Info("Received shutdown signal, exiting...")
			return
		case <-clock.C:
			if err := emu.Cycle(ctx); err != nil {
				if err == emulator.ErrQuit {
					logger.Info("Quit requested, exiting...")
					return
				}
				logger.Error("Emulation cycle error", "error", err)
			}
		case <-timerTicker.C:
			emu.UpdateTimers()
			if err := emu.HandleInput(ctx); err != nil {
				if err == emulator.ErrQuit {
					logger.Info("Quit requested, exiting...")
					return
				}
				logger.Error("Input handling error", "error", err)
			}
			if err := emu.Render(ctx); err != nil {
				if err == emulator.ErrQuit {
					logger.Info("Quit requested, exiting...")
					return
				}
				logger.Error("Render error", "error", err)
			}
		}
	}
}
