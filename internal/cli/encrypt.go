// Package cli provides the encrypt command for the eniGOma CLI.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/coredds/eniGOma"
	"github.com/coredds/eniGOma/pkg/enigma"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt text or files using an Enigma machine",
	Long: `Encrypt plaintext using a configured Enigma machine.

You can encrypt text directly, read from a file, or read from stdin.
The machine can be configured using presets, custom settings, or configuration files.

Examples:
  eniGOma encrypt --text "Hello World" --preset classic
  eniGOma encrypt --file input.txt --output encrypted.txt --preset high
  eniGOma encrypt --text "Secret Message" --alphabet greek --security medium
  eniGOma encrypt --file data.txt --config my-enigma.json`,
	RunE: runEncrypt,
}

func init() {
	// Input options
	encryptCmd.Flags().StringP("text", "t", "", "Text to encrypt")
	encryptCmd.Flags().StringP("file", "f", "", "File to encrypt")
	encryptCmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")

	// Machine configuration
	encryptCmd.Flags().StringP("preset", "p", "", "Use a preset configuration (classic, simple, high, extreme)")
	encryptCmd.Flags().StringP("alphabet", "a", "latin", "Alphabet to use (latin, greek, cyrillic, portuguese, ascii, alphanumeric)")
	encryptCmd.Flags().StringP("security", "s", "medium", "Security level (low, medium, high, extreme)")

	// Advanced options
	encryptCmd.Flags().StringSliceP("rotors", "r", nil, "Rotor positions (e.g., 1,5,12)")
	encryptCmd.Flags().StringSliceP("plugboard", "", nil, "Plugboard pairs (e.g., A:Z,B:Y)")
	encryptCmd.Flags().BoolP("reset", "", false, "Reset machine to initial state before encryption")

	// Output formatting
	encryptCmd.Flags().StringP("format", "", "text", "Output format (text, hex, base64)")
	encryptCmd.Flags().BoolP("preserve-case", "", false, "Preserve original case (when possible)")
}

func runEncrypt(cmd *cobra.Command, args []string) error {
	setupVerbose(cmd)

	// Get input text
	text, err := getInputText(cmd)
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

	// Encrypt text
	encrypted, err := machine.Encrypt(text)
	if err != nil {
		return fmt.Errorf("encryption failed: %v", err)
	}

	// Format output
	formatted, err := formatOutput(encrypted, cmd)
	if err != nil {
		return fmt.Errorf("failed to format output: %v", err)
	}

	// Write output
	return writeOutput(formatted, cmd)
}

func getInputText(cmd *cobra.Command) (string, error) {
	// Check for direct text input
	if text, _ := cmd.Flags().GetString("text"); text != "" {
		return text, nil
	}

	// Check for file input
	if filename, _ := cmd.Flags().GetString("file"); filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %v", filename, err)
		}
		return string(data), nil
	}

	// TODO: Add stdin support in future version
	return "", nil
}

func createMachineFromFlags(cmd *cobra.Command) (*enigma.Enigma, error) {
	// Check if config file is specified
	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		return createMachineFromConfig(configFile)
	}

	// Check for preset
	if preset, _ := cmd.Flags().GetString("preset"); preset != "" {
		return createMachineFromPreset(preset)
	}

	// Create machine from individual flags
	return createMachineFromSettings(cmd)
}

func createMachineFromConfig(configFile string) (*enigma.Enigma, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	return enigma.NewFromJSON(string(data))
}

func createMachineFromPreset(preset string) (*enigma.Enigma, error) {
	switch strings.ToLower(preset) {
	case "classic":
		return enigma.NewEnigmaClassic()
	case "simple":
		return enigma.NewEnigmaSimple(eniGOma.AlphabetLatinUpper)
	case "low":
		return enigma.New(
			enigma.WithAlphabet(eniGOma.AlphabetLatinUpper),
			enigma.WithRandomSettings(enigma.Low),
		)
	case "medium":
		return enigma.New(
			enigma.WithAlphabet(eniGOma.AlphabetLatinUpper),
			enigma.WithRandomSettings(enigma.Medium),
		)
	case "high":
		return enigma.New(
			enigma.WithAlphabet(eniGOma.AlphabetLatinUpper),
			enigma.WithRandomSettings(enigma.High),
		)
	case "extreme":
		return enigma.New(
			enigma.WithAlphabet(eniGOma.AlphabetLatinUpper),
			enigma.WithRandomSettings(enigma.Extreme),
		)
	default:
		return nil, fmt.Errorf("unknown preset: %s. Available: classic, simple, low, medium, high, extreme", preset)
	}
}

