# eniGOma Usage Guide

## Auto-Detection (since v0.3.0)

The easiest way to use eniGOma: just encrypt any text in any language.

```bash
# Works automatically with any language!
eniGOma encrypt --text "Olá mundo! Como você está?"
eniGOma encrypt --text "Mixed: Hello! Привет! 日本語!"
eniGOma encrypt --text "Symbols: αβγ δεζ 🙂 test!"
```

Note: The default alphabet is auto-detected (equivalent to --alphabet=auto). No need to specify alphabets; eniGOma automatically detects the optimal character set from your text.

## Brazilian Portuguese Support

eniGOma includes **built-in support for Brazilian Portuguese** with both auto-detection and the `AlphabetPortuguese` predefined alphabet.

### Quick Start with Portuguese

```go
package main

import (
    "fmt"
    "log"
    "github.com/coredds/eniGOma"
    "github.com/coredds/eniGOma/pkg/enigma"
)

func main() {
    // Create Enigma machine with Portuguese alphabet
    machine, err := enigma.NewEnigmaSimple(eniGOma.AlphabetPortuguese)
    if err != nil {
        log.Fatal(err)
    }

    // Your message with full Portuguese accents
    message := "Hoje eu fui almoçar na casa da vovó."
    
    // Encrypt
    encrypted, err := machine.Encrypt(message)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Encrypted: %s\n", encrypted)

    // Decrypt
    machine.Reset()
    decrypted, err := machine.Decrypt(encrypted)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decrypted: %s\n", decrypted)
    
    // Perfect round-trip with all accents preserved!
}
```

### Portuguese Alphabet Features

- **88 characters total** (even number for reflector compatibility)
- **Full accent support**: à, á, â, ã, ç, é, ê, í, ó, ô, õ, ú (uppercase and lowercase)
- **Complete Latin alphabet**: A-Z, a-z
- **Common punctuation**: space, period, comma, exclamation, question mark, etc.

### Supported Characters

The `AlphabetPortuguese` includes:

```
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
ÀÁÂÃÇÉÊÍÓÔÕÚàáâãçéêíóôõú .,!?;:-'"()
```

### Test Your Portuguese Text

```go
// All these phrases work perfectly:
phrases := []string{
    "Olá, como você está?",
    "Bom dia! Está tudo bem?", 
    "Vamos à praia amanhã.",
    "Não posso ir à reunião.",
    "São Paulo é uma cidade incrível!",
    "Brasília é a capital do Brasil.",
    "Açúcar, café e pão de açúcar.",
}
```

## Other Languages

### Greek Example
```go
machine, err := enigma.NewEnigmaSimple(eniGOma.AlphabetGreek)
message := "Αβγδε ζητα" // Greek text
```

### Custom Language Support
```go
// Create alphabet for any language
customAlphabet := []rune{
    'A', 'B', 'C', /* your characters */
    'ñ', 'ü', 'ß', /* special characters */
    ' ', '.', '!',  /* punctuation */
}

machine, err := enigma.New(enigma.WithAlphabet(customAlphabet))
```

## Advanced Portuguese Usage

### Security Levels with Portuguese
```go
machine, err := enigma.New(
    enigma.WithAlphabet(eniGOma.AlphabetPortuguese),
    enigma.WithRandomSettings(enigma.High), // High security
)
```

### Save/Load Portuguese Settings
```go
// Save Portuguese machine configuration
jsonData, err := machine.SaveSettingsToJSON()

// Later: restore exact same configuration
newMachine, err := enigma.NewFromJSON(jsonData)
```

## Example Configurations

The [`examples/`](./examples/) directory contains ready-to-use configuration files:

### Portuguese Examples
```bash
# Use the Portuguese configuration example
eniGOma encrypt --text "Bom dia, Brasil!" --auto-config pt-key.json
eniGOma decrypt --text "ENCRYPTED" --config pt-key.json

# Generate your own Portuguese configuration
eniGOma keygen --alphabet portuguese --security medium --output my-portuguese.json
```

### Other Language Examples
```bash
# Greek text encryption
eniGOma encrypt --text "Γεια σας!" --config examples/languages/greek-simple.json

# Mixed character sets
eniGOma encrypt --text "Password123!" --config examples/languages/mixed-alphabet-extreme.json
```

### Security Examples
```bash
# Document protection
eniGOma encrypt --file document.txt --config examples/use-cases/document-protection.json

# High security communication
eniGOma encrypt --text "TOP SECRET" --config examples/security-levels/extreme-key.json

# Historical simulation
eniGOma encrypt --text "ENIGMA MACHINE" --config examples/use-cases/historical-simulation.json
```

Browse all examples: [`examples/README.md`](./examples/README.md)

---

Brazilian Portuguese is now a first-class citizen in eniGOma.

## CLI: Stdin and Encoding Examples

```bash
# Stdin encryption (auto-detected alphabet by default)
echo "Hello via stdin" | eniGOma encrypt --auto-config my-key.json

# Base64 output and decrypt
eniGOma encrypt --text "Hello" --auto-config my-key.json --format base64
eniGOma decrypt --text "SGVsbG8=" --config my-key.json --format base64

# Hex output and decrypt
eniGOma encrypt --text "Hello" --auto-config my-key.json --format hex
eniGOma decrypt --text "48656c6c6f" --config my-key.json --format hex
```