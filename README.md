# eniGOma

[![Version](https://img.shields.io/badge/version-0.3.0-blue.svg)](https://github.com/coredds/eniGOma/releases)
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

### Unicode & Smart Features  
- ✅ **Auto-Alphabet Detection**: Automatically detects and uses the optimal character set from your input text
- ✅ Support for any Unicode character set (Latin, Greek, Cyrillic, Portuguese, Japanese, etc.)
- ✅ Mixed-language text support (e.g., "Hello! Privет! 日本語!")
- ✅ Predefined alphabets for advanced users (Latin, Greek, Cyrillic, Portuguese, ASCII)
- ✅ Custom alphabet support for specialized use cases
- ✅ Adjustable complexity levels (Low, Medium, High, Extreme)

### Developer Experience
- ✅ Functional options pattern for clean configuration
- ✅ Comprehensive error handling
- ✅ Full JSON serialization of machine state
- ✅ Deep cloning support
- ✅ Extensive unit tests (>95% coverage)

## Installation

```bash
go get github.com/coredds/eniGOma@v0.3.0
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

### Command Line Interface

eniGOma includes a powerful CLI for encryption, decryption, and key management:

```bash
# Install the CLI
go install github.com/coredds/eniGOma/cmd/eniGOma@latest

# Basic encryption with preset
eniGOma encrypt --text "HELLOWORLD" --preset classic

# Decrypt the result
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --preset classic

# Generate a random key configuration
eniGOma keygen --security high --output my-key.json

# Encrypt with custom configuration
eniGOma encrypt --text "Secret Message" --config my-key.json

# Use example configurations
eniGOma encrypt --text "Confidential" --config examples/security-levels/high-security.json
eniGOma encrypt --text "Olá mundo!" --config examples/languages/portuguese-basic.json

# List available presets
eniGOma preset --list

# Get detailed preset information
eniGOma preset --describe classic --verbose

# Encrypt with Unicode support
eniGOma encrypt --text "Hello World!" --alphabet portuguese --security medium

# Validate a configuration file
eniGOma config --validate my-key.json --detailed
```

#### CLI Commands

- **`encrypt`** - Encrypt text or files using an Enigma machine
- **`decrypt`** - Decrypt text or files using an Enigma machine  
- **`keygen`** - Generate random Enigma machine configurations
- **`preset`** - List and describe available machine presets
- **`config`** - Manage and validate configuration files

#### Available Presets

| Preset   | Security | Rotors | Plugboard | Use Case |
|----------|----------|---------|-----------|----------|
| `classic` | Low     | 3       | 2         | Historical simulation, learning |
| `simple`  | Medium  | 5       | 8         | General purpose encryption |
| `high`    | High    | 8       | 15        | Strong obfuscation |
| `extreme` | Extreme | 12      | 20        | Maximum complexity |

#### CLI Examples

```bash
# 🔑 Step 1: Generate a configuration file (required)
eniGOma keygen --output my-key.json

# 🔒 Step 2: Encrypt using the configuration
eniGOma encrypt --text "Hello World!" --config my-key.json

# 🔓 Step 3: Decrypt using the same configuration  
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config my-key.json

# ⚡ Quick workflow: Auto-generate config during encryption
eniGOma encrypt --text "Hello World!" --auto-config my-key.json
eniGOma encrypt --text "Olá Mundo! Café é ótimo!" --auto-config portuguese-key.json  
eniGOma encrypt --text "Mixed: English Русский 日本語!" --auto-config unicode-key.json

# 📁 File encryption/decryption workflow
eniGOma encrypt --file document.txt --config my-key.json --output encrypted.txt
eniGOma decrypt --file encrypted.txt --config my-key.json --output decrypted.txt

# 🎯 Using presets with saved configurations
eniGOma encrypt --text "TOP SECRET" --preset high --save-config high-security.json
eniGOma decrypt --text "ENCRYPTED_OUTPUT" --config high-security.json

# 📊 Configuration management
eniGOma config --show my-key.json --detailed
eniGOma config --test my-key.json --text "TEST MESSAGE"
eniGOma keygen --preset extreme --describe --stats --output extreme-key.json
```

## 🔑 Configuration-First Approach

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

### Why Configuration Files?

**Problem with old approach:**
```bash
eniGOma encrypt --text "Hello World!"  # ❌ No way to decrypt later!
```

**Solution with configuration files:**
```bash
eniGOma encrypt --text "Hello World!" --auto-config my-key.json  # ✅ Always decryptable!
eniGOma decrypt --text "ENCRYPTED" --config my-key.json         # ✅ Works perfectly!
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

### 📁 Configuration Examples

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

### 🔧 Code Examples

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

Current version: **0.2.3**

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