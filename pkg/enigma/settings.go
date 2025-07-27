// Package enigma provides settings management for the Enigma machine.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package enigma

import (
	"encoding/json"
	"fmt"

	"github.com/coredds/eniGOma/internal/alphabet"
	"github.com/coredds/eniGOma/internal/plugboard"
	"github.com/coredds/eniGOma/internal/reflector"
	"github.com/coredds/eniGOma/internal/rotor"
)

// EnigmaSettings represents the serializable configuration and state of an Enigma machine.
type EnigmaSettings struct {
	Alphabet              []rune                  `json:"alphabet"`
	RotorSpecs            []rotor.RotorSpec       `json:"rotor_specs"`
	ReflectorSpec         reflector.ReflectorSpec `json:"reflector_spec"`
	PlugboardPairs        map[rune]rune           `json:"plugboard_pairs"`
	CurrentRotorPositions []int                   `json:"current_rotor_positions"`
}

// GetSettings returns the current configuration and state of the Enigma machine.
func (e *Enigma) GetSettings() (*EnigmaSettings, error) {
	if e.alphabet == nil {
		return nil, fmt.Errorf("alphabet is not initialized")
	}

	// Get alphabet runes
	alphabetRunes := e.alphabet.Runes()

	// Get rotor specifications
	rotorSpecs := make([]rotor.RotorSpec, len(e.rotors))
	for i, r := range e.rotors {
		spec, err := rotor.ToSpec(r, e.alphabet)
		if err != nil {
			return nil, fmt.Errorf("failed to get spec for rotor %d: %v", i, err)
		}
		rotorSpecs[i] = spec
	}

	// Get reflector specification
	reflectorSpec, err := reflector.ToSpec(e.reflector, e.alphabet)
	if err != nil {
		return nil, fmt.Errorf("failed to get reflector spec: %v", err)
	}

	// Get plugboard pairs
	plugboardPairs, err := e.plugboard.GetPairsMap()
	if err != nil {
		return nil, fmt.Errorf("failed to get plugboard pairs: %v", err)
	}

	// Get current rotor positions
	currentPositions := e.GetCurrentRotorPositions()

	return &EnigmaSettings{
		Alphabet:              alphabetRunes,
		RotorSpecs:            rotorSpecs,
		ReflectorSpec:         reflectorSpec,
		PlugboardPairs:        plugboardPairs,
		CurrentRotorPositions: currentPositions,
	}, nil
}

// LoadSettings initializes the Enigma machine with the provided settings.
func (e *Enigma) LoadSettings(settings *EnigmaSettings) error {
	if settings == nil {
		return fmt.Errorf("settings cannot be nil")
	}

	// Create alphabet
	alph, err := alphabet.New(settings.Alphabet)
	if err != nil {
		return fmt.Errorf("failed to create alphabet: %v", err)
	}
	e.alphabet = alph

	// Create rotors
	rotors := make([]rotor.Rotor, len(settings.RotorSpecs))
	for i, spec := range settings.RotorSpecs {
		r, err := rotor.CreateFromSpec(spec, e.alphabet)
		if err != nil {
			return fmt.Errorf("failed to create rotor %d: %v", i, err)
		}
		rotors[i] = r
	}
	e.rotors = rotors

	// Create reflector
	refl, err := reflector.CreateFromSpec(settings.ReflectorSpec, e.alphabet)
	if err != nil {
		return fmt.Errorf("failed to create reflector: %v", err)
	}
	e.reflector = refl

	// Create plugboard
	pb, err := plugboard.New(e.alphabet)
	if err != nil {
		return fmt.Errorf("failed to create plugboard: %v", err)
	}

	if len(settings.PlugboardPairs) > 0 {
		err = pb.SetPairsFromMap(settings.PlugboardPairs)
		if err != nil {
			return fmt.Errorf("failed to set plugboard pairs: %v", err)
		}
	}
	e.plugboard = pb

	// Set current rotor positions if provided
	if len(settings.CurrentRotorPositions) > 0 {
		if len(settings.CurrentRotorPositions) != len(e.rotors) {
			return fmt.Errorf("current position count (%d) doesn't match rotor count (%d)",
				len(settings.CurrentRotorPositions), len(e.rotors))
		}

		for i, pos := range settings.CurrentRotorPositions {
			e.rotors[i].SetPosition(pos)
		}
	}

	// Store initial settings for reset functionality
	// Make a copy without current positions for reset
	initialSettings := *settings
	initialSettings.CurrentRotorPositions = make([]int, len(settings.RotorSpecs))
	for i, spec := range settings.RotorSpecs {
		initialSettings.CurrentRotorPositions[i] = spec.Position
	}
	e.initialSettings = initialSettings

	return nil
}

