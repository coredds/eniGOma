// Package alphabet provides utilities for managing character sets (alphabets)
// used by the Enigma machine components.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package alphabet

import (
	"fmt"
	"sort"
)

// Alphabet represents a character set used by the Enigma machine.
// It provides bidirectional mapping between runes and their indices.
type Alphabet struct {
	runes    []rune
	runeToID map[rune]int
	size     int
}

// New creates a new Alphabet from the provided runes.
// It validates that there are no duplicate characters.
func New(runes []rune) (*Alphabet, error) {
	if len(runes) == 0 {
		return nil, fmt.Errorf("alphabet cannot be empty")
	}

	// Check for duplicates
	seen := make(map[rune]bool)
	for _, r := range runes {
		if seen[r] {
			return nil, fmt.Errorf("duplicate character found: %c", r)
		}
		seen[r] = true
	}

	// Create a copy and sort for consistent ordering
	runesCopy := make([]rune, len(runes))
	copy(runesCopy, runes)
	sort.Slice(runesCopy, func(i, j int) bool {
		return runesCopy[i] < runesCopy[j]
	})

	// Build the mapping
	runeToID := make(map[rune]int, len(runesCopy))
	for i, r := range runesCopy {
		runeToID[r] = i
	}

	return &Alphabet{
		runes:    runesCopy,
		runeToID: runeToID,
		size:     len(runesCopy),
	}, nil
}

// Size returns the number of characters in the alphabet.
func (a *Alphabet) Size() int {
	return a.size
}

// Runes returns a copy of the runes in the alphabet.
func (a *Alphabet) Runes() []rune {
	result := make([]rune, len(a.runes))
	copy(result, a.runes)
	return result
}

// RuneToIndex converts a rune to its index in the alphabet.
// Returns an error if the rune is not in the alphabet.
func (a *Alphabet) RuneToIndex(r rune) (int, error) {
	idx, exists := a.runeToID[r]
	if !exists {
		return 0, fmt.Errorf("character %c not found in alphabet", r)
	}
	return idx, nil
}

// IndexToRune converts an index to its corresponding rune.
// Returns an error if the index is out of bounds.
func (a *Alphabet) IndexToRune(idx int) (rune, error) {
	if idx < 0 || idx >= a.size {
		return 0, fmt.Errorf("index %d out of bounds [0, %d)", idx, a.size)
	}
	return a.runes[idx], nil
}

// Contains checks if a rune is present in the alphabet.
func (a *Alphabet) Contains(r rune) bool {
	_, exists := a.runeToID[r]
	return exists
}

// ValidateString checks if all runes in the string are present in the alphabet.
// Returns the first invalid rune found, or 0 if all are valid.
func (a *Alphabet) ValidateString(s string) (rune, error) {
	for _, r := range s {
		if !a.Contains(r) {
			return r, fmt.Errorf("character %c not found in alphabet", r)
		}
	}
	return 0, nil
}

// StringToIndices converts a string to a slice of indices.
func (a *Alphabet) StringToIndices(s string) ([]int, error) {
	result := make([]int, 0, len(s))
	for _, r := range s {
		idx, err := a.RuneToIndex(r)
		if err != nil {
			return nil, err
		}
		result = append(result, idx)
	}
	return result, nil
}

// IndicesToString converts a slice of indices to a string.
func (a *Alphabet) IndicesToString(indices []int) (string, error) {
	runes := make([]rune, 0, len(indices))
	for _, idx := range indices {
		r, err := a.IndexToRune(idx)
		if err != nil {
			return "", err
		}
		runes = append(runes, r)
	}
	return string(runes), nil
}
