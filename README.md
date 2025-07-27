# eniGOma

[![Version](https://img.shields.io/badge/version-0.2.0-blue.svg)](https://github.com/coredds/eniGOma/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/coredds/eniGOma.svg)](https://pkg.go.dev/github.com/coredds/eniGOma)

A highly customizable, Unicode-capable Enigma machine implementation in Go.

## Overview

eniGOma is a Go library that simulates the famous Enigma machine used during World War II, with modern enhancements including:

- **Unicode Support**: Unlike traditional simulators limited to the Latin alphabet, eniGOma supports any Unicode character set
- **Configurable Complexity**: Adjust encryption complexity with different security levels
- **Modular Design**: Clean, extensible architecture with well-defined interfaces
- **State Management**: Save and load complete machine configurations
- **Historical Accuracy**: Maintains core Enigma behaviors including reciprocal encryption and rotor stepping

## Features

### Core Functionality
- ✅ Encryption and decryption with reciprocal property
- ✅ Configurable rotors with custom mappings and notch positions
- ✅ Reflector with reciprocal character mapping
- ✅ Plugboard for additional character swapping
- ✅ Proper rotor stepping including double-stepping

### Unicode & Customization
- ✅ Support for any Unicode character set (Latin, Greek, Cyrillic, Portuguese, etc.)
- ✅ Predefined alphabets for common use cases
- ✅ Custom alphabet support
- ✅ Built-in Brazilian Portuguese support with full accent support
- ✅ Adjustable complexity levels (Low, Medium, High, Extreme)

### Developer Experience
- ✅ Functional options pattern for clean configuration
- ✅ Comprehensive error handling
- ✅ Full JSON serialization of machine state
- ✅ Deep cloning support
- ✅ Extensive unit tests (>95% coverage)

## Installation

```bash
go get github.com/coredds/eniGOma@v0.2.0
```

Or get the latest version:
```bash
go get github.com/coredds/eniGOma@latest
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/coredds/eniGOma"
    "github.com/coredds/eniGOma/pkg/enigma"
)

func main() {
    // Create a classic Enigma machine
    machine, err := enigma.NewEnigmaClassic()
    if err != nil {
        log.Fatal(err)
    }

    message := "HELLO WORLD"
    
    // Encrypt
    encrypted, err := machine.Encrypt(message)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Encrypted: %s\n", encrypted)

    // Reset to initial state and decrypt
    machine.Reset()
    decrypted, err := machine.Decrypt(encrypted)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### Unicode Support

```go
// Create an Enigma with Greek alphabet
machine, err := enigma.NewEnigmaSimple(eniGOma.AlphabetGreek)
if err != nil {
    log.Fatal(err)
}

message := "Αβγδε ζητα"
encrypted, _ := machine.Encrypt(message)
fmt.Printf("Encrypted Greek: %s\n", encrypted)
```

### Custom Configuration

```go
// Create a high-security Enigma
machine, err := enigma.New(
    enigma.WithAlphabet(eniGOma.AlphabetASCIIPrintable),
    enigma.WithRandomSettings(enigma.High),
    enigma.WithPlugboardConfiguration(map[rune]rune{
        'A': 'Z', 'Z': 'A',
        '1': '9', '9': '1',
    }),
)
```

## Security Levels

eniGOma provides predefined security levels for quick setup:

| Level    | Rotors | Plugboard Pairs | Use Case |
|----------|--------|-----------------|----------|
| Low      | 3      | 2               | Historical simulation, learning |
| Medium   | 5      | 8               | Moderate complexity |
| High     | 8      | 15              | Strong obfuscation |
| Extreme  | 12     | 20              | Maximum complexity |

## Predefined Alphabets

- `AlphabetLatinUpper` - A-Z (26 characters)
- `AlphabetLatinLower` - a-z (26 characters)  
- `AlphabetDigits` - 0-9 (10 characters)
- `AlphabetASCIIPrintable` - All printable ASCII (95 characters)
- `AlphabetAlphaNumeric` - Letters and digits (62 characters)
- `AlphabetGreek` - Greek letters (48 characters)
- `AlphabetCyrillic` - Cyrillic letters (66 characters)
- `AlphabetPortuguese` - **Brazilian Portuguese with accents (88 characters)**

## Advanced Features

### State Serialization

```go
// Save machine state
settings, err := machine.GetSettings()
jsonData, err := machine.SaveSettingsToJSON()

// Restore machine state
newMachine, err := enigma.NewFromJSON(jsonData)
```

### Machine Cloning

```go
// Create independent copy
clone, err := machine.Clone()

// Clones maintain same initial behavior but operate independently
```

### Custom Components

```go
// Create custom rotors, reflectors, and plugboards
rotorSpecs := []rotor.RotorSpec{
    {
        ID: "CustomI",
        ForwardMapping: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
        Notches: []rune{'Q'},
        Position: 0,
        RingSetting: 0,
    },
}

machine, err := enigma.New(
    enigma.WithAlphabet(eniGOma.AlphabetLatinUpper),
    enigma.WithRotorConfiguration(rotorSpecs),
    // ... other options
)
```

## Architecture

eniGOma follows a modular architecture:

```
eniGOma/
├── pkg/enigma/          # Main Enigma machine implementation
├── internal/
│   ├── alphabet/        # Character set management
│   ├── rotor/          # Rotor component
│   ├── reflector/      # Reflector component
│   └── plugboard/      # Plugboard component
├── cmd/example/        # Example applications
└── alphabets.go        # Predefined alphabets
```

### Key Interfaces

- `Rotor` - Defines rotor behavior (forward/backward mapping, stepping)
- `Reflector` - Defines reflector behavior (reciprocal mapping)
- `Plugboard` - Manages character pair swapping

## Testing

Run the comprehensive test suite:

```bash
go test ./...
```

The library includes:
- Unit tests for all components (>95% coverage)
- Integration tests for complete workflows
- Property-based testing for Enigma invariants
- Benchmarks for performance validation

## Examples

See the `cmd/example/` directory for complete examples:

- Basic encryption/decryption
- Unicode alphabet usage
- Security level comparisons
- Settings serialization
- Custom component configuration

## Historical Accuracy

eniGOma maintains historical Enigma machine behaviors:

- **Reciprocal encryption**: If A encrypts to B, then B encrypts to A
- **Rotor stepping**: Proper single and double-stepping mechanics
- **No self-encryption**: No character encrypts to itself (with plugboard and reflector)
- **Deterministic behavior**: Same settings always produce same results

## Version History

Current version: **0.2.0**

See [CHANGELOG.md](CHANGELOG.md) for detailed version history and release notes.

## Performance

Typical performance on modern hardware:
- Single character: ~1μs
- 1KB message: ~1ms
- Setup/configuration: ~100μs

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## Security Notice

⚠️ **Important**: This library is for educational and simulation purposes only. 

Do not use eniGOma for securing sensitive data in production systems. Modern cryptographic algorithms (AES-GCM, ChaCha20-Poly1305) should be used for real-world security applications.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Author

**David Duarte** - *Creator and Lead Developer*

- GitHub: [@coredds](https://github.com/coredds)
- Project: [eniGOma](https://github.com/coredds/eniGOma)

## Acknowledgments

- Historical Enigma machine designers and operators
- The cryptographic community for preserving this important history
- Go community for excellent tooling and libraries

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

*eniGOma - Experience the legendary Enigma machine with modern Go power!* 