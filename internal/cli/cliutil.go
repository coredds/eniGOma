package cli

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetInputText reads input text from a file or stdin.
func GetInputText(filePath string) (string, error) {
	if filePath == "-" {
		// Read from stdin
		info, err := os.Stdin.Stat()
		if err != nil {
			return "", fmt.Errorf("failed to stat stdin: %w", err)
		}
		if (info.Mode() & os.ModeCharDevice) != 0 {
			return "", fmt.Errorf("stdin is not a pipe")
		}
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read from stdin: %w", err)
		}
		return string(input), nil
	}

	// Read from file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	return string(data), nil
}

// FormatOutput formats the output text based on the specified format.
func FormatOutput(text, format string) (string, error) {
	switch strings.ToLower(format) {
	case "hex":
		return fmt.Sprintf("%x", text), nil
	case "base64":
		return base64.StdEncoding.EncodeToString([]byte(text)), nil
	default:
		return text, nil
	}
}

// ParseInputFormat parses the input text based on the specified format.
func ParseInputFormat(text, format string) (string, error) {
	switch strings.ToLower(format) {
	case "hex":
		decoded, err := hex.DecodeString(text)
		if err != nil {
			return "", fmt.Errorf("failed to decode hex: %w", err)
		}
		return string(decoded), nil
	case "base64":
		decoded, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			return "", fmt.Errorf("failed to decode base64: %w", err)
		}
		return string(decoded), nil
	default:
		return text, nil
	}
}

// WriteOutput writes the output text to a file or stdout.
func WriteOutput(text, filePath string) error {
	if filePath == "-" {
		// Write to stdout
		fmt.Println(text)
		return nil
	}

	// Write to file
	return os.WriteFile(filePath, []byte(text), 0600)
}
