// Package cli provides the decrypt command for the eniGOma CLI.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"fmt"
    "encoding/base64"
    "encoding/hex"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt text or files using an Enigma machine",
	Long: `Decrypt ciphertext using a configured Enigma machine.

Due to the reciprocal nature of the Enigma machine, decryption uses the same
process as encryption. The machine must be configured with the same settings
that were used for encryption.

Examples:
  # Decrypt using the same configuration file used for encryption
  eniGOma decrypt --text "ENCRYPTED_TEXT" --config my-key.json
  eniGOma decrypt --file encrypted.txt --config my-key.json --output decrypted.txt
  
  # Legacy mode: specify alphabet manually (not recommended)
  eniGOma decrypt --text "CIPHER" --alphabet latin --security medium
  
Note: Always use the same configuration file that was used for encryption.`,
	RunE: runDecrypt,
}

func init() {
	// Input options
	decryptCmd.Flags().StringP("text", "t", "", "Text to decrypt")
	decryptCmd.Flags().StringP("file", "f", "", "File to decrypt")
	decryptCmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")

	// Machine configuration
	decryptCmd.Flags().StringP("preset", "p", "", "Use a preset configuration (classic, simple, high, extreme)")
	decryptCmd.Flags().StringP("alphabet", "a", "auto", "Alphabet to use (auto, latin, greek, cyrillic, portuguese, ascii, alphanumeric)")
	decryptCmd.Flags().StringP("security", "s", "medium", "Security level (low, medium, high, extreme)")

	// Advanced options
	decryptCmd.Flags().StringSliceP("rotors", "r", nil, "Rotor positions (e.g., 1,5,12)")
	decryptCmd.Flags().StringSliceP("plugboard", "", nil, "Plugboard pairs (e.g., A:Z,B:Y)")
	decryptCmd.Flags().BoolP("reset", "", false, "Reset machine to initial state before decryption")

	// Input format
	decryptCmd.Flags().StringP("format", "", "text", "Input format (text, hex, base64)")
}

func runDecrypt(cmd *cobra.Command, args []string) error {
	setupVerbose(cmd)

	// Get input text
	text, err := getInputTextForDecrypt(cmd)
	if err != nil {
		return fmt.Errorf("failed to get input text: %v", err)
	}

	if text == "" {
		return fmt.Errorf("no input text provided. Use --text, --file, or pipe to stdin")
	}

	// Create Enigma machine
	machine, err := createMachineFromFlags(cmd)
	if err != nil {
		return fmt.Errorf("failed to create Enigma machine: %v", err)
	}

	// Reset machine if requested
	if reset, _ := cmd.Flags().GetBool("reset"); reset {
		if err := machine.Reset(); err != nil {
			return fmt.Errorf("failed to reset machine: %v", err)
		}
	}

	// Decrypt text (same as encrypt due to Enigma's reciprocal nature)
	decrypted, err := machine.Decrypt(text)
	if err != nil {
		return fmt.Errorf("decryption failed: %v", err)
	}

	// Write output (decrypt always outputs as text)
	return writeOutput(decrypted, cmd)
}

func getInputTextForDecrypt(cmd *cobra.Command) (string, error) {
	// Check for direct text input
	if text, _ := cmd.Flags().GetString("text"); text != "" {
		return parseInputFormat(text, cmd)
	}

	// Check for file input
	if filename, _ := cmd.Flags().GetString("file"); filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %v", filename, err)
		}
		return parseInputFormat(string(data), cmd)
	}

	return "", nil
}

func parseInputFormat(text string, cmd *cobra.Command) (string, error) {
	format, _ := cmd.Flags().GetString("format")

	switch strings.ToLower(format) {
	case "text", "":
		return text, nil
	case "hex":
        decoded, err := hex.DecodeString(strings.TrimSpace(text))
        if err != nil {
            return "", fmt.Errorf("invalid hex input: %w", err)
        }
        return string(decoded), nil
	case "base64":
        decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(text))
        if err != nil {
            return "", fmt.Errorf("invalid base64 input: %w", err)
        }
        return string(decoded), nil
	default:
		return "", fmt.Errorf("unknown format: %s. Available: text, hex, base64", format)
	}
}
