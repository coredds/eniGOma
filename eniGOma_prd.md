# Product Requirements Document: eniGOma Go Library

**Document Version:** 0.2  
**Date:** July 27, 2025  
**Author:** David Duarte

## 1. Introduction

### 1.1 Purpose

This Product Requirements Document (PRD) outlines the specifications for eniGOma, a Go library designed to implement a highly customizable virtual Enigma machine. Unlike traditional Enigma simulators limited to the Latin alphabet, eniGOma will support Unicode characters, providing a flexible platform for exploring historical cryptography concepts with modern character sets. A key feature will be the ability for users to adjust the encryption's complexity, allowing for various levels of "protection" or demonstration.

### 1.2 Goals

- Provide a robust, idiomatic Go library for simulating an Enigma machine.
- Support encryption and decryption of messages using Unicode characters.
- Offer granular control over Enigma machine parameters (rotors, reflector, plugboard, stepping mechanism) to adjust encryption complexity.
- Ensure a clear, well-documented, and easy-to-use API for Go developers.
- Emphasize correctness and predictability, while acknowledging that this is a simulation and not intended for production-grade cryptographic security for sensitive data (as modern ciphers are far more robust).
- Maintain the core "Enigma behavior" accurately, including the reciprocal nature and stepping.

### 1.3 Target Audience

- Cryptology enthusiasts and students.
- Developers building educational tools or simulations.
- Researchers interested in historical cipher analysis.
- Anyone needing a customizable text obfuscation tool based on Enigma principles (for non-sensitive data).

### 1.4 Non-Goals

- To create a production-grade cryptographic library for securing highly sensitive data. Modern ciphers (e.g., AES-GCM) should always be preferred for real-world security.
- To implement every historical variant of the Enigma machine (e.g., specific M3/M4 differences, naval vs. army versions, unless explicitly added as configurable options). The focus is on a generalized configurable Enigma.
- To provide a graphical user interface (GUI); this is purely a Go library.

## 2. Features

The eniGOma library will expose an API that allows users to configure and operate a virtual Enigma machine.

### 2.1 Core Enigma Machine Simulation

#### Initialization:
- Ability to create a new Enigma machine instance.
- Configuration through a Config struct or functional options for clear, flexible setup.

#### Encryption/Decryption:
- `Encrypt(plaintext string) (string, error)`: Encrypts a given Unicode string.
- `Decrypt(ciphertext string) (string, error)`: Decrypts a given Unicode string using the current machine state. (Due to the reciprocal nature of Enigma, decryption is the same operation as encryption if the machine state is identical).

#### Reset:
- `Reset()`: Resets the machine's rotor positions to their initial configuration, but keeps the chosen rotors, reflector, and plugboard settings.

### 2.2 Customizable Components

The library will allow detailed configuration of the Enigma's components to adjust complexity.

#### 2.2.1 Character Set (Alphabet)

**Requirement:** Users must be able to define the set of Unicode runes (characters) that the Enigma machine will operate on.

**Implementation Details:**
- A `[]rune` slice or `map[rune]int` could represent the active alphabet.
- All internal mappings (rotors, reflector, plugboard) will be dynamically adjusted to the size of this alphabet.
- Pre-defined common alphabets (e.g., `eniGOma.AlphabetLatinUpper`, `eniGOma.AlphabetASCIIPrintable`) should be provided for convenience.
- **Validation:** The library must ensure the chosen alphabet does not contain duplicate characters.

#### 2.2.2 Rotors

**Requirement:**
- Users can specify the number of rotors to be used (minimum 3, potentially up to 10 or more).
- Users can select specific rotor configurations from a pool of available designs.
- Users can define the initial position (offset) and ring setting for each rotor.
- Users can define "notch" positions for each rotor, controlling when subsequent rotors step.

**Implementation Details:**
- A `Rotor` struct/interface representing a single rotor with `ForwardMap`, `BackwardMap`, and `NotchPositions`.
- `NewRotor(id string, alphabet []rune, forwardMapping string, notches []rune) (Rotor, error)`: Function to create custom rotors. `forwardMapping` would be a string of runes representing the output for each input rune in order of the alphabet.
- `RandomRotor(id string, alphabet []rune) (Rotor, error)`: Generates a cryptographically random, valid rotor permutation and random notch positions.
- The Enigma machine constructor will take a slice of `Rotor` objects.
- Internal logic will handle the stepping (single step, double step, etc.) and permutation logic for each character.

#### 2.2.3 Reflector

**Requirement:**
- Users can select a specific reflector configuration from a pool of available designs.
- Users can optionally generate a random, valid reflector.

**Implementation Details:**
- A `Reflector` struct/interface with a `ReflectMap`.
- `NewReflector(id string, alphabet []rune, mapping string) (Reflector, error)`: Function to create custom reflectors. `mapping` would be a string of runes representing the output for each input rune in order of the alphabet.
- `RandomReflector(id string, alphabet []rune) (Reflector, error)`: Generates a cryptographically random, valid reflector permutation (ensuring no character maps to itself, and all mappings are reciprocal).