func createMachineFromSettings(cmd *cobra.Command) (*enigma.Enigma, error) {
	// Get alphabet
	alphabet, err := getAlphabetFromFlag(cmd)
	if err != nil {
		return nil, err
	}

	// Get security level
	securityLevel, err := getSecurityLevelFromFlag(cmd)
	if err != nil {
		return nil, err
	}

	// Create machine with basic settings
	machine, err := enigma.New(
		enigma.WithAlphabet(alphabet),
		enigma.WithRandomSettings(securityLevel),
	)
	if err != nil {
		return nil, err
	}

	// Apply rotor positions if specified
	if rotorPositions, _ := cmd.Flags().GetStringSlice("rotors"); len(rotorPositions) > 0 {
		positions, err := parseRotorPositions(rotorPositions)
		if err != nil {
			return nil, fmt.Errorf("invalid rotor positions: %v", err)
		}
		if err := machine.SetRotorPositions(positions); err != nil {
			return nil, fmt.Errorf("failed to set rotor positions: %v", err)
		}
	}

	return machine, nil
}

func getAlphabetFromFlag(cmd *cobra.Command) ([]rune, error) {
	alphabetName, _ := cmd.Flags().GetString("alphabet")

	switch strings.ToLower(alphabetName) {
	case "latin", "latin-upper":
		return eniGOma.AlphabetLatinUpper, nil
	case "latin-lower":
		return eniGOma.AlphabetLatinLower, nil
	case "greek":
		return eniGOma.AlphabetGreek, nil
	case "cyrillic":
		return eniGOma.AlphabetCyrillic, nil
	case "portuguese":
		return eniGOma.AlphabetPortuguese, nil
	case "ascii":
		return eniGOma.AlphabetASCIIPrintable, nil
	case "alphanumeric":
		return eniGOma.AlphabetAlphaNumeric, nil
	case "digits":
		return eniGOma.AlphabetDigits, nil
	default:
		return nil, fmt.Errorf("unknown alphabet: %s. Available: latin, greek, cyrillic, portuguese, ascii, alphanumeric, digits", alphabetName)
	}
}

func getSecurityLevelFromFlag(cmd *cobra.Command) (enigma.SecurityLevel, error) {
	securityName, _ := cmd.Flags().GetString("security")

	switch strings.ToLower(securityName) {
	case "low":
		return enigma.Low, nil
	case "medium":
		return enigma.Medium, nil
	case "high":
		return enigma.High, nil
	case "extreme":
		return enigma.Extreme, nil
	default:
		return enigma.Medium, fmt.Errorf("unknown security level: %s. Available: low, medium, high, extreme", securityName)
	}
}

func parseRotorPositions(positions []string) ([]int, error) {
	result := make([]int, len(positions))
	for i, pos := range positions {
		var err error
		result[i], err = parseIntFromString(pos)
		if err != nil {
			return nil, fmt.Errorf("invalid position '%s': %v", pos, err)
		}
	}
	return result, nil
}

func parseIntFromString(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(strings.TrimSpace(s), "%d", &result)
	return result, err
}

func formatOutput(text string, cmd *cobra.Command) (string, error) {
	format, _ := cmd.Flags().GetString("format")

	switch strings.ToLower(format) {
	case "text", "":
		return text, nil
	case "hex":
		return fmt.Sprintf("%x", []byte(text)), nil
	case "base64":
		// Simple base64-like encoding (for demo purposes)
		return fmt.Sprintf("base64:%s", text), nil
	default:
		return "", fmt.Errorf("unknown format: %s. Available: text, hex, base64", format)
	}
}

func writeOutput(text string, cmd *cobra.Command) error {
	outputFile, _ := cmd.Flags().GetString("output")

	if outputFile == "" {
		fmt.Print(text)
		return nil
	}

	return os.WriteFile(outputFile, []byte(text), 0644)
}
