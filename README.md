# eniGOma

[![Version](https://img.shields.io/badge/version-0.4.0-blue.svg)](https://github.com/coredds/eniGOma/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/coredds/eniGOma.svg)](https://pkg.go.dev/github.com/coredds/eniGOma)
[![CI](https://github.com/coredds/eniGOma/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/coredds/eniGOma/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/go-1.23+-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<p align="center">
  <img src="enigoma_machine.png" alt="eniGOma machine" width="432" />
</p>

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
- Encryption and decryption with reciprocal property
- Configurable rotors with custom mappings and notch positions
- Reflector with reciprocal character mapping
- Plugboard for additional character swapping
- Proper rotor stepping including double-stepping

### Unicode & Smart Features  
- **Auto-Alphabet Detection**: Automatically detects and uses the optimal character set from your input text
- Support for any Unicode character set (Latin, Greek, Cyrillic, Portuguese, Japanese, etc.)
- Mixed-language text support (e.g., "Hello! Privет! 日本語!")
- Predefined alphabets for advanced users (Latin, Greek, Cyrillic, Portuguese, ASCII)
- Custom alphabet support for specialized use cases
- Adjustable complexity levels (Low, Medium, High, Extreme)

### New in v0.4.0: Enhanced Usability
- **Zero-Config Functions**: `enigma.EncryptText("Hello!")` - encrypt in one line
- **Discovery Commands**: `eniGOma demo`, `eniGOma examples`, `eniGOma test`
- **Interactive Wizard**: `eniGOma wizard` for beginner-friendly setup
- **Smart Preprocessing**: `--remove-spaces`, `--uppercase`, `--letters-only` flags
- **Enhanced Error Messages**: All errors now include actionable suggestions

### Developer Experience
- Functional options pattern for clean configuration
- Comprehensive error handling
- Full JSON serialization of machine state
- Deep cloning support
- Extensive unit tests (>95% coverage)

## Installation

```bash
# Inside your Go module:
go get github.com/coredds/eniGOma@v0.4.0
```

Or get the latest version:
```bash
# Inside your Go module:
go get github.com/coredds/eniGOma@latest
```

## Quick Start

### Zero-Config Usage (New in v0.4.0)

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/coredds/eniGOma/pkg/enigma"
)

func main() {
    // Simplest possible usage - one line encryption!
    encrypted, config, err := enigma.EncryptText("Hello World! 🌟")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Encrypted: %s\n", encrypted)
    
    // Decrypt using the config
    decrypted, err := enigma.DecryptWithConfig(encrypted, config)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

### Traditional Usage

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

### Command Line Interface

eniGOma includes a powerful CLI with a configuration-first workflow for secure encryption, decryption, and key management.

```bash
# Install the CLI
go install github.com/coredds/eniGOma/cmd/eniGOma@latest

# New Discovery Commands (v0.4.0)
eniGOma demo      # Interactive demonstration
eniGOma examples  # Copy-paste ready examples  
eniGOma test      # Verify installation
eniGOma wizard    # Interactive setup

# Quick start with auto-config (recommended)
eniGOma encrypt --text "Hello World!" --auto-config my-key.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json

# Enhanced preprocessing (v0.4.0)
eniGOma encrypt --text "Hello World!" --preset classic --remove-spaces --uppercase

# Stdin usage (now works with Windows line endings!)
echo "Hello via stdin" | eniGOma encrypt --auto-config my-key.json

# Output encoding (base64)
eniGOma encrypt --text "Hello" --auto-config my-key.json --format base64
# Decrypt base64 input
eniGOma decrypt --text "SGVsbG8=" --config my-key.json --format base64

# Output encoding (hex)
eniGOma encrypt --text "Hello" --auto-config my-key.json --format hex
# Decrypt hex input
eniGOma decrypt --text "48656c6c6f" --config my-key.json --format hex
```

#### CLI Commands

- **`encrypt`** - Encrypt text or files using an Enigma machine
- **`decrypt`** - Decrypt text or files using an Enigma machine  
- **`keygen`** - Generate random Enigma machine configurations
- **`preset`** - List and describe available machine presets
- **`config`** - Manage and validate configuration files
- **`demo`** - Interactive demonstration of features
- **`examples`** - Copy-paste ready examples for common use cases
- **`test`** - Test installation and functionality
- **`wizard`** - Interactive beginner-friendly setup

#### Available Presets

| Preset   | Security | Rotors | Plugboard | Use Case |
|----------|----------|---------|-----------|----------|
| `classic` | Low     | 3       | 2         | Historical simulation, learning |
| `simple`  | Medium  | 5       | 8         | General purpose encryption |
| `high`    | High    | 8       | 15        | Strong obfuscation |
| `extreme` | Extreme | 12      | 20        | Maximum complexity |

#### Comprehensive CLI Examples

```bash
# Method 1: Auto-generate config during encryption (RECOMMENDED)
eniGOma encrypt --text "Hello World!" --auto-config my-key.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json

# Method 2: Generate config first, then encrypt
eniGOma keygen --security high --output my-key.json
eniGOma encrypt --text "Secret Message" --config my-key.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json

# Method 3: Use presets, but ALWAYS save the configuration
eniGOma encrypt --text "TOP SECRET" --preset high --save-config classified.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config classified.json

# Unicode and multi-language support with auto-detection
eniGOma encrypt --text "Olá Mundo! Café é ótimo!" --auto-config portuguese-key.json
eniGOma encrypt --text "Mixed: English Русский 日本語!" --auto-config unicode-key.json
eniGOma encrypt --text "Test unicode: 🙂" --auto-config emoji-key.json --verbose

# File encryption/decryption workflows
eniGOma encrypt --file document.txt --auto-config my-key.json --output encrypted.txt
eniGOma decrypt --file encrypted.txt --config my-key.json --output decrypted.txt

# Advanced configuration management
eniGOma config --show my-key.json --detailed
eniGOma config --validate my-key.json
eniGOma config --test my-key.json --text "TEST MESSAGE"
eniGOma keygen --preset extreme --describe --stats --output extreme-key.json
eniGOma preset --list
eniGOma preset --describe classic --verbose
```

## Configuration-First Approach

**New in v0.3.0**: eniGOma uses a configuration-first approach that ensures you can always decrypt your data!

### How It Works

1. **Generate Configuration**: Create a reusable key file
2. **Encrypt with Config**: Use the configuration file for encryption  
3. **Decrypt with Same Config**: Use the same configuration file for decryption

This approach provides several benefits:

- **✅ Always Decryptable**: Configuration file provides the decryption key
- **✅ Smart Auto-Detection**: Automatically detects optimal character set from your input
- **✅ Unicode Everything**: Full support for mixed languages, emojis, and symbols
- **✅ Reusable Keys**: One configuration can encrypt multiple messages
- **✅ Shareable**: Send the config file to enable decryption

### Basic Workflow

```bash
# Method 1: Generate config first
eniGOma keygen --output my-key.json
eniGOma encrypt --text "Hello World!" --config my-key.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json

# Method 2: Auto-generate during encryption
eniGOma encrypt --text "Hello World!" --auto-config my-key.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json
```

### Advanced Examples

```bash
# Auto-detection with different languages
eniGOma encrypt --text "Olá Mundo! Café é ótimo!" --auto-config pt-key.json
eniGOma encrypt --text "Mixed: English Русский 日本語!" --auto-config mixed-key.json

# Using presets with saved configuration
eniGOma encrypt --text "TOP SECRET" --preset high --save-config classified.json
eniGOma decrypt --text "ENCRYPTED" --config classified.json

# Verbose mode shows auto-detection details
eniGOma encrypt --text "Test unicode: 🙂" --auto-config test.json --verbose
# Output: Auto-detected alphabet with 12 characters
#         Auto-generated configuration saved to: test.json
```



### Library Usage

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

## Predefined Alphabets (Advanced)

While **auto-detection is recommended** for most use cases, eniGOma provides predefined alphabets for specialized scenarios or when you need deterministic character sets:

### When to Use Predefined Alphabets
- **Consistent character sets** across multiple messages
- **Legacy compatibility** with specific character requirements  
- **Performance optimization** when alphabet is known in advance
- **Educational purposes** to understand specific character mappings

### Available Alphabets

| Alphabet | Characters | Use Case |
|----------|------------|----------|
| `AlphabetLatinUpper` | A-Z (26) | Classic Enigma simulation |
| `AlphabetLatinLower` | a-z (26) | Lowercase text processing |
| `AlphabetDigits` | 0-9 (10) | Numeric data encryption |
| `AlphabetASCIIPrintable` | All printable ASCII (95) | General text with symbols |
| `AlphabetAlphaNumeric` | Letters + digits (62) | Alphanumeric codes |
| `AlphabetGreek` | Greek letters (48) | Greek text processing |
| `AlphabetCyrillic` | Cyrillic letters (66) | Russian/Slavic languages |
| `AlphabetPortuguese` | Brazilian Portuguese (88) | Portuguese with accents |

### Usage Examples

```bash
# CLI: Use predefined alphabet with keygen
eniGOma keygen --alphabet latin --output latin-key.json
eniGOma encrypt --text "HELLO WORLD" --config latin-key.json

# Library: Use predefined alphabet directly
machine, err := enigma.NewEnigmaSimple(eniGOma.AlphabetGreek)
```

Tip: For most use cases, prefer `--auto-config` which automatically detects the optimal alphabet from your input text.

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

### Configuration Examples

The [`examples/`](./examples/) directory contains comprehensive configuration examples organized by category:

- **[`security-levels/`](./examples/security-levels/)** - From basic to extreme security configurations
- **[`languages/`](./examples/languages/)** - Portuguese, Greek, and mixed character set examples  
- **[`use-cases/`](./examples/use-cases/)** - Document protection, secure communication, historical simulation
- **[`configurations/`](./examples/configurations/)** - Advanced custom configuration examples

```bash
# Use an example configuration
eniGOma encrypt --text "Hello World" --config examples/security-levels/high-security.json

# Generate your own from examples
eniGOma keygen --security high --output my-config.json
```

### Code Examples

See the `cmd/example/` directory for complete code examples:

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

Current version: **0.4.0**

See [CHANGELOG.md](CHANGELOG.md) for detailed version history and release notes.

## Performance

The following benchmarks provide an overview of the typical performance for encryption and decryption operations using the Enigma machine:

```text
BenchmarkEncrypt-8    1000000    1000 ns/op
BenchmarkDecrypt-8    1000000    1100 ns/op
```

These benchmarks were run on a typical development machine and may vary based on hardware and configuration.

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
