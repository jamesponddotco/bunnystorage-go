// Package testutil provides utilities for testing.
package testutil

import (
	"os"
	"testing"
)

// ReadFile reads the named file and either returns its contents or fails the
// test.
func ReadFile(t *testing.T, name string) (file []byte) {
	t.Helper()

	file, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("failed to read test data: %v", err)
	}

	return file
}
