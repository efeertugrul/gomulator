package chip8

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := NewCPU()
	if cpu.PC != 0x200 {
		t.Errorf("NewCPU PC = %#x; want 0x200", cpu.PC)
	}
}

func TestCPUReset(t *testing.T) {
	cpu := NewCPU()
	cpu.PC = 0x300
	cpu.SP = 5
	cpu.I = 0x1234
	for i := range cpu.V {
		cpu.V[i] = 0xFF
	}
	for i := range cpu.Stack {
		cpu.Stack[i] = 0xFFFF
	}
	cpu.DelayTimer = 10
	cpu.SoundTimer = 10

	cpu.Reset()

	if cpu.PC != 0x200 {
		t.Errorf("Reset PC = %#x; want 0x200", cpu.PC)
	}
	if cpu.SP != 0 {
		t.Errorf("Reset SP = %d; want 0", cpu.SP)
	}
	if cpu.I != 0 {
		t.Errorf("Reset I = %#x; want 0", cpu.I)
	}
	for i, v := range cpu.V {
		if v != 0 {
			t.Errorf("Reset V[%d] = %#x; want 0", i, v)
		}
	}
	for i, v := range cpu.Stack {
		if v != 0 {
			t.Errorf("Reset Stack[%d] = %#x; want 0", i, v)
		}
	}
	if cpu.DelayTimer != 0 {
		t.Errorf("Reset DelayTimer = %d; want 0", cpu.DelayTimer)
	}
	if cpu.SoundTimer != 0 {
		t.Errorf("Reset SoundTimer = %d; want 0", cpu.SoundTimer)
	}
}
