// Package cli provides unit tests for the eniGOma CLI.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// TestRootCommand tests the basic root command functionality.
func TestRootCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "version flag",
			args:     []string{"--version"},
			wantErr:  false,
			contains: "0.2.1",
		},
		{
			name:     "help flag",
			args:     []string{"--help"},
			wantErr:  false,
			contains: "A highly customizable, Unicode-capable Enigma machine",
		},
		{
			name:     "invalid command",
			args:     []string{"invalid-command"},
			wantErr:  true,
			contains: "unknown command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture output
			var out bytes.Buffer

			// Create a new root command for testing
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := out.String()
			if tt.contains != "" && !strings.Contains(output, tt.contains) {
				t.Errorf("Output should contain '%s', got: %s", tt.contains, output)
			}
		})
	}
}

// TestEncryptCommand tests the encrypt command functionality.
func TestEncryptCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		setup   func(t *testing.T) string // Returns temp file path if needed
		cleanup func(string)
	}{
		{
			name:    "encrypt with text and preset",
			args:    []string{"encrypt", "--text", "HELLO", "--preset", "classic"},
			wantErr: false,
		},
		{
			name:    "encrypt with alphabet",
			args:    []string{"encrypt", "--text", "HELLO", "--alphabet", "latin", "--security", "low"},
			wantErr: false,
		},
		{
			name:    "encrypt with invalid preset",
			args:    []string{"encrypt", "--text", "HELLO", "--preset", "invalid"},
			wantErr: true,
		},
		{
			name:    "encrypt without input",
			args:    []string{"encrypt", "--preset", "classic"},
			wantErr: true,
		},
		{
			name:    "encrypt with file input",
			args:    []string{"encrypt", "--file", "", "--preset", "classic"},
			wantErr: false,
			setup: func(t *testing.T) string {
				tmpFile, err := os.CreateTemp("", "test-input-*.txt")
				if err != nil {
					t.Fatalf("Failed to create temp file: %v", err)
				}
				_, err = tmpFile.WriteString("HELLOWORLD")
				if err != nil {
					t.Fatalf("Failed to write to temp file: %v", err)
				}
				tmpFile.Close()
				return tmpFile.Name()
			},
			cleanup: func(path string) {
				os.Remove(path)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tempFile string
			if tt.setup != nil {
				tempFile = tt.setup(t)
				// Replace empty file path with actual temp file
				for i, arg := range tt.args {
					if arg == "--file" && i+1 < len(tt.args) && tt.args[i+1] == "" {
						tt.args[i+1] = tempFile
					}
				}
			}

			if tt.cleanup != nil {
				defer tt.cleanup(tempFile)
			}

			var out bytes.Buffer
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestDecryptCommand tests the decrypt command functionality.
func TestDecryptCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "decrypt with text and preset",
			args:    []string{"decrypt", "--text", "JHLGQ", "--preset", "classic"},
			wantErr: false,
		},
		{
			name:    "decrypt with invalid preset",
			args:    []string{"decrypt", "--text", "HELLO", "--preset", "invalid"},
			wantErr: true,
		},
		{
			name:    "decrypt without input",
			args:    []string{"decrypt", "--preset", "classic"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestKeygenCommand tests the keygen command functionality.
func TestKeygenCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "keygen with basic settings",
			args:    []string{"keygen", "--security", "low", "--alphabet", "latin"},
			wantErr: false,
		},
		{
			name:    "keygen with preset",
			args:    []string{"keygen", "--preset", "classic"},
			wantErr: false,
		},
		{
			name:    "keygen with description",
			args:    []string{"keygen", "--preset", "classic", "--describe"},
			wantErr: false,
		},
		{
			name:    "keygen with stats",
			args:    []string{"keygen", "--preset", "classic", "--stats"},
			wantErr: false,
		},
		{
			name:    "keygen with invalid alphabet",
			args:    []string{"keygen", "--alphabet", "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// TestPresetCommand tests the preset command functionality.
func TestPresetCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "list presets",
			args:     []string{"preset", "--list"},
			wantErr:  false,
			contains: "classic",
		},
		{
			name:     "describe classic preset",
			args:     []string{"preset", "--describe", "classic"},
			wantErr:  false,
			contains: "Historical M3 Enigma",
		},
		{
			name:     "describe all presets",
			args:     []string{"preset", "--describe", "all"},
			wantErr:  false,
			contains: "classic",
		},
		{
			name:    "describe invalid preset",
			args:    []string{"preset", "--describe", "invalid"},
			wantErr: true,
		},
		{
			name:     "export preset",
			args:     []string{"preset", "--export", "classic"},
			wantErr:  false,
			contains: "alphabet",
		},
		{
			name:    "export invalid preset",
			args:    []string{"preset", "--export", "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := out.String()
			if tt.contains != "" && !strings.Contains(output, tt.contains) {
				t.Errorf("Output should contain '%s', got: %s", tt.contains, output)
			}
		})
	}
}

// TestConfigCommand tests the config command functionality.
func TestConfigCommand(t *testing.T) {
	// Create a test configuration file
	testConfig := `{
		"alphabet": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"rotor_specs": [{
			"id": "TestRotor",
			"forward_mapping": "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			"notches": [81],
			"position": 0,
			"ring_setting": 0
		}],
		"reflector_spec": {
			"id": "TestReflector", 
			"mapping": "YRUHQSLDPXNGOKMIEBFZCWVJAT"
		},
		"plugboard_pairs": {},
		"current_rotor_positions": [0]
	}`

	tmpFile, err := os.CreateTemp("", "test-config-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(testConfig)
	if err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}
	tmpFile.Close()

	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		contains string
	}{
		{
			name:     "validate config",
			args:     []string{"config", "--validate", tmpFile.Name()},
			wantErr:  false,
			contains: "VALID",
		},
		{
			name:     "show config",
			args:     []string{"config", "--show", tmpFile.Name()},
			wantErr:  false,
			contains: "Configuration File",
		},
		{
			name:     "show config detailed",
			args:     []string{"config", "--show", tmpFile.Name(), "--detailed"},
			wantErr:  false,
			contains: "Detailed Settings",
		},
		{
			name:     "test config",
			args:     []string{"config", "--test", tmpFile.Name()},
			wantErr:  false,
			contains: "Round-trip",
		},
		{
			name:    "validate nonexistent config",
			args:    []string{"config", "--validate", "nonexistent.json"},
			wantErr: true,
		},
		{
			name:    "convert without output",
			args:    []string{"config", "--convert", tmpFile.Name()},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			cmd := createTestRootCmd()
			cmd.SetOut(&out)
			cmd.SetErr(&out)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			output := out.String()
			if tt.contains != "" && !strings.Contains(output, tt.contains) {
				t.Errorf("Output should contain '%s', got: %s", tt.contains, output)
			}
		})
	}
}

