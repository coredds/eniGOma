# Security Level Examples

This directory contains example configurations demonstrating different security levels available in eniGOma.

## Available Examples

### 🔓 `classic-low-security.json` - Classic Historical
- **Security Level**: Low
- **Rotors**: 3 (historically accurate)
- **Plugboard Pairs**: 2
- **Alphabet**: Latin uppercase (A-Z)
- **Use Case**: Educational, historical simulation, learning Enigma mechanics

```bash
# Example usage
eniGOma encrypt --text "HELLO WORLD" --config examples/security-levels/classic-low-security.json
```

### 🔒 `simple-medium-security.json` - Balanced Security
- **Security Level**: Medium
- **Rotors**: 5
- **Plugboard Pairs**: 8
- **Alphabet**: Latin uppercase (A-Z)
- **Use Case**: General purpose encryption, file protection

```bash
# Example usage
eniGOma encrypt --file document.txt --config examples/security-levels/simple-medium-security.json
```

### 🔒🔒 `high-security.json` - Strong Protection
- **Security Level**: High
- **Rotors**: 8
- **Plugboard Pairs**: 15
- **Alphabet**: Latin uppercase (A-Z)
- **Use Case**: Sensitive documents, secure communication

```bash
# Example usage
eniGOma encrypt --text "CONFIDENTIAL DATA" --config examples/security-levels/high-security.json
```

### 🔒🔒🔒 `extreme-key.json` - Maximum Complexity
- **Security Level**: Extreme
- **Rotors**: 12
- **Plugboard Pairs**: 20+
- **Alphabet**: Latin uppercase (A-Z)
- **Use Case**: Research, maximum obfuscation needs

```bash
# Example usage
eniGOma encrypt --text "TOP SECRET" --config examples/security-levels/extreme-key.json
```

## Security Level Comparison

| Level | Rotors | Plugboard | Keyspace Size | Performance | Use Case |
|-------|--------|-----------|---------------|-------------|----------|
| Low | 3 | 2 | ~10^15 | Fastest | Learning, historical |
| Medium | 5 | 8 | ~10^25 | Fast | General use |
| High | 8 | 15 | ~10^40 | Moderate | Sensitive data |
| Extreme | 12 | 20+ | ~10^60 | Slower | Maximum security |

## Generating Your Own

You can generate configurations at different security levels:

```bash
# Generate a custom high-security configuration
eniGOma keygen --security high --output my-high-security.json

# Generate with specific parameters
eniGOma keygen --rotors 6 --plugboard-pairs 10 --output custom-config.json
```

## Validation and Testing

Always validate your configurations:

```bash
# Validate configuration
eniGOma config --validate examples/security-levels/high-security.json

# Test with sample text
eniGOma config --test examples/security-levels/high-security.json --text "TEST MESSAGE"

# Show detailed configuration info
eniGOma config --show examples/security-levels/extreme-key.json --detailed
```

## Historical Context

The **Classic** configuration closely matches the Wehrmacht M3 Enigma used during WWII:
- 3 rotors (I, II, III historically)
- Minimal plugboard connections
- 26-character Latin alphabet
- Deterministic rotor stepping

Higher security levels represent modern enhancements that would not have been possible with 1940s technology.

## Performance Notes

- **Low/Medium**: Suitable for real-time applications
- **High**: Good for batch processing
- **Extreme**: Best for offline encryption of critical data

Remember: For production security needs, use modern cryptographic algorithms like AES-GCM or ChaCha20-Poly1305.