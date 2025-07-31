package alphabet

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		runes     []rune
		wantError bool
	}{
		{
			name:      "valid alphabet",
			runes:     []rune{'A', 'B', 'C'},
			wantError: false,
		},
		{
			name:      "empty alphabet",
			runes:     []rune{},
			wantError: true,
		},
		{
			name:      "duplicate characters",
			runes:     []rune{'A', 'B', 'A'},
			wantError: true,
		},
		{
			name:      "unicode characters",
			runes:     []rune{'Ω', 'α', 'β'},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alphabet, err := New(tt.runes)
			if tt.wantError {
				if err == nil {
					t.Errorf("New() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("New() unexpected error: %v", err)
				return
			}
			if alphabet.Size() != len(tt.runes) {
				t.Errorf("Size() = %d, want %d", alphabet.Size(), len(tt.runes))
			}
		})
	}
}

func TestAlphabet_RuneToIndex(t *testing.T) {
	alphabet, err := New([]rune{'C', 'A', 'B'}) // Preserves original order: C, A, B
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name      string
		rune      rune
		wantIndex int
		wantError bool
	}{
		{"first character", 'C', 0, false},
		{"middle character", 'A', 1, false},
		{"last character", 'B', 2, false},
		{"not in alphabet", 'D', 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, err := alphabet.RuneToIndex(tt.rune)
			if tt.wantError {
				if err == nil {
					t.Errorf("RuneToIndex() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("RuneToIndex() unexpected error: %v", err)
				return
			}
			if index != tt.wantIndex {
				t.Errorf("RuneToIndex() = %d, want %d", index, tt.wantIndex)
			}
		})
	}
}

func TestAlphabet_IndexToRune(t *testing.T) {
	alphabet, err := New([]rune{'C', 'A', 'B'}) // Preserves original order: C, A, B
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name      string
		index     int
		wantRune  rune
		wantError bool
	}{
		{"first index", 0, 'C', false},
		{"middle index", 1, 'A', false},
		{"last index", 2, 'B', false},
		{"negative index", -1, 0, true},
		{"index too large", 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rune, err := alphabet.IndexToRune(tt.index)
			if tt.wantError {
				if err == nil {
					t.Errorf("IndexToRune() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("IndexToRune() unexpected error: %v", err)
				return
			}
			if rune != tt.wantRune {
				t.Errorf("IndexToRune() = %c, want %c", rune, tt.wantRune)
			}
		})
	}
}

func TestAlphabet_Contains(t *testing.T) {
	alphabet, err := New([]rune{'A', 'B', 'C'})
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name string
		rune rune
		want bool
	}{
		{"contains A", 'A', true},
		{"contains B", 'B', true},
		{"contains C", 'C', true},
		{"does not contain D", 'D', false},
		{"does not contain lowercase", 'a', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabet.Contains(tt.rune); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphabet_ValidateString(t *testing.T) {
	alphabet, err := New([]rune{'A', 'B', 'C'})
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name      string
		input     string
		wantError bool
		errorRune rune
	}{
		{"valid string", "ABC", false, 0},
		{"valid repeated", "AAB", false, 0},
		{"invalid character", "ABD", true, 'D'},
		{"empty string", "", false, 0},
		{"first char invalid", "XBC", true, 'X'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invalidRune, err := alphabet.ValidateString(tt.input)
			if tt.wantError {
				if err == nil {
					t.Errorf("ValidateString() expected error but got none")
				}
				if invalidRune != tt.errorRune {
					t.Errorf("ValidateString() returned rune %c, want %c", invalidRune, tt.errorRune)
				}
				return
			}
			if err != nil {
				t.Errorf("ValidateString() unexpected error: %v", err)
			}
		})
	}
}

func TestAlphabet_StringToIndices(t *testing.T) {
	alphabet, err := New([]rune{'C', 'A', 'B'}) // Preserves original order: C, A, B
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		want    []int
		wantErr bool
	}{
		{"simple conversion", "CAB", []int{0, 1, 2}, false},
		{"repeated characters", "CCA", []int{0, 0, 1}, false},
		{"empty string", "", []int{}, false},
		{"invalid character", "ABD", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := alphabet.StringToIndices(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("StringToIndices() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("StringToIndices() unexpected error: %v", err)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("StringToIndices() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("StringToIndices()[%d] = %d, want %d", i, v, tt.want[i])
				}
			}
		})
	}
}

func TestAlphabet_IndicesToString(t *testing.T) {
	alphabet, err := New([]rune{'C', 'A', 'B'}) // Preserves original order: C, A, B
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	tests := []struct {
		name    string
		input   []int
		want    string
		wantErr bool
	}{
		{"simple conversion", []int{0, 1, 2}, "CAB", false},
		{"repeated indices", []int{0, 0, 1}, "CCA", false},
		{"empty slice", []int{}, "", false},
		{"invalid index", []int{0, 1, 5}, "", true},
		{"negative index", []int{-1}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := alphabet.IndicesToString(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("IndicesToString() expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("IndicesToString() unexpected error: %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("IndicesToString() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestAlphabet_Roundtrip(t *testing.T) {
	alphabet, err := New([]rune{'A', 'B', 'C', 'D', 'E'})
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	testString := "ABCDE"
	indices, err := alphabet.StringToIndices(testString)
	if err != nil {
		t.Fatalf("StringToIndices() error: %v", err)
	}

	result, err := alphabet.IndicesToString(indices)
	if err != nil {
		t.Fatalf("IndicesToString() error: %v", err)
	}

	if result != testString {
		t.Errorf("Roundtrip failed: %s -> %v -> %s", testString, indices, result)
	}
}

func TestAlphabet_Runes(t *testing.T) {
	originalRunes := []rune{'C', 'A', 'B'}
	alphabet, err := New(originalRunes)
	if err != nil {
		t.Fatalf("Failed to create alphabet: %v", err)
	}

	runes := alphabet.Runes()

	// Should preserve original order
	expected := []rune{'C', 'A', 'B'}
	if len(runes) != len(expected) {
		t.Errorf("Runes() length = %d, want %d", len(runes), len(expected))
	}

	for i, r := range runes {
		if r != expected[i] {
			t.Errorf("Runes()[%d] = %c, want %c", i, r, expected[i])
		}
	}

	// Verify it's a copy (modifying shouldn't affect original)
	runes[0] = 'X'
	newRunes := alphabet.Runes()
	if newRunes[0] == 'X' {
		t.Errorf("Runes() should return a copy, but modification affected original")
	}
}
