package chip8

import (
	"context"
	"testing"
)

func TestNewChip8(t *testing.T) {
	c := New()
	if c.CPU == nil || c.Memory == nil || c.Display == nil || c.Input == nil {
		t.Error("New() did not initialize all components")
	}
}

func TestUpdateTimers(t *testing.T) {
	c := New()
	c.CPU.DelayTimer = 2
	c.CPU.SoundTimer = 2
	c.UpdateTimers()
	if c.CPU.DelayTimer != 1 {
		t.Errorf("DelayTimer = %d; want 1", c.CPU.DelayTimer)
	}
	if c.CPU.SoundTimer != 1 {
		t.Errorf("SoundTimer = %d; want 1", c.CPU.SoundTimer)
	}
}

func TestCleanupNoPanic(t *testing.T) {
	c := New()
	// Should not panic
	c.Cleanup()
}

func TestRenderNoPanic(t *testing.T) {
	t.Skip("SDL/GUI code cannot be tested in Go unit tests due to main thread requirements.")
}

func Example() {
	c8 := New()
	ctx := context.Background()
	_ = c8.Initialize(ctx)
	_ = c8.LoadROM(ctx, "../../testdata/gomulator.ch8")
	for i := 0; i < 5; i++ {
		_ = c8.Cycle(ctx)
		c8.UpdateTimers()
	}
	c8.Cleanup()
	// Output:
}
