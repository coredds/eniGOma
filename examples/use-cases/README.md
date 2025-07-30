# Use Case Examples

This directory contains example configurations tailored for specific real-world applications and scenarios.

## Available Examples

### 📄 `document-protection.json` - File and Document Encryption
- **Security Level**: Low-Medium
- **Best For**: Personal documents, files, notes
- **Features**: Balanced security and performance for everyday use

```bash
# Encrypt a document
eniGOma encrypt --file important-document.txt --config examples/use-cases/document-protection.json --output encrypted-doc.txt

# Decrypt back to original
eniGOma decrypt --file encrypted-doc.txt --config examples/use-cases/document-protection.json --output decrypted-doc.txt
```

### 🔐 `secure-communication.json` - Secure Messaging
- **Security Level**: High
- **Best For**: Confidential communications, secure messaging
- **Features**: Strong security for sensitive communications

```bash
# Encrypt a message
eniGOma encrypt --text "Meet at the usual place at 3pm" --config examples/use-cases/secure-communication.json

# For secure file transfer
eniGOma encrypt --file confidential-report.pdf --config examples/use-cases/secure-communication.json
```

### 🏛️ `historical-simulation.json` - Educational and Historical
- **Security Level**: Historical accuracy
- **Best For**: Learning about Enigma, historical projects, education
- **Features**: Authentic WWII-era Enigma M3 configuration

```bash
# Historical encryption simulation
eniGOma encrypt --text "WEATHER REPORT CLOUDY" --config examples/use-cases/historical-simulation.json

# Educational demonstration
eniGOma config --show examples/use-cases/historical-simulation.json --detailed
```

## Detailed Use Case Scenarios

### 🗂️ Personal File Protection

**Scenario**: Protecting personal documents, photos, or sensitive files on your computer.

**Recommended Configuration**: `document-protection.json`

```bash
# Protect a folder of documents
for file in Documents/*.txt; do
    eniGOma encrypt --file "$file" --config examples/use-cases/document-protection.json --output "encrypted-${file##*/}"
done

# Create encrypted backup
tar -czf backup.tar.gz Documents/
eniGOma encrypt --file backup.tar.gz --config examples/use-cases/document-protection.json --output secure-backup.enc
```

### 💬 Secure Communication

**Scenario**: Sending confidential messages or files that need protection during transit.

**Recommended Configuration**: `secure-communication.json`

**Workflow**:
1. Both parties share the same configuration file securely
2. Sender encrypts message/file
3. Encrypted content sent via any channel
4. Receiver decrypts using shared configuration

```bash
# Sender
eniGOma encrypt --text "The quarterly reports are ready for review" \
    --config examples/use-cases/secure-communication.json > encrypted-message.txt

# Receiver  
cat encrypted-message.txt | eniGOma decrypt --config examples/use-cases/secure-communication.json
```

### 🎓 Educational Projects

**Scenario**: Teaching cryptography, computer science, or WWII history.

**Recommended Configuration**: `historical-simulation.json`

**Learning Activities**:
```bash
# Demonstrate Enigma principles
eniGOma encrypt --text "ENIGMA" --config examples/use-cases/historical-simulation.json

# Show rotor movement
eniGOma config --test examples/use-cases/historical-simulation.json --text "AAAAA"

# Compare with modern security
eniGOma preset --describe classic
eniGOma preset --describe extreme
```

## Advanced Use Cases

### 📊 Data Processing Pipeline

```bash
# Encrypt CSV data for processing
eniGOma encrypt --file sales-data.csv --config examples/use-cases/document-protection.json --output encrypted-sales.csv

# Process encrypted data (theoretical - you'd decrypt first)
# ... data processing ...

# Re-encrypt results
eniGOma encrypt --file processed-results.csv --config examples/use-cases/document-protection.json
```

### 🔄 Configuration Rotation

```bash
# Daily rotation for high-security communications
eniGOma keygen --security high --output "daily-key-$(date +%Y%m%d).json"

# Weekly rotation for document protection  
eniGOma keygen --security medium --output "weekly-key-$(date +%Y-W%U).json"
```

### 🧪 Research and Testing

```bash
# Generate test configurations for research
for level in low medium high extreme; do
    eniGOma keygen --security $level --output "research-${level}.json"
    eniGOma config --validate "research-${level}.json" --stats
done
```

## Integration Examples

### Shell Script Integration

```bash
#!/bin/bash
# secure-backup.sh

CONFIG="examples/use-cases/document-protection.json"
BACKUP_DIR="$HOME/Documents"
ENCRYPTED_DIR="$HOME/secure-backups"

# Create encrypted backup
tar -czf temp-backup.tar.gz "$BACKUP_DIR"
eniGOma encrypt --file temp-backup.tar.gz --config "$CONFIG" --output "$ENCRYPTED_DIR/backup-$(date +%Y%m%d).enc"
rm temp-backup.tar.gz

echo "Secure backup created: $ENCRYPTED_DIR/backup-$(date +%Y%m%d).enc"
```

### Python Integration

```python
import subprocess
import json

def encrypt_with_enigma(text, config_file):
    """Encrypt text using eniGOma configuration"""
    result = subprocess.run([
        'eniGOma', 'encrypt', 
        '--text', text,
        '--config', config_file
    ], capture_output=True, text=True)
    
    return result.stdout.strip()

# Usage
encrypted = encrypt_with_enigma(
    "Sensitive data here", 
    "examples/use-cases/secure-communication.json"
)
```

## Performance Guidelines

### File Size Recommendations

| Use Case | File Size | Recommended Config | Expected Time |
|----------|-----------|-------------------|---------------|
| Text files | < 1MB | `document-protection.json` | < 1 second |
| Documents | 1-10MB | `document-protection.json` | 1-10 seconds |
| Secure comms | < 100KB | `secure-communication.json` | < 1 second |
| Large files | > 10MB | Consider chunking | Varies |

### Security vs Performance Trade-offs

- **Document Protection**: Optimized for daily use, good security
- **Secure Communication**: Higher security, acceptable performance impact  
- **Historical Simulation**: Fast processing, educational accuracy

## Best Practices

### 🔑 Key Management
- Never share configuration files over insecure channels
- Use different configurations for different purposes
- Rotate configurations periodically for high-security use cases

### 🛡️ Security Considerations
- ⚠️ eniGOma is for educational/simulation purposes only
- For production security, use AES-GCM or ChaCha20-Poly1305
- Test configurations thoroughly before deployment

### 📋 Operational Guidelines
- Always validate configurations before use
- Keep backups of your configuration files
- Document which configurations are used for what purposes

---

*These use case examples demonstrate practical applications of eniGOma while maintaining appropriate security practices.*