#### 2.2.4 Plugboard (Steckerbrett)

**Requirement:**
- Users can specify pairs of Unicode characters to swap on the plugboard.
- Users can clear existing plugboard settings.
- Users can optionally generate a random set of plugboard connections up to a specified number of pairs.

**Implementation Details:**
- A `Plugboard` struct with a `Map` (e.g., `map[rune]rune`).
- `AddPair(r1, r2 rune) error`: Adds a reciprocal swap.
- `RemovePair(r rune) error`: Removes a pair involving the given rune.
- `RandomPairs(n int)`: Generates n random reciprocal pairs.
- **Validation:** Ensure characters are part of the active alphabet and not already part of another pair.

### 2.3 Complexity Levels (Convenience API)

**Requirement:** Provide high-level functions to quickly configure machines for different "security" levels.

**Implementation Details:**
- `NewEnigma(options ...Option) (*Enigma, error)`: Utilize Go's functional options pattern for flexible configuration.
- `WithRandomSettings(level SecurityLevel)`: An option that, based on `SecurityLevel` enum (e.g., Low, Medium, High, Extreme), automatically configures:
  - Number of rotors (e.g., Low=3, Medium=5, High=8, Extreme=10+)
  - Selection of rotors (from a pre-defined or randomly generated pool)
  - Initial rotor positions (random)
  - Ring settings (random)
  - Reflector (randomly chosen or generated)
  - Plugboard connections (e.g., Low=0-5 pairs, Medium=5-15 pairs, High=max unique pairs)
  - Character Set (e.g., Low=A-Z, Medium=ASCII printable, High=user-defined large Unicode range).
- `WithAlphabet(alphabet []rune)`: Option to set the character set.

### 2.4 State Management and Export/Import

**Requirement:** Ability to get and set the full machine configuration and current state.

**Implementation Details:**
- `GetSettings() *EnigmaSettings`: Returns a serializable struct containing all configuration (rotor types, order, initial positions, ring settings, plugboard, reflector, current rotor positions).
- `LoadSettings(settings *EnigmaSettings) error`: Initializes the machine with the provided settings.
- `MarshalJSON()` / `UnmarshalJSON()` methods for `EnigmaSettings` to allow easy saving/loading of machine state. This enables sharing a "key" for encryption/decryption.

## 3. Architecture

- **Modular Design:** The library will be structured into distinct packages or modules for each core component (e.g., enigma, rotor, reflector, plugboard).
- **Interfaces:** Where appropriate, interfaces will be used (e.g., `Rotor` interface) to allow for future extensibility and custom implementations.
- **Stateless Operations for Components:** Individual rotor/reflector operations will be designed as pure functions (or methods that operate on their internal state) to minimize side effects, making them easier to test. The `Enigma` struct will manage the overall state.
- **Error Handling:** Idiomatic Go error handling (error return values) will be used consistently.
- **Unicode Handling:** Internally, `rune` will be used for character manipulation. Input/output will be `string`, with conversion to/from `[]rune` as needed.

## 4. Technical Considerations

### 4.1 Go Language Specifics

- **Go Modules:** Standard Go module structure.
- **Concurrency:** Enigma encryption/decryption is typically a sequential operation. No complex concurrency is expected within the core encryption logic itself. However, the library should be safe for concurrent use if multiple Enigma instances are created and used independently.
- **Performance:** While not a primary goal for production-grade crypto, the implementation should be reasonably performant for typical message lengths. Benchmarks will be created.

### 4.2 Cryptographic Best Practices (for Simulation)

- **Randomness:** When generating random rotors, reflectors, or plugboard settings, `crypto/rand` package must be used to ensure cryptographically secure pseudo-random number generation.
- **Avoid Homebrew Crypto:** While eniGOma is a simulation of a historical cipher, the underlying implementation of permutations and mappings should use standard, proven techniques. Avoid custom bit manipulations that could introduce subtle flaws.
- **Constant-Time Operations (where applicable):** While less critical for a simulation, any lookup tables or permutation logic should ideally aim for constant-time operations to prevent timing side-channel attacks, if practically feasible without excessive complexity. This primarily applies to lookups within a rotor's mapping.

### 4.3 Input/Output

Input and output for encryption/decryption will be `string`.

The library will handle mapping incoming runes to the defined internal alphabet's integer indices and vice-versa. Characters not in the defined alphabet will need a clear error handling strategy (e.g., error out, skip, or substitute with a default unknown character). **Decision:** For simplicity and clear behavior, characters not present in the configured alphabet will result in an error or be skipped/passthrough, depending on a configurable option. An initial decision is to return an error to ensure explicit handling.

## 5. API Design (High-Level Examples)

