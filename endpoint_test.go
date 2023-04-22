package bunnystorage_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func TestEndpoint_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		endpoint bunnystorage.Endpoint
		expected string
	}{
		{
			name:     "Falkenstein",
			endpoint: bunnystorage.EndpointFalkenstein,
			expected: "https://storage.bunnycdn.com",
		},
		{
			name:     "New York",
			endpoint: bunnystorage.EndpointNewYork,
			expected: "https://ny.storage.bunnycdn.com",
		},
		{
			name:     "Los Angeles",
			endpoint: bunnystorage.EndpointLosAngeles,
			expected: "https://la.storage.bunnycdn.com",
		},
		{
			name:     "Singapore",
			endpoint: bunnystorage.EndpointSingapore,
			expected: "https://sg.storage.bunnycdn.com",
		},
		{
			name:     "Sydney",
			endpoint: bunnystorage.EndpointSydney,
			expected: "https://syd.storage.bunnycdn.com",
		},
		{
			name:     "Unknown",
			endpoint: bunnystorage.Endpoint(-1),
			expected: "https://storage.bunnycdn.com",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := tt.endpoint.String()
			if result != tt.expected {
				t.Errorf("Expected endpoint '%s', got '%s'", tt.expected, result)
			}
		})
	}
}
