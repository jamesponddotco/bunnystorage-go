package bunnystorage_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"git.sr.ht/~jamesponddotco/httpx-go"
)

func TestApplication_UserAgent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		application *bunnystorage.Application
		expected    *httpx.UserAgent
	}{
		{
			name: "Valid application",
			application: &bunnystorage.Application{
				Name:    "TestApp",
				Version: "1.0.0",
				Contact: "test@example.com",
			},
			expected: &httpx.UserAgent{
				Token:   "TestApp",
				Version: "1.0.0",
				Comment: []string{
					"test@example.com",
				},
			},
		},
		{
			name: "Empty application name",
			application: &bunnystorage.Application{
				Name:    "",
				Version: "1.0.0",
				Contact: "test@example.com",
			},
			expected: &httpx.UserAgent{},
		},
		{
			name: "Empty application version",
			application: &bunnystorage.Application{
				Name:    "TestApp",
				Version: "",
				Contact: "test@example.com",
			},
			expected: &httpx.UserAgent{},
		},
		{
			name: "Empty application contact",
			application: &bunnystorage.Application{
				Name:    "TestApp",
				Version: "1.0.0",
				Contact: "",
			},
			expected: &httpx.UserAgent{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ua := tt.application.UserAgent()
			if ua.Token != tt.expected.Token || ua.Version != tt.expected.Version || len(ua.Comment) != len(tt.expected.Comment) {
				t.Errorf("Expected user agent %v, but got %v", tt.expected, ua)
			}

			if len(ua.Comment) > 0 && ua.Comment[0] != tt.expected.Comment[0] {
				t.Errorf("Expected user agent comment %s, but got %s", tt.expected.Comment[0], ua.Comment[0])
			}
		})
	}
}

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