```go
package eniGOma

import "errors"

// SecurityLevel defines pre-set complexity levels.
type SecurityLevel int

const (
	Low SecurityLevel = iota
	Medium
	High
	Extreme
)

// Enigma represents a configurable Enigma machine.
type Enigma struct {
	// internal fields for rotors, reflector, plugboard, current positions etc.
}

// Option is a functional option for Enigma configuration.
type Option func(*Enigma) error

// NewEnigma creates a new Enigma machine with the given options.
// Example: eni, err := eniGOma.NewEnigma(eniGOma.WithRandomSettings(eniGOma.High))
// Example: eni, err := eniGOma.NewEnigma(eniGOma.WithAlphabet(eniGOma.AlphabetLatinUpper))
func NewEnigma(opts ...Option) (*Enigma, error) {
	// ... implementation ...
}

// WithAlphabet sets the character set for the Enigma machine.
// All rotors, plugboard, and reflector will be built/validated against this alphabet.
func WithAlphabet(alphabet []rune) Option {
	return func(e *Enigma) error {
		// ... set alphabet and re-init internal mappings ...
	}
}

// WithRandomSettings configures the Enigma with random components based on a security level.
// This is a convenience for quickly setting up a machine.
func WithRandomSettings(level SecurityLevel) Option {
	return func(e *Enigma) error {
		// ... logic to configure rotors, reflector, plugboard based on level ...
	}
}

// WithCustomComponents allows detailed manual configuration of components.
// Example: eniGOma.WithCustomComponents(rotors, reflector, plugboard)
func WithCustomComponents(rotors []Rotor, reflector Reflector, plugboard *Plugboard) Option {
	return func(e *Enigma) error {
		// ... set components, perform validation ...
	}
}

// Encrypt encrypts the given plaintext using the current machine state.
func (e *Enigma) Encrypt(plaintext string) (string, error) {
	// ... implementation ...
}

// Decrypt decrypts the given ciphertext using the current machine state.
// (Identical to Encrypt in behavior for Enigma).
func (e *Enigma) Decrypt(ciphertext string) (string, error) {
	// ... implementation ...
}

// Reset resets the rotor positions to their initial configuration.
func (e *Enigma) Reset() {
	// ... implementation ...
}

// GetSettings returns the current configuration and state of the Enigma machine.
// This struct can be serialized to share the "key".
type EnigmaSettings struct {
	Alphabet              []rune
	RotorSpecs            []RotorSpec // Specifies ID, initial position, ring setting, notches
	ReflectorSpec         ReflectorSpec
	PlugboardPairs        map[rune]rune
	CurrentRotorPositions []int // Current positions of rotors for ongoing encryption
}

// LoadSettings initializes the Enigma machine with the provided settings.
func (e *Enigma) LoadSettings(settings *EnigmaSettings) error {
	// ... implementation ...
}

// Rotor represents a single rotor with its internal wiring and notch positions.
type Rotor interface {
	ID() string
	Forward(inputIdx int) int
	Backward(inputIdx int) int
	IsAtNotch() bool
	Step()
	SetPosition(pos int)
	SetRingSetting(ring int)
	// ... potentially other methods ...
}

// Reflector represents the reflector component.
type Reflector interface {
	ID() string
	Reflect(inputIdx int) int
}

// Plugboard represents the plugboard.
type Plugboard struct {
	// ... internal map ...
}

// AddPair adds a reciprocal swap on the plugboard.
func (p *Plugboard) AddPair(r1, r2 rune) error { /* ... */ }

// RemovePair removes a pair involving the given rune.
func (p *Plugboard) RemovePair(r rune) error { /* ... */ }

// Clear clears all plugboard connections.
func (p *Plugboard) Clear() { /* ... */ }

// Process applies the plugboard mapping to a rune.
func (p *Plugboard) Process(r rune) rune { /* ... */ }

// Pre-defined alphabets
var (
	AlphabetLatinUpper     = []rune{'A', 'B', 'C', /* ... */ 'Z'}
	AlphabetASCIIPrintable = []rune{/* ... all printable ASCII ... */}
	// etc.
)
```

## 6. Future Considerations

- **Performance Optimizations:** For very large Unicode alphabets, consider using lookup tables (arrays) instead of `map[rune]int` for character-to-index and index-to-character conversions if performance becomes a bottleneck.

- **Error Handling for Invalid Characters:** Refine the behavior when input text contains characters not in the configured alphabet. Options include returning an error, silently skipping them, or defining a "default unknown character" mapping. The current PRD leans towards an error, which is generally safer.

- **More Advanced Stepping:** While eniGOma allows defining notches, more complex historical stepping mechanisms (e.g., specific double-stepping behavior of some Enigma models) could be added as advanced options or specific Rotor implementations.

- **Visualizer Integration:** While not part of the library, consider how the `EnigmaSettings` structure could facilitate easy integration with a separate visualization tool.

- **Benchmarking:** Thorough benchmarks to understand performance characteristics with different alphabet sizes and rotor counts.

## 7. Success Metrics

- **API Usability:** Positive feedback from developers regarding ease of integration and clarity of documentation.

- **Correctness:** All test cases (including historical Enigma examples and new Unicode examples) pass consistently.

- **Configurability:** Users can effectively customize the machine parameters to achieve desired complexity levels.

- **Maintainability:** Codebase adheres to Go best practices, making it easy to understand, test, and extend.