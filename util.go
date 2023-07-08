package bunnystorage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// ComputeSHA256 returns the SHA256 hash of the given string as a hex string.
func ComputeSHA256(r io.Reader) (string, error) {
	hasher := sha256.New()
	if _, err := io.Copy(hasher, r); err != nil {
		return "", fmt.Errorf("%w", err)
	}

	var (
		hash    = hasher.Sum(nil)
		hashHex = hex.EncodeToString(hash)
	)

	return hashHex, nil
}
