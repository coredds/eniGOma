# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.3.1] - 2025-01-31

### Added
- CLI flags: `--auto-config` and `--save-config`
- Stdin support for `encrypt` (pipe input directly)
- Proper hex/base64 encoding (encrypt) and decoding (decrypt) via stdlib

### Changed
- Default alphabet documented and enforced as auto-detected for `encrypt` (equivalent to `--alphabet=auto`)
- README and USAGE updated with configuration-first workflow, stdin, and encoding examples
- Documentation consistency improvements (removed decorative emojis; reorganized CLI examples)

### Fixed
- Version tests updated to avoid pinning an exact patch version

## [0.3.0] - 2025-01-30

### Major Features Added
- **Smart Auto-Alphabet Detection** - Automatically detects optimal character set from input text
- **Universal Unicode Support** - Encrypt any text in any language without alphabet selection
- **Mixed-Language Text Support** - Handle text combining multiple languages (e.g., "Hello! Привет! 日本語!")
- **Emoji & Symbol Support** - Full support for Unicode symbols, emojis, and special characters

### Added
- `alphabet.AutoDetectFromText()` function with configurable options
- Auto-detection as default CLI behavior (`--alphabet auto` is now default)
- Automatic even-size padding for reflector compatibility
- Safety limits for auto-detected alphabets (max 1000 characters)
- Verbose mode showing detected alphabet statistics
- Support for any Unicode character set including CJK, symbols, and emojis

### Changed
- **Breaking**: CLI default changed from `--alphabet latin` to `--alphabet auto`
- CLI examples updated to showcase auto-detection workflow
- Help text updated to reflect simplified usage patterns
- Documentation extensively updated with auto-detection examples

### Enhanced
- Backwards compatibility maintained - all existing alphabet flags still work
- Performance optimizations for large Unicode character sets
- Deterministic alphabet ordering for consistent behavior
- Comprehensive test coverage for auto-detection functionality

### Technical
- Added auto-detection unit tests covering various Unicode scenarios
- Updated CLI tests to verify auto-detection behavior
- Enhanced documentation with migration guide and best practices
- Maintained 100% backwards compatibility with existing code

## [0.2.1] - 2025-01-28

### Added
- **Comprehensive CLI Tool** - Complete command-line interface using Cobra framework
- **5 CLI Commands**: encrypt, decrypt, keygen, preset, config
- **4 Security Presets**: classic, simple, high, extreme with detailed descriptions
- **Unicode CLI Support**: All 8 predefined alphabets accessible via CLI
- **Configuration Management**: JSON import/export, validation, and testing
- **File I/O Operations**: Encrypt/decrypt files with multiple output formats
- **Enhanced Documentation**: Updated README with CLI usage examples and installation guide

### Changed
- Updated project documentation to reflect dual-interface nature (library + CLI)
- Enhanced README with comprehensive CLI examples and command reference

### Technical
- Added Cobra and Viper dependencies for robust CLI experience
- Comprehensive unit tests for CLI functionality (>95% coverage)
- Integration tests for complete encrypt/decrypt workflows
- Maintained full backward compatibility with library interface

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