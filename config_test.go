package bunnystorage_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func TestConfig_AccessKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		config      *bunnystorage.Config
		operation   bunnystorage.Operation
		expectedKey string
	}{
		{
			name: "Read operation with read-only key available",
			config: &bunnystorage.Config{
				Key:         "main-key",
				ReadOnlyKey: "read-only-key",
			},
			operation:   bunnystorage.OperationRead,
			expectedKey: "read-only-key",
		},
		{
			name: "Read operation without read-only key",
			config: &bunnystorage.Config{
				Key: "main-key",
			},
			operation:   bunnystorage.OperationRead,
			expectedKey: "main-key",
		},
		{
			name: "Write operation with read-only key available",
			config: &bunnystorage.Config{
				Key:         "main-key",
				ReadOnlyKey: "read-only-key",
			},
			operation:   bunnystorage.OperationWrite,
			expectedKey: "main-key",
		},
		{
			name: "Write operation without read-only key",
			config: &bunnystorage.Config{
				Key: "main-key",
			},
			operation:   bunnystorage.OperationWrite,
			expectedKey: "main-key",
		},
		{
			name: "No keys available",
			config: &bunnystorage.Config{
				Key:         "",
				ReadOnlyKey: "",
			},
			operation:   bunnystorage.OperationRead,
			expectedKey: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			accessKey := tt.config.AccessKey(tt.operation)
			if accessKey != tt.expectedKey {
				t.Errorf("Expected access key %s, but got %s", tt.expectedKey, accessKey)
			}
		})
	}
}
