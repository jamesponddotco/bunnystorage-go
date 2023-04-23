// Package cryptoutil provides utilities to extend Go's crypto package.
package cryptoutil

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ComputeSHA256 returns the SHA256 hash of the given string as a hex string.
func ComputeSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var (
		hash    = hasher.Sum(nil)
		hashHex = hex.EncodeToString(hash)
	)

	return hashHex, nil
}
