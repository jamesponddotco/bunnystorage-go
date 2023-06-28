// Package osutil provides useful functions and utilities for working with the
// operating system.
package osutil

import (
	"os"
	"path/filepath"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
)

// ApplicationConfig returns the path to the user's configuration file for the
// application.
func ApplicationConfig() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, meta.Name, "config.json")
}
