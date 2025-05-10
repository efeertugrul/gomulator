package chip8

import (
	"context"
	"testing"
)

func TestCycleAndDecodeAndExecute_LogicOpcodes(t *testing.T) {
	c := New()
	ctx := context.Background()

	t.Run("0x6000: LD Vx, byte", func(t *testing.T) {
		c.Memory.RAM[0x200] = 0x60
		c.Memory.RAM[0x201] = 0x0A
		c.CPU.PC = 0x200
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.V[0] != 0x0A {
			t.Errorf("V0 = %#x; want 0x0A", c.CPU.V[0])
		}
	})

	t.Run("0x7000: ADD Vx, byte", func(t *testing.T) {
		c.Memory.RAM[0x202] = 0x70
		c.Memory.RAM[0x203] = 0x01
		c.CPU.PC = 0x202
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.V[0] != 0x0B {
			t.Errorf("V0 = %#x; want 0x0B", c.CPU.V[0])
		}
	})

	t.Run("0xA000: LD I, addr", func(t *testing.T) {
		c.Memory.RAM[0x204] = 0xA2
		c.Memory.RAM[0x205] = 0x34
		c.CPU.PC = 0x204
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.I != 0x234 {
			t.Errorf("I = %#x; want 0x234", c.CPU.I)
		}
	})

	t.Run("0x3000: SE Vx, byte", func(t *testing.T) {
		c.CPU.V[1] = 0x12
		c.Memory.RAM[0x206] = 0x31
		c.Memory.RAM[0x207] = 0x12
		c.CPU.PC = 0x206
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x20A {
			t.Errorf("PC = %#x; want 0x20A (skipped)", c.CPU.PC)
		}
	})

	t.Run("0x4000: SNE Vx, byte", func(t *testing.T) {
		c.CPU.V[2] = 0x34
		c.Memory.RAM[0x20A] = 0x42
		c.Memory.RAM[0x20B] = 0x33
		c.CPU.PC = 0x20A
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x20E {
			t.Errorf("PC = %#x; want 0x20E (skipped)", c.CPU.PC)
		}
	})

	t.Run("0x5000: SE Vx, Vy", func(t *testing.T) {
		c.CPU.V[3] = 0x55
		c.CPU.V[4] = 0x55
		c.Memory.RAM[0x20E] = 0x53
		c.Memory.RAM[0x20F] = 0x40
		c.CPU.PC = 0x20E
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x212 {
			t.Errorf("PC = %#x; want 0x212 (skipped)", c.CPU.PC)
		}
	})

	t.Run("0x9000: SNE Vx, Vy", func(t *testing.T) {
		c.CPU.V[5] = 0x77
		c.CPU.V[6] = 0x88
		c.Memory.RAM[0x212] = 0x95
		c.Memory.RAM[0x213] = 0x60
		c.CPU.PC = 0x212
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x216 {
			t.Errorf("PC = %#x; want 0x216 (skipped)", c.CPU.PC)
		}
	})

	t.Run("0x1000: JP addr", func(t *testing.T) {
		c.Memory.RAM[0x216] = 0x12
		c.Memory.RAM[0x217] = 0x34
		c.CPU.PC = 0x216
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x234 {
			t.Errorf("PC = %#x; want 0x234", c.CPU.PC)
		}
	})

	t.Run("0xB000: JP V0, addr", func(t *testing.T) {
		c.CPU.V[0] = 0x10
		c.Memory.RAM[0x218] = 0xB1
		c.Memory.RAM[0x219] = 0x00
		c.CPU.PC = 0x218
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x110 {
			t.Errorf("PC = %#x; want 0x110", c.CPU.PC)
		}
	})

	t.Run("0x2000: CALL addr and 0x00EE: RET", func(t *testing.T) {
		c.CPU.PC = 0x220
		c.Memory.RAM[0x220] = 0x22
		c.Memory.RAM[0x221] = 0x30 // CALL 0x230
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x230 {
			t.Errorf("PC = %#x; want 0x230", c.CPU.PC)
		}
		if c.CPU.SP != 1 {
			t.Errorf("SP = %d; want 1", c.CPU.SP)
		}
		// Now test RET
		c.Memory.RAM[0x230] = 0x00
		c.Memory.RAM[0x231] = 0xEE
		c.CPU.PC = 0x230
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.SP != 0 {
			t.Errorf("SP = %d; want 0", c.CPU.SP)
		}
	})

	t.Run("0x8000 group: LD, OR, AND, XOR, ADD, SUB, SHR, SUBN, SHL", func(t *testing.T) {
		c.CPU.V[1] = 0xF0
		c.CPU.V[2] = 0x0F
		c.Memory.RAM[0x222] = 0x81
		c.Memory.RAM[0x223] = 0x20 // LD V1, V2
		c.CPU.PC = 0x222
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0x0F {
			t.Errorf("LD V1, V2 failed")
		}
		c.CPU.V[1] = 0xF0
		c.Memory.RAM[0x224] = 0x81
		c.Memory.RAM[0x225] = 0x21 // OR V1, V2
		c.CPU.PC = 0x224
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0xFF {
			t.Errorf("OR V1, V2 failed")
		}
		c.CPU.V[1] = 0xF0
		c.Memory.RAM[0x226] = 0x81
		c.Memory.RAM[0x227] = 0x22 // AND V1, V2
		c.CPU.PC = 0x226
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0x00 {
			t.Errorf("AND V1, V2 failed")
		}
		c.CPU.V[1] = 0xF0
		c.Memory.RAM[0x228] = 0x81
		c.Memory.RAM[0x229] = 0x23 // XOR V1, V2
		c.CPU.PC = 0x228
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0xFF {
			t.Errorf("XOR V1, V2 failed")
		}
		c.CPU.V[1] = 0xF0
		c.Memory.RAM[0x22A] = 0x81
		c.Memory.RAM[0x22B] = 0x24 // ADD V1, V2
		c.CPU.PC = 0x22A
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0xFF || c.CPU.V[15] != 0x0 {
			t.Errorf("ADD V1, V2 failed")
		}
		c.CPU.V[1] = 0x10
		c.CPU.V[2] = 0x20
		c.Memory.RAM[0x22C] = 0x81
		c.Memory.RAM[0x22D] = 0x25 // SUB V1, V2
		c.CPU.PC = 0x22C
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0xF0 || c.CPU.V[15] != 0x0 {
			t.Errorf("SUB V1, V2 failed")
		}
		c.CPU.V[1] = 0x01
		c.Memory.RAM[0x22E] = 0x81
		c.Memory.RAM[0x22F] = 0x26 // SHR V1
		c.CPU.PC = 0x22E
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0x00 || c.CPU.V[15] != 0x1 {
			t.Errorf("SHR V1 failed")
		}
		c.CPU.V[1] = 0x10
		c.CPU.V[2] = 0x20
		c.Memory.RAM[0x230] = 0x81
		c.Memory.RAM[0x231] = 0x27 // SUBN V1, V2
		c.CPU.PC = 0x230
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0x10 || c.CPU.V[15] != 0x1 {
			t.Errorf("SUBN V1, V2 failed")
		}
		c.CPU.V[1] = 0x80
		c.Memory.RAM[0x232] = 0x81
		c.Memory.RAM[0x233] = 0x2E // SHL V1
		c.CPU.PC = 0x232
		_ = c.Cycle(ctx)
		if c.CPU.V[1] != 0x00 || c.CPU.V[15] != 0x1 {
			t.Errorf("SHL V1 failed")
		}
	})

	t.Run("0xF007: LD Vx, DT", func(t *testing.T) {
		c.CPU.DelayTimer = 0x42
		c.Memory.RAM[0x240] = 0xF1
		c.Memory.RAM[0x241] = 0x07 // LD V1, DT
		c.CPU.PC = 0x240
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.V[1] != 0x42 {
			t.Errorf("V1 = %#x; want 0x42 (delay timer)", c.CPU.V[1])
		}
	})

	t.Run("Unknown opcode", func(t *testing.T) {
		c.Memory.RAM[0x234] = 0xFF
		c.Memory.RAM[0x235] = 0xFF
		c.CPU.PC = 0x234
		_ = c.Cycle(ctx)
		// Should not panic or error
	})

	t.Run("0xE09E: SKP Vx (skip if key pressed)", func(t *testing.T) {
		c.CPU.V[2] = 0x1
		c.Input.Keys[0x1] = true
		c.Memory.RAM[0x250] = 0xE2
		c.Memory.RAM[0x251] = 0x9E
		c.CPU.PC = 0x250
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x254 {
			t.Errorf("PC = %#x; want 0x254 (skipped)", c.CPU.PC)
		}
	})

	t.Run("0xE0A1: SKNP Vx (skip if key not pressed)", func(t *testing.T) {
		c.CPU.V[3] = 0x2
		c.Input.Keys[0x2] = false
		c.Memory.RAM[0x260] = 0xE3
		c.Memory.RAM[0x261] = 0xA1
		c.CPU.PC = 0x260
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x264 {
			t.Errorf("PC = %#x; want 0x264 (skipped)", c.CPU.PC)
		}
	})

	t.Run("0xF00A: LD Vx, K (wait for key)", func(t *testing.T) {
		c.CPU.V[4] = 0
		c.Input.Keys[0x3] = true
		c.Memory.RAM[0x270] = 0xF4
		c.Memory.RAM[0x271] = 0x0A
		c.CPU.PC = 0x270
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.V[4] != 0x3 {
			t.Errorf("V4 = %#x; want 0x3 (key index)", c.CPU.V[4])
		}
	})

	t.Run("0xF015: LD DT, Vx", func(t *testing.T) {
		c.CPU.V[5] = 0x77
		c.Memory.RAM[0x280] = 0xF5
		c.Memory.RAM[0x281] = 0x15
		c.CPU.PC = 0x280
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.DelayTimer != 0x77 {
			t.Errorf("DelayTimer = %#x; want 0x77", c.CPU.DelayTimer)
		}
	})

	t.Run("0xF018: LD ST, Vx", func(t *testing.T) {
		c.CPU.V[6] = 0x88
		c.Memory.RAM[0x290] = 0xF6
		c.Memory.RAM[0x291] = 0x18
		c.CPU.PC = 0x290
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.SoundTimer != 0x88 {
			t.Errorf("SoundTimer = %#x; want 0x88", c.CPU.SoundTimer)
		}
	})

	t.Run("0xF01E: ADD I, Vx", func(t *testing.T) {
		c.CPU.V[7] = 0x10
		c.CPU.I = 0x100
		c.Memory.RAM[0x2A0] = 0xF7
		c.Memory.RAM[0x2A1] = 0x1E
		c.CPU.PC = 0x2A0
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.I != 0x110 {
			t.Errorf("I = %#x; want 0x110", c.CPU.I)
		}
	})

	t.Run("0xF029: LD F, Vx", func(t *testing.T) {
		c.CPU.V[8] = 0xA
		c.Memory.RAM[0x2B0] = 0xF8
		c.Memory.RAM[0x2B1] = 0x29
		c.CPU.PC = 0x2B0
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.I != 0x32 {
			t.Errorf("I = %#x; want 0x32 (font address)", c.CPU.I)
		}
	})

	t.Run("0xF033: LD B, Vx", func(t *testing.T) {
		c.CPU.V[9] = 123
		c.CPU.I = 0x300
		c.Memory.RAM[0x2C0] = 0xF9
		c.Memory.RAM[0x2C1] = 0x33
		c.CPU.PC = 0x2C0
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.Memory.RAM[0x300] != 1 || c.Memory.RAM[0x301] != 2 || c.Memory.RAM[0x302] != 3 {
			t.Errorf("BCD conversion failed: got %d %d %d; want 1 2 3", c.Memory.RAM[0x300], c.Memory.RAM[0x301], c.Memory.RAM[0x302])
		}
	})

	t.Run("0xF055: LD [I], Vx", func(t *testing.T) {
		c.CPU.I = 0x350
		c.CPU.V[0] = 0xAA
		c.CPU.V[1] = 0xBB
		c.Memory.RAM[0x2D0] = 0xF1
		c.Memory.RAM[0x2D1] = 0x55
		c.CPU.PC = 0x2D0
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.Memory.RAM[0x350] != 0xAA || c.Memory.RAM[0x351] != 0xBB {
			t.Errorf("LD [I], Vx failed: got %x %x; want AA BB", c.Memory.RAM[0x350], c.Memory.RAM[0x351])
		}
	})

	t.Run("0xF065: LD Vx, [I]", func(t *testing.T) {
		c.CPU.I = 0x360
		c.Memory.RAM[0x360] = 0x11
		c.Memory.RAM[0x361] = 0x22
		c.Memory.RAM[0x362] = 0x33
		c.Memory.RAM[0x2E0] = 0xF2
		c.Memory.RAM[0x2E1] = 0x65
		c.CPU.PC = 0x2E0
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.V[0] != 0x11 || c.CPU.V[1] != 0x22 || c.CPU.V[2] != 0x33 {
			t.Errorf("LD Vx, [I] failed: got %x %x %x; want 11 22 33", c.CPU.V[0], c.CPU.V[1], c.CPU.V[2])
		}
	})

	t.Run("Unknown opcode (explicit)", func(t *testing.T) {
		c.CPU.PC = 0x300
		c.Memory.RAM[0x300] = 0xF0
		c.Memory.RAM[0x301] = 0xFF // 0xF0FF is not a valid CHIP-8 opcode
		if err := c.Cycle(ctx); err != nil {
			t.Fatalf("Cycle failed: %v", err)
		}
		if c.CPU.PC != 0x302 {
			t.Errorf("PC = %#x; want 0x302 (should advance by 2)", c.CPU.PC)
		}
	})
}
