package harness

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper function to create a temporary file and set the environment variable
func setupTestFile(t *testing.T, envVar string) string {
	t.Helper()
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.env")
	err := os.WriteFile(tmpFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	os.Setenv(envVar, tmpFile)
	return tmpFile
}

// Helper function to create a temporary .out file and set the environment variable
func setupOutTestFile(t *testing.T, envVar string) string {
	t.Helper()
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.out")
	err := os.WriteFile(tmpFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	os.Setenv(envVar, tmpFile)
	return tmpFile
}

// Test for SetSecret
func TestSetSecret(t *testing.T) {
	filePath := setupTestFile(t, HarnessOutputSecretFile)

	err := SetSecret("testKey", "testValue")
	if err != nil {
		t.Fatalf("SetSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "testKey=testValue\n" {
		t.Errorf("Expected 'testKey=testValue', got: %s", string(content))
	}
}

// Test for UpdateSecret
func TestUpdateSecret(t *testing.T) {
	filePath := setupTestFile(t, HarnessOutputSecretFile)

	// Write initial content
	err := os.WriteFile(filePath, []byte("testKey=oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial file content: %v", err)
	}

	err = UpdateSecret("testKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "testKey=updatedValue\n" {
		t.Errorf("Expected 'testKey=updatedValue', got: %s", string(content))
	}
}

// Test for DeleteSecret
func TestDeleteSecret(t *testing.T) {
	filePath := setupTestFile(t, HarnessOutputSecretFile)

	// Write initial content
	err := os.WriteFile(filePath, []byte("testKey=testValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial file content: %v", err)
	}

	err = DeleteSecret("testKey")
	if err != nil {
		t.Fatalf("DeleteSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "" {
		t.Errorf("Expected file to be empty, got: %s", string(content))
	}
}

// Test for SetOutput
func TestSetOutput(t *testing.T) {
	filePath := setupTestFile(t, DroneOutputFile)

	err := SetOutput("outputKey", "outputValue")
	if err != nil {
		t.Fatalf("SetOutput failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "outputKey=outputValue\n" {
		t.Errorf("Expected 'outputKey=outputValue', got: %s", string(content))
	}
}

// Test for SetErrorMetadata
func TestSetErrorMetadata(t *testing.T) {
	filePath := setupTestFile(t, MetadataFile)

	err := SetErrorMetadata("errorMessage", "errorCode", "errorCategory")
	if err != nil {
		t.Fatalf("SetErrorMetadata failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expected := "ERROR_MESSAGE=errorMessage\nERROR_CODE=errorCode\nERROR_CATEGORY=errorCategory\n"
	if string(content) != expected {
		t.Errorf("Expected '%s', got: %s", expected, string(content))
	}
}

// Test for SetSecret with .out file
func TestSetSecretOutFile(t *testing.T) {
	filePath := setupOutTestFile(t, HarnessOutputSecretFile)

	err := SetSecret("testKey", "testValue")
	if err != nil {
		t.Fatalf("SetSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "testKey testValue\n" {
		t.Errorf("Expected 'testKey testValue', got: %s", string(content))
	}
}

// Test for UpdateSecret with .out file
func TestUpdateSecretOutFile(t *testing.T) {
	filePath := setupOutTestFile(t, HarnessOutputSecretFile)

	// Write initial content
	err := os.WriteFile(filePath, []byte("testKey oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial file content: %v", err)
	}

	err = UpdateSecret("testKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "testKey updatedValue\n" {
		t.Errorf("Expected 'testKey updatedValue', got: %s", string(content))
	}
}

// Test for DeleteSecret with .out file
func TestDeleteSecretOutFile(t *testing.T) {
	filePath := setupOutTestFile(t, HarnessOutputSecretFile)

	// Write initial content
	err := os.WriteFile(filePath, []byte("testKey testValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial file content: %v", err)
	}

	err = DeleteSecret("testKey")
	if err != nil {
		t.Fatalf("DeleteSecret failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(content) != "" {
		t.Errorf("Expected file to be empty, got: %s", string(content))
	}
}

// Test for SetErrorMetadata with .out file
func TestSetErrorMetadataOutFile(t *testing.T) {
	filePath := setupOutTestFile(t, MetadataFile)

	err := SetErrorMetadata("errorMessage", "errorCode", "errorCategory")
	if err != nil {
		t.Fatalf("SetErrorMetadata failed: %v", err)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expected := "ERROR_MESSAGE errorMessage\nERROR_CODE errorCode\nERROR_CATEGORY errorCategory\n"
	if string(content) != expected {
		t.Errorf("Expected '%s', got: %s", expected, string(content))
	}
}
