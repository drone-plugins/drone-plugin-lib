package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	PluginErrorMessageKey = "PLUGIN_ERROR_MESSAGE"
)

// SetSecret sets a new secret by adding it to the output
func SetSecret(name, value string) error {
	return WriteEnvToOutputFile(name, value)
}

// UpdateSecret updates an existing secret with a new value
func UpdateSecret(name, value string) error {
	return WriteEnvToOutputFile(name, value)
}

// DeleteSecret removes a secret from the output
func DeleteSecret(name string) error {
	return WriteEnvToFile(name, "")
}

// SetError sets the error message and writes it to the DRONE_OUTPUT file
func SetError(message string) error {
	return WriteEnvToFile(PluginErrorMessageKey, message)
}

// WriteEnvToFile writes a key-value pair to the DRONE_OUTPUT file
func WriteEnvToOutputFile(key, value string) error {
	outputFilePath := os.Getenv("DRONE_OUTPUT")
	
	// Check the extension of the output file (.env or .out)
	ext := strings.ToLower(filepath.Ext(outputFilePath))

	if ext == ".env" {
		// Write in .env format (KEY=VALUE)
		return writeToFile(outputFilePath, fmt.Sprintf("%s=%s\n", key, value))
	} else if ext == ".out" {
		// Write in .out format (export KEY="VALUE")
		return writeToFile(outputFilePath, fmt.Sprintf("%s %s\n", key, value))
	} else {
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

// Helper function to append content to the file
func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}