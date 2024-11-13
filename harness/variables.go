package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	CIErrorMessageKey  = "CI_ERROR_MESSAGE"
	CIErrorCodeKey     = "CI_ERROR_CODE"
	CIMetadataFileEnv  = "CI_ERROR_METADATA"
	DroneOutputFileEnv = "DRONE_OUTPUT"
)

// SetSecret sets a new secret by adding it to the DRONE_OUTPUT file
func SetSecret(name, value string) error {
	return WriteEnvToFile(DroneOutputFileEnv, name, value)
}

// UpdateSecret updates an existing secret with a new value in the DRONE_OUTPUT file
func UpdateSecret(name, value string) error {
	return WriteEnvToFile(DroneOutputFileEnv, name, value)
}

// DeleteSecret removes a secret by setting it to an empty value in the DRONE_OUTPUT file
func DeleteSecret(name string) error {
	return WriteEnvToFile(DroneOutputFileEnv, name, "")
}

// SetError sets the error message and error code, writing them to the CI_ERROR_METADATA file
func SetError(message, code string) error {
	if err := WriteEnvToFile(CIMetadataFileEnv, CIErrorMessageKey, message); err != nil {
		return err
	}
	return WriteEnvToFile(CIMetadataFileEnv, CIErrorCodeKey, code)
}

// WriteEnvToFile writes a key-value pair to the specified file, determined by an environment variable
func WriteEnvToFile(envVar, key, value string) error {
	// Get the file path from the specified environment variable
	filePath := os.Getenv(envVar)
	if filePath == "" {
		return fmt.Errorf("environment variable %s is not set", envVar)
	}

	// Check the extension of the file (.env or .out)
	ext := strings.ToLower(filepath.Ext(filePath))

	var content string
	if ext == ".env" {
		// Write in .env format (KEY=VALUE)
		content = fmt.Sprintf("%s=%s\n", key, value)
	} else if ext == ".out" {
		// Write in .out format (export KEY="VALUE")
		content = fmt.Sprintf("%s \"%s\"\n", key, value)
	} else {
		return fmt.Errorf("unsupported file extension: %s", ext)
	}

	return writeToFile(filePath, content)
}

// Helper function to append content to the file
func writeToFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
