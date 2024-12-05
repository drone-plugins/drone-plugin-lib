package harness

import (
	"os"
	"testing"

	v3 "github.com/harness/godotenv/v3"
	"github.com/stretchr/testify/assert"
)

// Helper function to create the .env file with the given content
func createEnvFile(t *testing.T, filePath string, content map[string]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()

	for key, value := range content {
		_, err := file.WriteString(key + "=" + value + "\n")
		if err != nil {
			t.Fatalf("Failed to write to file: %v", err)
			return err
		}
	}

	return nil
}

// Helper function to create an output file (.out) with the given content
func createOutFile(t *testing.T, filePath string, content []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()

	// Write the content to the .out file manually
	for _, line := range content {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			t.Fatalf("Failed to write to file: %v", err)
			return err
		}
	}
	return nil
}

// Test Set, Update, and Delete functions for .env files
func TestSetUpdateDeleteEnvFile(t *testing.T) {
	// Set the environment variable to point to the test.env file
	envFilePath := "test.env"
	os.Setenv("HARNESS_OUTPUT_SECRET_FILE", envFilePath)

	// Setup: Create the .env file at the path specified by HARNESS_OUTPUT_SECRET_FILE
	envContent := map[string]string{
		"KEY1": "value1",
		"KEY2": "value2",
	}
	err := createEnvFile(t, envFilePath, envContent)
	assert.NoError(t, err)

	// Defer file removal after the test completes
	defer os.Remove(envFilePath)

	// Set a new key-value pair
	err = SetSecret("KEY3", "value3")
	assert.NoError(t, err)

	// Update an existing key-value pair
	err = UpdateSecret("KEY1", "new_value1")
	assert.NoError(t, err)

	// Delete an existing key
	err = DeleteSecret("KEY2")
	assert.NoError(t, err)

	// Read the file content and verify the changes
	data, err := v3.Read(envFilePath)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, "new_value1", data["KEY1"])
	assert.Equal(t, "value3", data["KEY3"])
	assert.NotContains(t, data, "KEY2")

	// Clean up
	defer os.Unsetenv("HARNESS_OUTPUT_SECRET_FILE")
}

// Test Set, Update, and Delete functions for .out files (single-line values)
func TestSetUpdateDeleteOutFile(t *testing.T) {
	// Set the environment variable to point to the test.out file
	outFilePath := "test.out"
	os.Setenv("DRONE_OUTPUT", outFilePath)

	// Setup: Create the .out file at the path specified by DRONE_OUTPUT
	outContent := []string{
		"KEY1 value1",
		"KEY2 value2",
	}
	err := createOutFile(t, outFilePath, outContent)
	assert.NoError(t, err)

	// Defer file removal after the test completes
	defer os.Remove(outFilePath)

	// Set a new key-value pair
	err = SetOutput("KEY3", "value3")
	assert.NoError(t, err)

	// Update an existing key-value pair
	err = UpdateOutput("KEY1", "new_value1")
	assert.NoError(t, err)

	// Delete an existing key
	err = DeleteOutput("KEY2")
	assert.NoError(t, err)

	// Verify changes
	lines, err := ReadLines(outFilePath)
	assert.NoError(t, err)

	// Assertions
	assert.Contains(t, lines, "KEY1 new_value1")
	assert.Contains(t, lines, "KEY3 value3")
	assert.NotContains(t, lines, "KEY2")

	// Clean up
	defer os.Unsetenv("DRONE_OUTPUT")
}

// Test Set, Update, and Delete functions for multiline values in .env file
func TestSetUpdateDeleteMultilineEnvFile(t *testing.T) {
	// Set the environment variable to point to the test.env file
	envFilePath := "test.env"
	os.Setenv("HARNESS_OUTPUT_SECRET_FILE", envFilePath)

	// Setup: Create a temporary .env file with multiline value for a key
	envContent := map[string]string{}
	err := createEnvFile(t, envFilePath, envContent)
	assert.NoError(t, err)

	// Defer file removal after the test completes
	defer os.Remove(envFilePath)

	// Set a new multiline value
	err = SetSecret("KEY1", "line1\nline2\nline3")
	assert.NoError(t, err)

	// Verify the multiline update
	data, err := v3.Read(envFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "line1\nline2\nline3", data["KEY1"])

	err = UpdateSecret("KEY1", "line4\nline5\nline6")
	assert.NoError(t, err)

	// Verify the multiline update
	data, err = v3.Read(envFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "line4\nline5\nline6", data["KEY1"])

	// Delete the key
	err = DeleteSecret("KEY1")
	assert.NoError(t, err)

	// Verify deletion
	data, err = v3.Read(envFilePath)
	assert.NoError(t, err)
	assert.NotContains(t, data, "KEY1")

	// Clean up
	defer os.Unsetenv("HARNESS_OUTPUT_SECRET_FILE")
}

// Test invalid cases for setting a secret in .out (should not allow multiline)
func TestSetInvalidMultilineOut(t *testing.T) {
	// Set the environment variable to point to the test.out file
	outFilePath := "test.out"
	os.Setenv("DRONE_OUTPUT", outFilePath)

	// Setup: Create a temporary .out file
	outContent := []string{
		"KEY1 value1",
		"KEY2 value2",
	}
	err := createOutFile(t, outFilePath, outContent)
	assert.NoError(t, err)

	// Defer file removal after the test completes
	defer os.Remove(outFilePath)

	// Attempt to set a multiline value (should fail)
	err = SetOutput("KEY3", "line1\nline2")
	assert.Error(t, err)

	// Verify the content did not change
	lines, err := ReadLines(outFilePath)
	assert.NoError(t, err)

	// Assertions
	assert.Contains(t, lines, "KEY1 value1")
	assert.Contains(t, lines, "KEY2 value2")

	// Clean up
	defer os.Unsetenv("DRONE_OUTPUT")
}

// Test that the functions return the appropriate error when the file does not exist
func TestFileNotExistError(t *testing.T) {
	// Attempt to set a value when the .env file is missing
	err := SetSecret("KEY1", "value1")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "environment variable HARNESS_OUTPUT_SECRET_FILE is not set")
}
