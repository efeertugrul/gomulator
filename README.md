# Gomulator - Multi-System Emulator in Go

Gomulator is a multi-system emulator project written in Go, designed with extensibility in mind to support various classic gaming systems. Currently, it implements a CHIP-8 emulator with plans to support additional systems in the future.

## Features

### CHIP-8 Emulator
- Complete CHIP-8 instruction set implementation
- SDL-based display rendering
- Configurable clock speed (default: 500Hz)


## Requirements
- SDL2 library

## Usage

### CHIP-8 Emulator

To run a CHIP-8 ROM:
```bash
go run cmd/chip8-emulator/main.go <path-to-rom>
```

Or build and run the binary:
```bash
go build -o chip8 cmd/chip8-emulator/main.go
./chip8 <path-to-rom>
```


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## Acknowledgments

- Thanks to the Go SDL2 binding developers
- CHIP-8 Technical Reference used in implementation