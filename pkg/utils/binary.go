// Package common provides shared utilities for emulator implementations
package utils

// BoolArrayToByte converts a boolean array to a byte, with configurable offset and wrapping
func BoolArrayToByte(arr []bool, offset int, wrap bool, limit int) byte {
	var result byte
	for i := range 8 {
		location := offset + i
		if location >= limit {
			if !wrap {
				return result
			} else {
				offset = 0
			}
		}

		if arr[offset+i] {
			result |= 1 << i
		}
	}

	return result
}

// ChangeBits modifies a boolean array based on bits in a byte, with configurable offset and wrapping
func ChangeBits(arr []bool, offset int, b byte, wrap bool, limit int) {
	for i := range 8 {
		location := offset + i
		if location >= limit {
			if !wrap {
				return
			} else {
				offset = 0
			}
		}

		if b&1 == 1 {
			arr[offset+i] = true
		} else {
			arr[offset+i] = false
		}

		b >>= 1
	}
}
