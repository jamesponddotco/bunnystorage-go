package bunnystorage_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

type ErrorReader struct{}

func (*ErrorReader) Read(_ []byte) (n int, err error) {
	return 0, fmt.Errorf("forced reader error")
}

func TestComputeSHA256(t *testing.T) {
	t.Parallel()

	fileContent := []byte("Test data for hashing.")

	var (
		expectedHash    = sha256.Sum256(fileContent)
		expectedHashHex = hex.EncodeToString(expectedHash[:])
	)

	emptyHash := sha256.New()
	emptyHashHex := hex.EncodeToString(emptyHash.Sum(nil))

	testCases := []struct {
		name         string
		reader       io.Reader
		expectedHash string
		expectError  bool
	}{
		{
			name:         "ValidFileContent",
			reader:       bytes.NewReader(fileContent),
			expectedHash: expectedHashHex,
			expectError:  false,
		},
		{
			name:         "EmptyFileContent",
			reader:       bytes.NewReader([]byte{}),
			expectedHash: emptyHashHex,
			expectError:  false,
		},
		{
			name:         "ReaderError",
			reader:       &ErrorReader{},
			expectedHash: "",
			expectError:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			hash, err := bunnystorage.ComputeSHA256(tc.reader)
			if (err != nil) != tc.expectError {
				t.Fatalf("Expected error: %v, got: %v", tc.expectError, err)
			}

			if err == nil && hash != tc.expectedHash {
				t.Errorf("Expected hash: %s, got: %s", tc.expectedHash, hash)
			}

			log.Printf("Hash: %s", hash)
		})
	}
}
