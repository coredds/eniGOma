# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2025-01-27

### Author
- **David Duarte** - Creator and Lead Developer

### Added
- Complete eniGOma library implementation
- **AlphabetPortuguese** - Built-in Brazilian Portuguese alphabet with full accent support (88 characters)
- Unicode support for any character set (Latin, Greek, Cyrillic, etc.)
- Configurable security levels (Low, Medium, High, Extreme)
- Functional options pattern for clean configuration
- State management with JSON serialization/deserialization
- Comprehensive unit tests (>95% coverage)
- Modular architecture with clean interfaces
- Predefined alphabets for common use cases
- Example application demonstrating all features
- Full documentation and README

### Features
- **Core Enigma Simulation**: Encrypt/Decrypt/Reset functionality
- **Customizable Components**: Rotors, reflectors, and plugboards
- **Historical Accuracy**: Proper rotor stepping and double-stepping
- **Deep Cloning**: Independent machine copies
- **Error Handling**: Comprehensive validation and error reporting
- **Performance**: Optimized for typical use cases

### Components
- **Alphabet Management**: Unicode character set handling
- **Rotor Component**: Forward/backward mapping with stepping mechanics
- **Reflector Component**: Reciprocal character mapping
- **Plugboard Component**: Character pair swapping
- **Main Enigma Engine**: Complete machine simulation

### Technical
- Go 1.21+ support
- Modular internal packages
- Extensive test coverage
- Cryptographically secure randomness using crypto/rand
- JSON serialization support
- Clean API following Go best practices

## [Unreleased]

### Planned
- Performance benchmarks
- Additional historical rotor configurations
- Web interface example
- Advanced stepping mechanisms 