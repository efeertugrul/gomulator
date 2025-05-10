package main

import (
	"log"
	"os"
)

func main() {
	var program = []byte{
		0x00, 0xe0,
		0xa0, 0x96,
		0x60, 0x00,
		0x61, 0x00,
		0xD0, 0x15,

		0xA0, 0x9B,
		0x62, 0x06,
		0x63, 0x00,
		0xD2, 0x35,

		0xA0, 0x96,
		0x64, 0x0C,
		0x65, 0x00,
		0xD4, 0x55,
	}

	// JP 0x200 (loop forever)
	program = append(program, 0x12, 0x00)
	// check if the file exists
	if _, err := os.Stat("testdata/gomulator.ch8"); !os.IsNotExist(err) {
		os.Remove("testdata/gomulator.ch8")
	}

	if err := os.WriteFile("testdata/gomulator.ch8", program, 0644); err != nil {
		log.Fatalf("failed to write ROM: %v", err)
	}
	log.Println("ROM written to gomulator.ch8 (EFE)")
}
