// Package cli provides the command-line interface for eniGOma.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"fmt"

	"github.com/coredds/eniGOma"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eniGOma",
	Short: "A highly customizable, Unicode-capable Enigma machine implementation",
	Long: `eniGOma is a Go library and CLI tool that simulates the famous Enigma machine 
used during World War II, with modern enhancements including Unicode support,
configurable complexity, and modular design.

Examples:
  eniGOma encrypt --text "Hello World" --preset classic
  eniGOma decrypt --file encrypted.txt --config my-enigma.json
  eniGOma keygen --security high --alphabet latin --output my-key.json
  eniGOma preset --list`,
	Version: eniGOma.GetVersion(),
}

// Execute runs the root command and handles errors.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	rootCmd.AddCommand(keygenCmd)
	rootCmd.AddCommand(presetCmd)
	rootCmd.AddCommand(configCmd)

	// Global flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringP("config", "c", "", "Configuration file path")
}

// setupVerbose configures verbose logging if enabled.
func setupVerbose(cmd *cobra.Command) {
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		fmt.Println("Verbose mode enabled")
	}
}
