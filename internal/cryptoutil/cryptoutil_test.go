package cryptoutil_test

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/cryptoutil"
)

func TestComputeSHA256(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "cryptoutil_test_")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}

	defer t.Cleanup(func() {
		os.RemoveAll(tempDir)
	})

	var (
		fileContent = []byte("Test data for hashing.")
		filePath    = filepath.Join(tempDir, "testfile.txt")
	)

	if err := ioutil.WriteFile(filePath, fileContent, 0o600); err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}

	var (
		expectedHash    = sha256.Sum256(fileContent)
		expectedHashHex = hex.EncodeToString(expectedHash[:])
	)

	testCases := []struct {
		name          string
		filePath      string
		expectedHash  string
		expectedError bool
	}{
		{
			name:          "ValidFilePath",
			filePath:      filePath,
			expectedHash:  expectedHashHex,
			expectedError: false,
		},
		{
			name:          "InvalidFilePath",
			filePath:      filepath.Join(tempDir, "nonexistent.txt"),
			expectedHash:  "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := cryptoutil.ComputeSHA256(tc.filePath)
			if (err != nil) != tc.expectedError {
				t.Fatalf("Expected error: %v, got: %v", tc.expectedError, err)
			}

			if hash != tc.expectedHash {
				t.Errorf("Expected hash: %s, got: %s", tc.expectedHash, hash)
			}

			log.Printf("Hash: %s", hash)
		})
	}
}
