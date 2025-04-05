package chip8

import (
	"log"
	"math/rand"
)

func (c *Chip8) Cycle() error {
	opcode := uint16(c.Memory.RAM[c.CPU.PC])<<8 | uint16(c.Memory.RAM[c.CPU.PC+1])
	log.Printf("Current Operation 0x%X", opcode)
	c.decodeAndExecute(opcode)

	return nil
}

func (c *Chip8) decodeAndExecute(opcode uint16) {
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode {
		case 0x00E0: // cls
			c.Display.Clear()
		case 0x00EE: // return from a subroutine
			c.CPU.PC = c.CPU.Stack[c.CPU.SP]
			c.CPU.SP--
		}
	case 0x1000:
		c.CPU.PC = opcode & 0x0FFF
		return
	case 0x2000:
		c.CPU.SP++
		c.CPU.Stack[c.CPU.SP] = (c.CPU.PC)
		c.CPU.PC = opcode & 0x0FFF
		return
	case 0x3000:
		if c.CPU.V[(opcode&0x0F00)>>8] == byte((opcode & 0x00FF)) {
			c.CPU.PC += 2
		}
	case 0x4000:
		if c.CPU.V[(opcode&0x0F00)>>8] != byte((opcode & 0x00FF)) {
			c.CPU.PC += 2
		}
	case 0x5000:
		if c.CPU.V[(opcode&0x0F00)>>8] == c.CPU.V[(opcode&0x00F0)>>4] {
			c.CPU.PC += 2
		}
	case 0x6000:
		c.CPU.V[(opcode&0x0F00)>>8] = byte(opcode & 0x00FF)
	case 0x7000:
		c.CPU.V[(opcode&0x0F00)>>8] += byte(opcode & 0x00FF)
	case 0x8000:
		switch opcode & 0x000F {
		case 0: // LD
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x00F0)>>4]
		case 1: // OR
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] | c.CPU.V[(opcode&0x00F0)>>4]
		case 2: // AND
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] & c.CPU.V[(opcode&0x00F0)>>4]
		case 3: // XOR
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] ^ c.CPU.V[(opcode&0x00F0)>>4]
		case 4: // ADD
			result := uint16(c.CPU.V[(opcode&0x0F00)>>8]) + uint16(c.CPU.V[(opcode&0x00F0)>>4])
			c.CPU.V[(opcode&0x0F00)>>8] = byte(result)
			c.CPU.V[15] = byte(result >> 8)
		case 5: // SUB
			if c.CPU.V[(opcode&0x0F00)>>8] > c.CPU.V[(opcode&0x00F0)>>4] {
				c.CPU.V[15] = 1
			} else {
				c.CPU.V[15] = 0
			}
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] - c.CPU.V[(opcode&0x00F0)>>4]
		case 6: // SHR
			c.CPU.V[15] = c.CPU.V[(opcode&0x0F00)>>8] & 1
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] >> 1
		case 7: // SUBN
			if c.CPU.V[(opcode&0x0F00)>>8] < c.CPU.V[(opcode&0x00F0)>>4] {
				c.CPU.V[15] = 1
			} else {
				c.CPU.V[15] = 0
			}
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x00F0)>>4] - c.CPU.V[(opcode&0x0F00)>>8]
		case 0xE: // SHL
			c.CPU.V[15] = c.CPU.V[(opcode&0x0F00)>>8] >> 7
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.V[(opcode&0x0F00)>>8] << 1
		}
	case 0x9000:
		if c.CPU.V[(opcode&0x0F00)>>8] != c.CPU.V[(opcode&0x00F0)>>4] {
			c.CPU.PC += 2
		}
	case 0xA000:
		c.CPU.I = uint16(opcode & 0x0FFF)
	case 0xB000:
		c.CPU.PC = uint16(opcode&0x0FFF) + uint16(c.CPU.V[0])
		return
	case 0xC000:
		c.CPU.V[(opcode&0x0F00)>>8] = byte(rand.Uint32()) & byte(opcode&0x00FF)
	case 0xD000: // DRW Vx, Vy, nibble

		nibble := opcode & 0x000F
		c.CPU.V[15] = 0
		x := uint16(c.CPU.V[(opcode&0x0F00)>>8])
		y := uint16(c.CPU.V[(opcode&0x00F0)>>4])
		for yLine := uint16(0); yLine < nibble; yLine++ {
			sprite := c.Memory.RAM[c.CPU.I+yLine]
			for xLine := uint16(0); xLine < 8; xLine++ {
				if sprite&(0x80>>xLine) != 0 {
					if c.Display.Pixels[yLine+y][xLine+x] {
						c.CPU.V[15] = 1
					}

					c.Display.Pixels[yLine+y][xLine+x] = !c.Display.Pixels[yLine+y][xLine+x]
				}
			}
		}
	case 0xE000:
		switch opcode & 0x00FF {
		case 0x009E:
			if c.Input.Keys[c.CPU.V[(opcode&0x0F00)>>8]] {
				c.CPU.PC += 2
			}
		case 0x00A1:
			if !c.Input.Keys[c.CPU.V[(opcode&0x0F00)>>8]] {
				c.CPU.PC += 2
			}
		}
	case 0xF000:
		switch opcode & 0x00FF {
		case 0x0007:
			c.CPU.V[(opcode&0x0F00)>>8] = c.CPU.DelayTimer
		case 0x000A:
			keyPressed := false
			for i := range c.Input.Keys {
				if c.Input.Keys[i] {
					keyPressed = true
					c.CPU.V[(opcode&0x0F00)>>8] = byte(i)
					break
				}

				if !keyPressed {
					return
				}
			}
		case 0x0015:
			c.CPU.DelayTimer = c.CPU.V[(opcode&0x0F00)>>8]
		case 0x0018:
			c.CPU.SoundTimer = c.CPU.V[(opcode&0x0F00)>>8]
		case 0x001E:
			c.CPU.I += uint16(c.CPU.V[(opcode&0x0F00)>>8])
		case 0x0029:
			c.CPU.I = uint16(c.CPU.V[(opcode&0x0F00)>>8] * 5)
		case 0x0033:
			c.Memory.RAM[c.CPU.I] = c.CPU.V[(opcode&0x0F00)>>8] / 100
			c.Memory.RAM[c.CPU.I+1] = (c.CPU.V[(opcode&0x0F00)>>8] / 10) % 10
			c.Memory.RAM[c.CPU.I+2] = (c.CPU.V[(opcode&0x0F00)>>8] % 100) % 10
		case 0x0055:
			for i := 0; i <= int((opcode&0x0F00)>>8); i++ {
				c.Memory.RAM[c.CPU.I+uint16(i)] = c.CPU.V[i]
			}
			c.CPU.I += uint16((opcode&0x0F00)>>8) + 1
		case 0x0065:
			for i := 0; i <= int((opcode&0x0F00)>>8); i++ {
				c.CPU.V[i] = c.Memory.RAM[c.CPU.I+uint16(i)]
			}

			c.CPU.I += uint16((opcode&0x0F00)>>8) + 1
		}
	default:
		log.Printf("Unknown opcode: 0x0%X\n", opcode)
	}

	c.CPU.PC += 2
}
