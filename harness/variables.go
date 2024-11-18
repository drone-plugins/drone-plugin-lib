package harness

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// ErrorMessageKey is the key used to retrieve or store the error message content.
	ErrorMessageKey = "ERROR_MESSAGE"

	// ErrorCodeKey is the key used to identify the specific error code associated with an error.
	ErrorCodeKey = "ERROR_CODE"

	// ErrorCategoryKey is the key used to classify the category of the error, which can help in grouping similar types of errors.
	ErrorCategoryKey = "ERROR_CATEGORY"

	// MetadataFile is the key for the file that stores metadata associated with an error, such as details about the error's source or context.
	MetadataFile = "ERROR_METADATA_FILE"

	// DroneOutputFile is the key for the file where output related to the Drone CI process is stored.
	DroneOutputFile = "DRONE_OUTPUT"
)

// SetSecret sets a new secret by adding it to the DRONE_OUTPUT file
func SetSecret(name, value string) error {
	return WriteEnvToOutputFile(DroneOutputFile, name, value)
}

// UpdateSecret updates an existing secret with a new value in the DRONE_OUTPUT file
func UpdateSecret(name, value string) error {
	return WriteEnvToOutputFile(DroneOutputFile, name, value)
}

// DeleteSecret removes a secret by setting it to an empty value in the DRONE_OUTPUT file
func DeleteSecret(name string) error {
	return WriteEnvToOutputFile(DroneOutputFile, name, "")
}

// SetError sets the error message and error code, writing them to the CI_ERROR_METADATA file
// SetError sets the error message, error code, and error category, writing them to the CI_ERROR_METADATA file
func SetError(message, code, category string) error {
	// Write the error message
	if err := WriteEnvToOutputFile(MetadataFile, ErrorMessageKey, message); err != nil {
		return err
	}

	// Write the error code
	if err := WriteEnvToOutputFile(MetadataFile, ErrorCodeKey, code); err != nil {
		return err
	}

	// Write the error category
	if err := WriteEnvToOutputFile(MetadataFile, ErrorCategoryKey, category); err != nil {
		return err
	}

	return nil
}

// WriteEnvToOutputFile writes a key-value pair to the specified file, determined by an environment variable
func WriteEnvToOutputFile(envVar, key, value string) error {
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
