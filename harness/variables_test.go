package harness

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper function to create a temporary file with a given extension and set the environment variable
func setupTestFile(t *testing.T, envVar, extension string) string {
	t.Helper()
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test."+extension)
	err := os.WriteFile(tmpFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	os.Setenv(envVar, tmpFile)
	return tmpFile
}

// Test for SetSecret with both .env and .out file extensions
func TestSetSecret(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, HarnessOutputSecretFile, "env")

	err := SetSecret("testKey", "testValue")
	if err != nil {
		t.Fatalf("SetSecret failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read .env file: %v", err)
	}

	if string(content) != "testKey=testValue\n" {
		t.Errorf("Expected 'testKey=testValue', got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, HarnessOutputSecretFile, "out")

	err = SetSecret("testKey", "testValue")
	if err != nil {
		t.Fatalf("SetSecret failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read .out file: %v", err)
	}

	if string(content) != "testKey testValue\n" {
		t.Errorf("Expected 'testKey testValue', got: %s", string(content))
	}
}

// Test for UpdateSecret with both .env and .out file extensions
func TestUpdateSecret(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, HarnessOutputSecretFile, "env")

	// Write initial content
	err := os.WriteFile(envFilePath, []byte("testKey=oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .env file content: %v", err)
	}

	err = UpdateSecret("testKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateSecret failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read .env file: %v", err)
	}

	if string(content) != "testKey=updatedValue\n" {
		t.Errorf("Expected 'testKey=updatedValue', got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, HarnessOutputSecretFile, "out")

	// Write initial content
	err = os.WriteFile(outFilePath, []byte("testKey oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .out file content: %v", err)
	}

	err = UpdateSecret("testKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateSecret failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read .out file: %v", err)
	}

	if string(content) != "testKey updatedValue\n" {
		t.Errorf("Expected 'testKey updatedValue', got: %s", string(content))
	}
}

// Test for DeleteSecret with both .env and .out file extensions
func TestDeleteSecret(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, HarnessOutputSecretFile, "env")

	// Write initial content
	err := os.WriteFile(envFilePath, []byte("testKey=testValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .env file content: %v", err)
	}

	err = DeleteSecret("testKey")
	if err != nil {
		t.Fatalf("DeleteSecret failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read .env file: %v", err)
	}

	if string(content) != "" {
		t.Errorf("Expected .env file to be empty, got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, HarnessOutputSecretFile, "out")

	// Write initial content
	err = os.WriteFile(outFilePath, []byte("testKey testValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .out file content: %v", err)
	}

	err = DeleteSecret("testKey")
	if err != nil {
		t.Fatalf("DeleteSecret failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read .out file: %v", err)
	}

	if string(content) != "" {
		t.Errorf("Expected .out file to be empty, got: %s", string(content))
	}
}

// Test for SetErrorMetadata with both .env and .out file extensions
func TestSetErrorMetadata(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, MetadataFile, "env")

	err := SetErrorMetadata("errorMessage", "errorCode", "errorCategory")
	if err != nil {
		t.Fatalf("SetErrorMetadata failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read .env file: %v", err)
	}

	expected := "ERROR_MESSAGE=errorMessage\nERROR_CODE=errorCode\nERROR_CATEGORY=errorCategory\n"
	if string(content) != expected {
		t.Errorf("Expected '%s', got: %s", expected, string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, MetadataFile, "out")

	err = SetErrorMetadata("errorMessage", "errorCode", "errorCategory")
	if err != nil {
		t.Fatalf("SetErrorMetadata failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read .out file: %v", err)
	}

	expectedOut := "ERROR_MESSAGE errorMessage\nERROR_CODE errorCode\nERROR_CATEGORY errorCategory\n"
	if string(content) != expectedOut {
		t.Errorf("Expected '%s', got: %s", expectedOut, string(content))
	}
}

// Test for SetOutput with both .env and .out file extensions
func TestSetOutput(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, DroneOutputFile, "env")
	err := SetOutput("outputKey", "outputValue")
	if err != nil {
		t.Fatalf("SetOutput failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", envFilePath, err)
	}
	if string(content) != "outputKey=outputValue\n" {
		t.Errorf("Expected 'outputKey=outputValue', got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, DroneOutputFile, "out")
	err = SetOutput("outputKey", "outputValue")
	if err != nil {
		t.Fatalf("SetOutput failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", outFilePath, err)
	}
	if string(content) != "outputKey outputValue\n" {
		t.Errorf("Expected 'outputKey outputValue', got: %s", string(content))
	}
}

// Test for UpdateOutput with both .env and .out file extensions
func TestUpdateOutput(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, DroneOutputFile, "env")
	err := os.WriteFile(envFilePath, []byte("outputKey=oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .env file content: %v", err)
	}
	err = UpdateOutput("outputKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateOutput failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", envFilePath, err)
	}
	if string(content) != "outputKey=updatedValue\n" {
		t.Errorf("Expected 'outputKey=updatedValue', got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, DroneOutputFile, "out")
	err = os.WriteFile(outFilePath, []byte("outputKey oldValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .out file content: %v", err)
	}
	err = UpdateOutput("outputKey", "updatedValue")
	if err != nil {
		t.Fatalf("UpdateOutput failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", outFilePath, err)
	}
	if string(content) != "outputKey updatedValue\n" {
		t.Errorf("Expected 'outputKey updatedValue', got: %s", string(content))
	}
}

// Test for DeleteOutput with both .env and .out file extensions
func TestDeleteOutput(t *testing.T) {
	// Test for .env file
	envFilePath := setupTestFile(t, DroneOutputFile, "env")
	err := os.WriteFile(envFilePath, []byte("outputKey=outputValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .env file content: %v", err)
	}
	err = DeleteOutput("outputKey")
	if err != nil {
		t.Fatalf("DeleteOutput failed: %v", err)
	}

	content, err := os.ReadFile(envFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", envFilePath, err)
	}
	if string(content) != "" {
		t.Errorf("Expected file to be empty, got: %s", string(content))
	}

	// Test for .out file
	outFilePath := setupTestFile(t, DroneOutputFile, "out")
	err = os.WriteFile(outFilePath, []byte("outputKey outputValue\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to write initial .out file content: %v", err)
	}
	err = DeleteOutput("outputKey")
	if err != nil {
		t.Fatalf("DeleteOutput failed: %v", err)
	}

	content, err = os.ReadFile(outFilePath)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", outFilePath, err)
	}
	if string(content) != "" {
		t.Errorf("Expected file to be empty, got: %s", string(content))
	}
}
func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		line          string
		ext           string
		expectedKey   string
		expectedValue string
	}{
		// .env cases
		{"key=value", ".env", "key", "value"},
		{"key= value", ".env", "key", "value"},
		{"key =value", ".env", "key", "value"},
		{"key = value", ".env", "key", "value"},
		{"key=", ".env", "key", ""},
		{"key", ".env", "key", ""},
		{"key=multi word value", ".env", "key", "multi word value"},
		{"key=  spaced value  ", ".env", "key", "spaced value"},
		{"key=first line\nsecond line", ".env", "key", "first line\nsecond line"},
		{"key=first line\nsecond line\nthird line", ".env", "key", "first line\nsecond line\nthird line"},

		// .out cases
		{"key value", ".out", "key", "value"},
		{"key  value", ".out", "key", "value"},
		{" key value ", ".out", "key", "value"},
		{"key ", ".out", "key", ""},
		{"key", ".out", "key", ""},
		{"key multi word value", ".out", "key", "multi word value"},
		{"key  spaced value  ", ".out", "key", "spaced value"},

		// Unsupported extension cases
		{"key=value", ".unknown", "", ""},
		{"key value", ".unknown", "", ""},
	}

	for _, test := range tests {
		t.Run(test.line+"_"+test.ext, func(t *testing.T) {
			key, value := ParseKeyValue(test.line, test.ext)
			if key != test.expectedKey || value != test.expectedValue {
				t.Errorf("For line '%s' and ext '%s': expected ('%s', '%s'), got ('%s', '%s')",
					test.line, test.ext, test.expectedKey, test.expectedValue, key, value)
			}
		})
	}
}
