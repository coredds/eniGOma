package enigma

import (
	"testing"
)

// TestSettingsJSONRoundTrip ensures that saving then loading settings preserves core characteristics.
func TestSettingsJSONRoundTrip(t *testing.T) {
	// Simple Latin alphabet
	alphabet := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	machine, err := New(
		WithAlphabet(alphabet),
		WithRandomSettings(Low),
		WithRandomRotorPositionsSeed(42),
	)
	if err != nil {
		t.Fatalf("failed to create machine: %v", err)
	}

	jsonData, err := machine.SaveSettingsToJSON()
	if err != nil {
		t.Fatalf("failed to save settings: %v", err)
	}

	machine2, err := NewFromJSON(jsonData)
	if err != nil {
		t.Fatalf("failed to load settings: %v", err)
	}

	if machine2.GetAlphabetSize() != machine.GetAlphabetSize() {
		t.Fatalf("alphabet size mismatch: %d vs %d", machine2.GetAlphabetSize(), machine.GetAlphabetSize())
	}
	if machine2.GetRotorCount() != machine.GetRotorCount() {
		t.Fatalf("rotor count mismatch: %d vs %d", machine2.GetRotorCount(), machine.GetRotorCount())
	}
}
