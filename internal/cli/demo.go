// Package cli provides the demo command for the enigoma CLI.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package cli

import (
	"fmt"
	"time"

	"github.com/coredds/enigoma"
	"github.com/coredds/enigoma/pkg/enigma"
	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Interactive demonstration of enigoma features",
	Long: `Interactive demonstration showing enigoma's key features and capabilities.

This command runs a series of demonstrations to help you understand:
• Basic encryption and decryption
• Unicode and multi-language support  
• Different security levels
• Auto-detection capabilities
• Configuration management

Perfect for new users to see enigoma in action!

Example:
  enigoma demo`,
	RunE: runDemo,
}

func runDemo(cmd *cobra.Command, args []string) error {
	fmt.Printf("🎯 Welcome to the enigoma Interactive Demo!\n")
	fmt.Printf("Version: %s\n\n", enigoma.GetVersion())

	// Demo 1: Basic Encryption
	fmt.Println("📝 Demo 1: Basic Encryption & Decryption")
	fmt.Println("=========================================")

	message := "HELLOWORLD"
	fmt.Printf("Original message: %q\n", message)

	machine, err := enigma.NewEnigmaClassic()
	if err != nil {
		return fmt.Errorf("failed to create machine: %v", err)
	}

	encrypted, err := machine.Encrypt(message)
	if err != nil {
		return fmt.Errorf("encryption failed: %v", err)
	}
	fmt.Printf("Encrypted: %q\n", encrypted)

	if err := machine.Reset(); err != nil {
		return fmt.Errorf("failed to reset machine: %v", err)
	}
	decrypted, err := machine.Decrypt(encrypted)
	if err != nil {
		return fmt.Errorf("decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %q\n", decrypted)
	fmt.Printf("✅ Round-trip successful: %t\n\n", message == decrypted)

	time.Sleep(1 * time.Second)

	// Demo 2: Unicode Support
	fmt.Println("🌍 Demo 2: Unicode & Multi-Language Support")
	fmt.Println("===========================================")

	unicodeMessage := "Olá! Привет! 日本語! 🌟"
	fmt.Printf("Unicode message: %q\n", unicodeMessage)

	// Use auto-detection for Unicode
	unicodeMachine, err := enigma.NewFromText(unicodeMessage, enigma.Medium)
	if err != nil {
		return fmt.Errorf("failed to create Unicode machine: %v", err)
	}

	encryptedUnicode, err := unicodeMachine.Encrypt(unicodeMessage)
	if err != nil {
		return fmt.Errorf("Unicode encryption failed: %v", err)
	}
	fmt.Printf("Encrypted: %q\n", encryptedUnicode)

	if err := unicodeMachine.Reset(); err != nil {
		return fmt.Errorf("failed to reset Unicode machine: %v", err)
	}
	decryptedUnicode, err := unicodeMachine.Decrypt(encryptedUnicode)
	if err != nil {
		return fmt.Errorf("Unicode decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %q\n", decryptedUnicode)
	fmt.Printf("✅ Unicode round-trip successful: %t\n", unicodeMessage == decryptedUnicode)
	fmt.Printf("📊 Auto-detected alphabet size: %d characters\n\n", unicodeMachine.GetAlphabetSize())

	time.Sleep(1 * time.Second)

	// Demo 3: Security Levels
	fmt.Println("🛡️  Demo 3: Security Levels")
	fmt.Println("===========================")

	testMessage := "SECRETMESSAGE"
	levels := []enigma.SecurityLevel{enigma.Low, enigma.Medium, enigma.High, enigma.Extreme}
	levelNames := []string{"Low", "Medium", "High", "Extreme"}

	for i, level := range levels {
		fmt.Printf("%s Security:\n", levelNames[i])

		secMachine, err := enigma.New(
			enigma.WithAlphabet(enigoma.AlphabetLatinUpper),
			enigma.WithRandomSettings(level),
		)
		if err != nil {
			return fmt.Errorf("failed to create %s security machine: %v", levelNames[i], err)
		}

		fmt.Printf("  • Rotors: %d\n", secMachine.GetRotorCount())
		fmt.Printf("  • Plugboard pairs: %d\n", secMachine.GetPlugboardPairCount())

		secEncrypted, _ := secMachine.Encrypt(testMessage)
		fmt.Printf("  • Encrypted: %q\n", secEncrypted)

		if err := secMachine.Reset(); err != nil {
			return fmt.Errorf("failed to reset %s security machine: %v", levelNames[i], err)
		}
		secDecrypted, _ := secMachine.Decrypt(secEncrypted)
		fmt.Printf("  • ✅ Round-trip: %t\n\n", testMessage == secDecrypted)
	}

	time.Sleep(1 * time.Second)

	// Demo 4: Convenience Functions
	fmt.Println("⚡ Demo 4: Zero-Config Convenience Functions")
	fmt.Println("==========================================")

	quickMessage := "Quick encryption test"
	fmt.Printf("Message: %q\n", quickMessage)

	// Use the new convenience function
	quickEncrypted, quickConfig, err := enigma.EncryptText(quickMessage)
	if err != nil {
		return fmt.Errorf("quick encryption failed: %v", err)
	}
	fmt.Printf("Encrypted: %q\n", quickEncrypted)
	fmt.Printf("Config size: %d bytes\n", len(quickConfig))

	// Decrypt using the config
	quickDecrypted, err := enigma.DecryptWithConfig(quickEncrypted, quickConfig)
	if err != nil {
		return fmt.Errorf("quick decryption failed: %v", err)
	}
	fmt.Printf("Decrypted: %q\n", quickDecrypted)
	fmt.Printf("✅ Zero-config round-trip: %t\n\n", quickMessage == quickDecrypted)

	// Summary
	fmt.Println("🎉 Demo Complete!")
	fmt.Println("================")
	fmt.Println("You've seen:")
	fmt.Println("• ✅ Basic encryption/decryption")
	fmt.Println("• ✅ Unicode and multi-language support")
	fmt.Println("• ✅ Different security levels")
	fmt.Println("• ✅ Zero-config convenience functions")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("• Try: enigoma wizard (interactive setup)")
	fmt.Println("• Try: enigoma examples (copy-paste ready examples)")
	fmt.Println("• Try: enigoma encrypt --text \"Your text\" --auto-config key.json")
	fmt.Println()
	fmt.Println("Happy encrypting! 🔐")

	return nil
}
