package harness

import (
	"bufio"
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

	// DroneOutputFile is the key for the file where outputs can be exported and utilized in the subsequent steps in Harness CI pipeline.
	DroneOutputFile = "DRONE_OUTPUT"

	// HarnessOutputSecretFile is the key for the file where secrets can be exported and utilized in the subsequent steps in Harness CI pipeline.
	HarnessOutputSecretFile = "HARNESS_OUTPUT_SECRET_FILE"
)

// SetSecret sets a new secret by adding it to the HARNESS_OUTPUT_SECRET_FILE file
func SetSecret(name, value string) error {
	return WriteEnvToOutputFile(HarnessOutputSecretFile, name, value)
}

// UpdateSecret overwrites the value of an existing secret.
func UpdateSecret(name, value string) error {
	return UpdateOrRemoveKeyValue(HarnessOutputSecretFile, name, value, false)
}

// DeleteSecret removes a secret from the file entirely.
func DeleteSecret(name string) error {
	return UpdateOrRemoveKeyValue(HarnessOutputSecretFile, name, "", true)
}

// SetOutput sets a new secret by adding it to the DRONE_OUTPUT file
func SetOutput(name, value string) error {
	return WriteEnvToOutputFile(DroneOutputFile, name, value)
}

// UpdateOutput overwrites the value of an existing output.
func UpdateOutput(name, value string) error {
	return UpdateOrRemoveKeyValue(DroneOutputFile, name, value, false)
}

// DeleteOutput removes an output from the file entirely.
func DeleteOutput(name string) error {
	return UpdateOrRemoveKeyValue(DroneOutputFile, name, "", true)
}

// SetErrorMetadata sets the error message, error code, and error category, writing them to the CI_ERROR_METADATA file
func SetErrorMetadata(message, code, category string) error {
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
		// Write in .out format (KEY VALUE)
		content = fmt.Sprintf("%s %s\n", key, value)
	} else {
		return fmt.Errorf("unsupported file extension: %s", ext)
	}

	return WriteToFile(filePath, content)
}

// Helper function to append content to the file
func WriteToFile(filename, content string) error {
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


// UpdateOrRemoveKeyValue updates or deletes a key-value pair in the specified file.
func UpdateOrRemoveKeyValue(envVar, key, newValue string, delete bool) error {
	// Get the file path from the environment variable
	filePath := os.Getenv(envVar)
	if filePath == "" {
		return fmt.Errorf("environment variable %s is not set", envVar)
	}

	// Determine the file extension to handle formats
	ext := strings.ToLower(filepath.Ext(filePath))

	// Read the file contents into memory
	lines, err := ReadLines(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Process lines
	var updatedLines []string
	found := false
	for _, line := range lines {
		k, v := ParseKeyValue(line, ext)
		if k == key {
			found = true
			if delete {
				continue // Skip the line to delete it
			}
			updatedLines = append(updatedLines, FormatKeyValue(k, newValue, ext))
		} else {
			updatedLines = append(updatedLines, FormatKeyValue(k, v, ext))
		}
	}

	// Append new key-value if not found and not deleting
	if !found && !delete {
		updatedLines = append(updatedLines, FormatKeyValue(key, newValue, ext))
	}

	// Write updated lines back to the file
	return WriteLines(filePath, updatedLines)
}

// Helper function to read lines from a file.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Helper function to write lines to a file.
func WriteLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	return nil
}

// Helper function to parse a line into key and value, considering file format.
func ParseKeyValue(line, ext string) (string, string) {
	if ext == ".env" {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		}
		return strings.TrimSpace(parts[0]), ""
	} else if ext == ".out" {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		}
		return strings.TrimSpace(parts[0]), ""
	}
	return "", ""
}

// Helper function to format a key-value pair as a line, considering file format.
func FormatKeyValue(key, value, ext string) string {
	if ext == ".env" {
		return fmt.Sprintf("%s=%s", key, value)
	} else if ext == ".out" {
		return fmt.Sprintf("%s %s", key, value)
	}
	return ""
}