// TestEncryptDecryptRoundTrip tests the full encryption/decryption workflow.
func TestEncryptDecryptRoundTrip(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "enigma-test-")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	originalText := "HELLOWORLDTESTMESSAGE"

	// Step 1: Generate a key
	keyFile := filepath.Join(tempDir, "test-key.json")
	var out bytes.Buffer
	cmd := createTestRootCmd()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"keygen", "--preset", "classic", "--output", keyFile})

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Verify key file was created
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		t.Fatalf("Key file was not created")
	}

	// Step 2: Encrypt with the generated key
	encryptedFile := filepath.Join(tempDir, "encrypted.txt")
	out.Reset()
	cmd = createTestRootCmd()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"encrypt", "--text", originalText, "--config", keyFile, "--output", encryptedFile})

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// Read encrypted content
	encryptedData, err := os.ReadFile(encryptedFile)
	if err != nil {
		t.Fatalf("Failed to read encrypted file: %v", err)
	}
	encryptedText := string(encryptedData)

	// Verify the text was actually encrypted (changed)
	if encryptedText == originalText {
		t.Errorf("Text was not encrypted (remained the same)")
	}

	// Step 3: Decrypt with the same key
	decryptedFile := filepath.Join(tempDir, "decrypted.txt")
	out.Reset()
	cmd = createTestRootCmd()
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"decrypt", "--text", encryptedText, "--config", keyFile, "--output", decryptedFile})

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	// Read decrypted content
	decryptedData, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatalf("Failed to read decrypted file: %v", err)
	}
	decryptedText := string(decryptedData)

	// Verify round-trip worked
	if decryptedText != originalText {
		t.Errorf("Round-trip failed. Original: %q, Decrypted: %q", originalText, decryptedText)
	}
}

// createTestRootCmd creates a fresh root command for testing.
func createTestRootCmd() *cobra.Command {
	// Create a new root command to avoid state pollution between tests
	testRootCmd := &cobra.Command{
		Use:     "eniGOma",
		Short:   "A highly customizable, Unicode-capable Enigma machine implementation",
		Version: "0.2.1",
	}

	// Create fresh command instances to avoid state pollution
	freshEncryptCmd := createFreshEncryptCmd()
	freshDecryptCmd := createFreshDecryptCmd()
	freshKeygenCmd := createFreshKeygenCmd()
	freshPresetCmd := createFreshPresetCmd()
	freshConfigCmd := createFreshConfigCmd()

	// Add subcommands
	testRootCmd.AddCommand(freshEncryptCmd)
	testRootCmd.AddCommand(freshDecryptCmd)
	testRootCmd.AddCommand(freshKeygenCmd)
	testRootCmd.AddCommand(freshPresetCmd)
	testRootCmd.AddCommand(freshConfigCmd)

	// Global flags
	testRootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	testRootCmd.PersistentFlags().StringP("config", "c", "", "Configuration file path")

	return testRootCmd
}

