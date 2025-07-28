// Package cli provides the config command for the eniGOma CLI.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"fmt"
	"os"

	"github.com/coredds/eniGOma/pkg/enigma"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Enigma machine configuration files",
	Long: `Manage Enigma machine configuration files.

This command helps validate, inspect, and manipulate configuration files
used by the eniGOma CLI and library.

Examples:
  eniGOma config --validate my-config.json
  eniGOma config --show my-config.json
  eniGOma config --test my-config.json --text "Hello World"
  eniGOma config --convert old-config.json --output new-config.json`,
	RunE: runConfig,
}

func init() {
	configCmd.Flags().StringP("validate", "", "", "Validate a configuration file")
	configCmd.Flags().StringP("show", "s", "", "Show configuration details")
	configCmd.Flags().StringP("test", "t", "", "Test configuration with sample text")
	configCmd.Flags().StringP("text", "", "Hello World", "Text to use for testing")
	configCmd.Flags().StringP("convert", "", "", "Convert/update configuration format")
	configCmd.Flags().StringP("output", "o", "", "Output file for converted configuration")
	configCmd.Flags().BoolP("detailed", "d", false, "Show detailed information")
}

func runConfig(cmd *cobra.Command, args []string) error {
	setupVerbose(cmd)

	validate, _ := cmd.Flags().GetString("validate")
	show, _ := cmd.Flags().GetString("show")
	test, _ := cmd.Flags().GetString("test")
	convert, _ := cmd.Flags().GetString("convert")

	// Handle different operations
	if validate != "" {
		return validateConfig(validate, cmd)
	}

	if show != "" {
		return showConfig(show, cmd)
	}

	if test != "" {
		return testConfig(test, cmd)
	}

	if convert != "" {
		return convertConfig(convert, cmd)
	}

	// Default: show help if no operation specified
	return cmd.Help()
}

func validateConfig(configFile string, cmd *cobra.Command) error {
	fmt.Printf("Validating configuration file: %s\n", configFile)

	// Try to read and parse the configuration
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Try to create machine from configuration
	machine, err := enigma.NewFromJSON(string(data))
	if err != nil {
		fmt.Printf("❌ Configuration is INVALID: %v\n", err)
		return nil
	}

	// Additional validation
	fmt.Printf("✅ Configuration is VALID\n")
	fmt.Printf("   Alphabet Size: %d characters\n", machine.GetAlphabetSize())
	fmt.Printf("   Rotors: %d\n", machine.GetRotorCount())
	fmt.Printf("   Plugboard Pairs: %d\n", machine.GetPlugboardPairCount())
	fmt.Printf("   Current Rotor Positions: %v\n", machine.GetCurrentRotorPositions())

	return nil
}

func showConfig(configFile string, cmd *cobra.Command) error {
	detailed, _ := cmd.Flags().GetBool("detailed")

	// Read configuration
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Create machine from configuration
	machine, err := enigma.NewFromJSON(string(data))
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %v", err)
	}

	// Show basic information
	fmt.Printf("Configuration File: %s\n", configFile)
	fmt.Printf("==========================================\n")
	fmt.Printf("Alphabet Size: %d characters\n", machine.GetAlphabetSize())
	fmt.Printf("Rotors: %d\n", machine.GetRotorCount())
	fmt.Printf("Plugboard Pairs: %d\n", machine.GetPlugboardPairCount())
	fmt.Printf("Current Rotor Positions: %v\n", machine.GetCurrentRotorPositions())

	if detailed {
		fmt.Printf("\nDetailed Settings:\n")
		fmt.Printf("------------------\n")

		// Get full settings
		settings, err := machine.GetSettings()
		if err != nil {
			return fmt.Errorf("failed to get detailed settings: %v", err)
		}

		fmt.Printf("Alphabet: %s\n", string(settings.Alphabet))
		fmt.Printf("Rotor Count: %d\n", len(settings.RotorSpecs))

		for i, rotor := range settings.RotorSpecs {
			fmt.Printf("  Rotor %d: ID=%s, Position=%d, Ring=%d\n",
				i+1, rotor.ID, rotor.Position, rotor.RingSetting)
		}

		fmt.Printf("Reflector: ID=%s\n", settings.ReflectorSpec.ID)
		fmt.Printf("Plugboard Pairs: %d\n", len(settings.PlugboardPairs))

		if len(settings.PlugboardPairs) > 0 {
			fmt.Printf("  Pairs: ")
			for k, v := range settings.PlugboardPairs {
				fmt.Printf("%c↔%c ", k, v)
			}
			fmt.Printf("\n")
		}
	}

	return nil
}

func testConfig(configFile string, cmd *cobra.Command) error {
	testText, _ := cmd.Flags().GetString("text")

	fmt.Printf("Testing configuration: %s\n", configFile)
	fmt.Printf("Test text: %s\n", testText)
	fmt.Printf("========================\n")

	// Create machine from configuration
	machine, err := createMachineFromConfig(configFile)
	if err != nil {
		return fmt.Errorf("failed to create machine from config: %v", err)
	}

	// Test encryption
	encrypted, err := machine.Encrypt(testText)
	if err != nil {
		return fmt.Errorf("encryption test failed: %v", err)
	}

	fmt.Printf("Encrypted: %s\n", encrypted)

	// Reset machine and test decryption
	if err := machine.Reset(); err != nil {
		return fmt.Errorf("failed to reset machine: %v", err)
	}

	decrypted, err := machine.Decrypt(encrypted)
	if err != nil {
		return fmt.Errorf("decryption test failed: %v", err)
	}

	fmt.Printf("Decrypted: %s\n", decrypted)

	// Verify round-trip
	if testText == decrypted {
		fmt.Printf("✅ Round-trip test PASSED\n")
	} else {
		fmt.Printf("❌ Round-trip test FAILED\n")
		fmt.Printf("   Expected: %s\n", testText)
		fmt.Printf("   Got:      %s\n", decrypted)
	}

	return nil
}

func convertConfig(configFile string, cmd *cobra.Command) error {
	outputFile, _ := cmd.Flags().GetString("output")

	if outputFile == "" {
		return fmt.Errorf("output file required for conversion (use --output)")
	}

	fmt.Printf("Converting configuration: %s → %s\n", configFile, outputFile)

	// Read and validate input configuration
	machine, err := createMachineFromConfig(configFile)
	if err != nil {
		return fmt.Errorf("failed to read input configuration: %v", err)
	}

	// Export to new format (currently just re-export as JSON)
	jsonData, err := machine.SaveSettingsToJSON()
	if err != nil {
		return fmt.Errorf("failed to convert configuration: %v", err)
	}

	// Write to output file
	err = writeStringToFile(jsonData, outputFile)
	if err != nil {
		return fmt.Errorf("failed to write converted configuration: %v", err)
	}

	fmt.Printf("✅ Configuration converted successfully\n")

	return nil
}
