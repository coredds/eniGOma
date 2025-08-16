// Package cli provides validation utilities for configuration files.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

// ValidateConfigAgainstSchema validates a configuration file against the JSON schema.
func ValidateConfigAgainstSchema(configFile string) error {
	// Read configuration file
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Parse JSON to ensure it's valid
	var configJSON interface{}
	if err := json.Unmarshal(configData, &configJSON); err != nil {
		return fmt.Errorf("invalid JSON in config file: %v", err)
	}

	// Find schema file
	schemaPath := filepath.Join("schemas", "config.v1.schema.json")

	// Check if schema file exists in the current directory
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		// If not found, try to find schema in the executable directory
		execPath, err := os.Executable()
		if err == nil {
			execDir := filepath.Dir(execPath)
			schemaPath = filepath.Join(execDir, "schemas", "config.v1.schema.json")
		}
	}

	// If schema file is still not found, return an error
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		return fmt.Errorf("schema file not found: %v", schemaPath)
	}

	// Load schema
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft7

	// Open schema file
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to open schema file: %v", err)
	}
	defer schemaFile.Close()

	// Read schema content
	schemaData, err := io.ReadAll(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Add schema to compiler
	if err := compiler.AddResource("config.v1.schema.json", bytes.NewReader(schemaData)); err != nil {
		return fmt.Errorf("failed to add schema resource: %v", err)
	}

	// Compile schema
	schema, err := compiler.Compile("config.v1.schema.json")
	if err != nil {
		return fmt.Errorf("failed to compile schema: %v", err)
	}

	// Validate configuration against schema
	if err := schema.Validate(configJSON); err != nil {
		return fmt.Errorf("configuration validation failed: %v", err)
	}

	return nil
}