// Helper functions to create fresh command instances
func createFreshEncryptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt text or files using an Enigma machine",
		RunE:  runEncrypt,
	}

	// Input options
	cmd.Flags().StringP("text", "t", "", "Text to encrypt")
	cmd.Flags().StringP("file", "f", "", "File to encrypt")
	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")

	// Machine configuration
	cmd.Flags().StringP("preset", "p", "", "Use a preset configuration (classic, simple, high, extreme)")
	cmd.Flags().StringP("alphabet", "a", "latin", "Alphabet to use (latin, greek, cyrillic, portuguese, ascii, alphanumeric)")
	cmd.Flags().StringP("security", "s", "medium", "Security level (low, medium, high, extreme)")

	// Advanced options
	cmd.Flags().StringSliceP("rotors", "r", nil, "Rotor positions (e.g., 1,5,12)")
	cmd.Flags().StringSliceP("plugboard", "", nil, "Plugboard pairs (e.g., A:Z,B:Y)")
	cmd.Flags().BoolP("reset", "", false, "Reset machine to initial state before encryption")

	// Output formatting
	cmd.Flags().StringP("format", "", "text", "Output format (text, hex, base64)")
	cmd.Flags().BoolP("preserve-case", "", false, "Preserve original case (when possible)")

	return cmd
}

func createFreshDecryptCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt text or files using an Enigma machine",
		RunE:  runDecrypt,
	}

	// Input options
	cmd.Flags().StringP("text", "t", "", "Text to decrypt")
	cmd.Flags().StringP("file", "f", "", "File to decrypt")
	cmd.Flags().StringP("output", "o", "", "Output file (default: stdout)")

	// Machine configuration
	cmd.Flags().StringP("preset", "p", "", "Use a preset configuration (classic, simple, high, extreme)")
	cmd.Flags().StringP("alphabet", "a", "latin", "Alphabet to use (latin, greek, cyrillic, portuguese, ascii, alphanumeric)")
	cmd.Flags().StringP("security", "s", "medium", "Security level (low, medium, high, extreme)")

	// Advanced options
	cmd.Flags().StringSliceP("rotors", "r", nil, "Rotor positions (e.g., 1,5,12)")
	cmd.Flags().StringSliceP("plugboard", "", nil, "Plugboard pairs (e.g., A:Z,B:Y)")
	cmd.Flags().BoolP("reset", "", false, "Reset machine to initial state before decryption")

	// Input format
	cmd.Flags().StringP("format", "", "text", "Input format (text, hex, base64)")

	return cmd
}

func createFreshKeygenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keygen",
		Short: "Generate random Enigma machine configurations",
		RunE:  runKeygen,
	}

	// Machine configuration
	cmd.Flags().StringP("preset", "p", "", "Base preset to modify (classic, simple, low, medium, high, extreme)")
	cmd.Flags().StringP("alphabet", "a", "latin", "Alphabet to use (latin, greek, cyrillic, portuguese, ascii, alphanumeric)")
	cmd.Flags().StringP("security", "s", "medium", "Security level (low, medium, high, extreme)")

	// Output options
	cmd.Flags().StringP("output", "o", "", "Output file for the configuration (default: stdout)")
	cmd.Flags().StringP("save-to", "", "", "Save configuration to file (alias for --output)")
	cmd.Flags().StringP("format", "f", "json", "Output format (json, yaml)")

	// Advanced options
	cmd.Flags().IntP("rotors", "r", 0, "Number of rotors (overrides security level)")
	cmd.Flags().IntP("plugboard-pairs", "", 0, "Number of plugboard pairs (overrides security level)")
	cmd.Flags().BoolP("random-positions", "", true, "Generate random rotor positions")

	// Information options
	cmd.Flags().BoolP("describe", "d", false, "Show description of generated configuration")
	cmd.Flags().BoolP("stats", "", false, "Show statistics about the configuration")

	return cmd
}

func createFreshPresetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preset",
		Short: "List and describe available Enigma machine presets",
		RunE:  runPreset,
	}

	cmd.Flags().BoolP("list", "l", false, "List all available presets")
	cmd.Flags().StringP("describe", "d", "", "Describe a specific preset (or 'all' for all presets)")
	cmd.Flags().StringP("export", "e", "", "Export preset configuration to file")
	cmd.Flags().StringP("output", "o", "", "Output file for exported configuration")
	cmd.Flags().BoolP("verbose", "v", false, "Show detailed information")

	return cmd
}

func createFreshConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage Enigma machine configuration files",
		RunE:  runConfig,
	}

	cmd.Flags().StringP("validate", "", "", "Validate a configuration file")
	cmd.Flags().StringP("show", "s", "", "Show configuration details")
	cmd.Flags().StringP("test", "t", "", "Test configuration with sample text")
	cmd.Flags().StringP("text", "", "HELLOWORLD", "Text to use for testing")
	cmd.Flags().StringP("convert", "", "", "Convert/update configuration format")
	cmd.Flags().StringP("output", "o", "", "Output file for converted configuration")
	cmd.Flags().BoolP("detailed", "d", false, "Show detailed information")

	return cmd
}
