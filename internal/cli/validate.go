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

// fixNotchesFormat converts integer notches to string notches in the configuration.
func fixNotchesFormat(configJSON map[string]interface{}) {
	rotorSpecs, ok := configJSON["rotor_specs"].([]interface{})
	if !ok {
		return
	}

	for _, rotorSpec := range rotorSpecs {
		rs, ok := rotorSpec.(map[string]interface{})
		if !ok {
			continue
		}

		notches, ok := rs["notches"].([]interface{})
		if !ok {
			continue
		}

		stringNotches := make([]interface{}, len(notches))
		for i, notch := range notches {
			if n, ok := notch.(float64); ok {
				// Convert integer notch to string (single character)
				stringNotches[i] = string(rune(int(n)))
			} else {
				stringNotches[i] = notch
			}
		}
		rs["notches"] = stringNotches
	}
}

// fixReflectorMapping converts string reflector mapping to object mapping in the configuration.
func fixReflectorMapping(configJSON map[string]interface{}) {
	reflectorSpec, ok := configJSON["reflector_spec"].(map[string]interface{})
	if !ok {
		return
	}

	mapping, ok := reflectorSpec["mapping"].(string)
	if !ok {
		return
	}

	alphabet, ok := configJSON["alphabet"].(string)
	if !ok {
		return
	}

	mappingObj := make(map[string]interface{})
	for i, r := range mapping {
		if i < len(alphabet) {
			mappingObj[string(alphabet[i])] = string(r)
		}
	}
	reflectorSpec["mapping"] = mappingObj
}

// findSchemaPath returns the path to the schema file.
func findSchemaPath() (string, error) {
	schemaPath := filepath.Join("schemas", "config.v1.schema.json")

	// Check if schema file exists in the current directory
	if _, err := os.Stat(schemaPath); err == nil {
		return schemaPath, nil
	}

	// Try to find schema in the executable directory
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %v", err)
	}

	execDir := filepath.Dir(execPath)
	schemaPath = filepath.Join(execDir, "schemas", "config.v1.schema.json")

	// Check if schema file exists in the executable directory
	if _, err := os.Stat(schemaPath); err != nil {
		return "", fmt.Errorf("schema file not found: %v", schemaPath)
	}

	return schemaPath, nil
}

// loadSchema loads and compiles the JSON schema from the given path.
func loadSchema(schemaPath string) (*jsonschema.Schema, error) {
	// Open schema file
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open schema file: %v", err)
	}
	defer schemaFile.Close()

	// Read schema content
	schemaData, err := io.ReadAll(schemaFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %v", err)
	}

	// Compile schema
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft7

	// Add schema to compiler
	if err := compiler.AddResource("config.v1.schema.json", bytes.NewReader(schemaData)); err != nil {
		return nil, fmt.Errorf("failed to add schema resource: %v", err)
	}

	// Compile schema
	schema, err := compiler.Compile("config.v1.schema.json")
	if err != nil {
		return nil, fmt.Errorf("failed to compile schema: %v", err)
	}

	return schema, nil
}

// ValidateConfigAgainstSchema validates a configuration file against the JSON schema.
func ValidateConfigAgainstSchema(configFile string) error {
	// Read configuration file
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Parse JSON to ensure it's valid
	var configJSON map[string]interface{}
	if err := json.Unmarshal(configData, &configJSON); err != nil {
		return fmt.Errorf("invalid JSON in config file: %v", err)
	}

	// Fix format issues for validation
	fixNotchesFormat(configJSON)
	fixReflectorMapping(configJSON)

	// Find and load schema
	schemaPath, err := findSchemaPath()
	if err != nil {
		return err
	}

	schema, err := loadSchema(schemaPath)
	if err != nil {
		return err
	}

	// Validate configuration against schema
	if err := schema.Validate(configJSON); err != nil {
		return fmt.Errorf("configuration validation failed: %v", err)
	}

	return nil
}
