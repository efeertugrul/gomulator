package chip8

import (
	"testing"
)

func TestMemoryReadWrite(t *testing.T) {
	mem := NewMemory()
	addr := uint16(0x200)
	value := byte(0xAB)
	mem.Write(addr, value)
	if got := mem.Read(addr); got != value {
		t.Errorf("Read(%#x) = %#x; want %#x", addr, got, value)
	}
}

func TestMemoryLoad(t *testing.T) {
	mem := NewMemory()
	addr := uint16(0x300)
	data := []byte{0x01, 0x02, 0x03}
	err := mem.Load(addr, data)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	for i, v := range data {
		if got := mem.Read(addr + uint16(i)); got != v {
			t.Errorf("Read(%#x) = %#x; want %#x", addr+uint16(i), got, v)
		}
	}
}

func TestMemoryLoadBounds(t *testing.T) {
	mem := NewMemory()
	addr := uint16(len(mem.RAM) - 2)
	data := []byte{0x01, 0x02, 0x03}
	err := mem.Load(addr, data)
	if err == nil {
		t.Error("expected error for out-of-bounds load, got nil")
	}
}
