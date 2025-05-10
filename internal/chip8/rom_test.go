package chip8

import (
	"context"
	"os"
	"testing"
)

func TestLoadROM_Success(t *testing.T) {
	c := New()
	f, err := os.CreateTemp("", "chip8rom*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(f.Name())
	rom := []byte{0x01, 0x02, 0x03}
	if _, err := f.Write(rom); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	if err := c.LoadROM(context.Background(), f.Name()); err != nil {
		t.Errorf("LoadROM failed: %v", err)
	}
	for i, v := range rom {
		if c.Memory.RAM[0x200+uint16(i)] != v {
			t.Errorf("RAM[0x%X] = %#x; want %#x", 0x200+uint16(i), c.Memory.RAM[0x200+uint16(i)], v)
		}
	}
}

func TestLoadROM_FileNotFound(t *testing.T) {
	c := New()
	if err := c.LoadROM(context.Background(), "nonexistent.rom"); err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestLoadROM_TooLarge(t *testing.T) {
	c := New()
	f, err := os.CreateTemp("", "chip8rom*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(f.Name())
	rom := make([]byte, len(c.Memory.RAM))
	if _, err := f.Write(rom); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	if err := c.LoadROM(context.Background(), f.Name()); err == nil {
		t.Error("expected error for oversized ROM, got nil")
	}
}