// MarshalJSON marshals the EnigmaSettings to JSON.
func (s *EnigmaSettings) MarshalJSON() ([]byte, error) {
	// Convert runes to strings for JSON compatibility
	type jsonSettings struct {
		Alphabet              string                  `json:"alphabet"`
		RotorSpecs            []rotor.RotorSpec       `json:"rotor_specs"`
		ReflectorSpec         reflector.ReflectorSpec `json:"reflector_spec"`
		PlugboardPairs        map[string]string       `json:"plugboard_pairs"`
		CurrentRotorPositions []int                   `json:"current_rotor_positions"`
	}

	js := jsonSettings{
		Alphabet:              string(s.Alphabet),
		RotorSpecs:            s.RotorSpecs,
		ReflectorSpec:         s.ReflectorSpec,
		CurrentRotorPositions: s.CurrentRotorPositions,
		PlugboardPairs:        make(map[string]string),
	}

	// Convert rune pairs to string pairs
	for k, v := range s.PlugboardPairs {
		js.PlugboardPairs[string(k)] = string(v)
	}

	return json.Marshal(js)
}

// UnmarshalJSON unmarshals JSON to EnigmaSettings.
func (s *EnigmaSettings) UnmarshalJSON(data []byte) error {
	type jsonSettings struct {
		Alphabet              string                  `json:"alphabet"`
		RotorSpecs            []rotor.RotorSpec       `json:"rotor_specs"`
		ReflectorSpec         reflector.ReflectorSpec `json:"reflector_spec"`
		PlugboardPairs        map[string]string       `json:"plugboard_pairs"`
		CurrentRotorPositions []int                   `json:"current_rotor_positions"`
	}

	var js jsonSettings
	if err := json.Unmarshal(data, &js); err != nil {
		return err
	}

	s.Alphabet = []rune(js.Alphabet)
	s.RotorSpecs = js.RotorSpecs
	s.ReflectorSpec = js.ReflectorSpec
	s.CurrentRotorPositions = js.CurrentRotorPositions
	s.PlugboardPairs = make(map[rune]rune)

	// Convert string pairs back to rune pairs
	for k, v := range js.PlugboardPairs {
		if len(k) != 1 || len(v) != 1 {
			return fmt.Errorf("invalid plugboard pair: %s->%s", k, v)
		}
		kRune := []rune(k)[0]
		vRune := []rune(v)[0]
		s.PlugboardPairs[kRune] = vRune
	}

	return nil
}

// SaveSettingsToJSON saves the current Enigma settings to a JSON string.
func (e *Enigma) SaveSettingsToJSON() (string, error) {
	settings, err := e.GetSettings()
	if err != nil {
		return "", fmt.Errorf("failed to get settings: %v", err)
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal settings: %v", err)
	}

	return string(data), nil
}

// LoadSettingsFromJSON loads Enigma settings from a JSON string.
func (e *Enigma) LoadSettingsFromJSON(jsonData string) error {
	var settings EnigmaSettings
	if err := json.Unmarshal([]byte(jsonData), &settings); err != nil {
		return fmt.Errorf("failed to unmarshal settings: %v", err)
	}

	return e.LoadSettings(&settings)
}

// NewFromSettings creates a new Enigma machine from the provided settings.
func NewFromSettings(settings *EnigmaSettings) (*Enigma, error) {
	e := &Enigma{}
	if err := e.LoadSettings(settings); err != nil {
		return nil, err
	}
	return e, nil
}

// NewFromJSON creates a new Enigma machine from JSON settings.
func NewFromJSON(jsonData string) (*Enigma, error) {
	var settings EnigmaSettings
	if err := json.Unmarshal([]byte(jsonData), &settings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings: %v", err)
	}

	return NewFromSettings(&settings)
}
