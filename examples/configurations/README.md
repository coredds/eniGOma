# Advanced Configuration Examples

This directory contains examples of advanced Enigma machine configurations demonstrating custom components, complex setups, and specialized configurations.

## Coming Soon

This directory will contain:

### 🔧 Custom Component Examples
- Custom rotor specifications
- Advanced reflector configurations  
- Complex plugboard setups
- Multi-stage encryption configurations

### ⚙️ Specialized Configurations
- High-performance setups
- Research configurations
- Compatibility configurations
- Edge case handling

### 🧬 Experimental Features
- Custom alphabet implementations
- Advanced security patterns
- Performance optimizations
- Integration examples

## Generating Advanced Configurations

You can create custom configurations using the CLI:

```bash
# Generate with specific parameters
enigoma keygen --rotors 6 --plugboard-pairs 15 --security high --output custom-advanced.json

# Create from preset and modify
enigoma preset --export extreme --output base-config.json
# Then manually edit the JSON for custom requirements
```

## Manual Configuration

For completely custom setups, you can manually create configuration files following this structure:

```json
{
  "alphabet": "your custom character set",
  "rotor_specs": [
    {
      "id": "CustomRotor1",
      "forward_mapping": "custom mapping string matching alphabet length",
      "notches": [array of notch positions],
      "position": 0,
      "ring_setting": 0
    }
  ],
  "reflector_spec": {
    "id": "CustomReflector",
    "mapping": "reflector mapping (must be reciprocal)"
  },
  "plugboard_pairs": {
    "A": "Z",
    "Z": "A"
  },
  "current_rotor_positions": [0, 0, 0]
}
```

## Validation

Always validate custom configurations:

```bash
enigoma config --validate your-custom-config.json --detailed
```

---

*Check back soon for comprehensive advanced configuration examples!*