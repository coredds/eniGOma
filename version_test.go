package eniGOma

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion() returned empty string")
	}

	if version != Version {
		t.Errorf("GetVersion() = %s, want %s", version, Version)
	}

	// Check version format (should be semantic versioning)
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		t.Errorf("Version format invalid: %s (should be X.Y.Z)", version)
	}

	// Version should match the constant (already checked) and follow semver
}

func TestVersionConstant(t *testing.T) {
	if Version == "" {
		t.Error("Version constant is empty")
	}

	// Version should start with "0.4"
	if !strings.HasPrefix(Version, "0.4") {
		t.Errorf("Version should start with '0.4', got %s", Version)
	}
}
