package chip8

import (
	"context"
	"testing"
)

func TestIntegration_MinimalROM(t *testing.T) {
	c := New()
	ctx := context.Background()

	// Minimal ROM: 0x6001 (LD V0, 1), 0x6102 (LD V1, 2), 0x8014 (ADD V0, V1)
	rom := []byte{
		0x60, 0x01, // LD V0, 1
		0x61, 0x02, // LD V1, 2
		0x80, 0x14, // ADD V0, V1
	}
	if err := c.Memory.Load(0x200, rom); err != nil {
		t.Fatalf("Failed to load ROM: %v", err)
	}
	c.CPU.PC = 0x200

	// Run three cycles
	for i := 0; i < 3; i++ {
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle %d failed: %v", i, err)
		}
	}

	if c.CPU.V[0] != 3 {
		t.Errorf("V0 = %d; want 3 (1+2)", c.CPU.V[0])
	}
	if c.CPU.V[1] != 2 {
		t.Errorf("V1 = %d; want 2", c.CPU.V[1])
	}
	if c.CPU.PC != 0x206 {
		t.Errorf("PC = %#x; want 0x206 (after 3 instructions)", c.CPU.PC)
	}
}

func TestIntegration_GomulatorROM(t *testing.T) {
	c := New()
	ctx := context.Background()

	if err := c.LoadROM(ctx, "../../testdata/gomulator.ch8"); err != nil {
		t.Fatalf("Failed to load gomulator.ch8: %v", err)
	}
	c.CPU.PC = 0x200

	// Run enough cycles to execute the ROM (44 instructions max, but loop forever)
	for i := 0; i < 40; i++ {
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle %d failed: %v", i, err)
		}
	}

	// Check that at least some pixels are set (text drawn)
	pixelsOn := 0
	for y := 0; y < DisplayHeight; y++ {
		for x := 0; x < DisplayWidth; x++ {
			if c.Display.Pixels[y][x] {
				pixelsOn++
			}
		}
	}
	if pixelsOn == 0 {
		t.Errorf("No pixels were set after running gomulator.ch8; expected text to be drawn")
	}
}
