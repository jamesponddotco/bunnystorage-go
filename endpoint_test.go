package bunnystorage_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bunnystorage.Endpoint
	}{
		{
			name:     "default case",
			input:    "http://default.com",
			expected: bunnystorage.EndpointFalkenstein,
		},
		{
			name:     "storage.bunnycdn.com",
			input:    "http://storage.bunnycdn.com",
			expected: bunnystorage.EndpointFalkenstein,
		},
		{
			name:     "ny.storage.bunnycdn.com",
			input:    "http://ny.storage.bunnycdn.com",
			expected: bunnystorage.EndpointNewYork,
		},
		{
			name:     "la.storage.bunnycdn.com",
			input:    "http://la.storage.bunnycdn.com",
			expected: bunnystorage.EndpointLosAngeles,
		},
		{
			name:     "sg.storage.bunnycdn.com",
			input:    "http://sg.storage.bunnycdn.com",
			expected: bunnystorage.EndpointSingapore,
		},
		{
			name:     "syd.storage.bunnycdn.com",
			input:    "http://syd.storage.bunnycdn.com",
			expected: bunnystorage.EndpointSydney,
		},
		{
			name:     "uk.storage.bunnycdn.com",
			input:    "http://uk.storage.bunnycdn.com",
			expected: bunnystorage.EndpointLondon,
		},
		{
			name:     "se.storage.bunnycdn.com",
			input:    "http://se.storage.bunnycdn.com",
			expected: bunnystorage.EndpointStockholm,
		},
		{
			name:     "br.storage.bunnycdn.com",
			input:    "http://br.storage.bunnycdn.com",
			expected: bunnystorage.EndpointSaoPaulo,
		},
		{
			name:     "jh.storage.bunnycdn.com",
			input:    "http://jh.storage.bunnycdn.com",
			expected: bunnystorage.EndpointJohannesburg,
		},
		{
			name:     "localhost",
			input:    "http://localhost:62769",
			expected: bunnystorage.EndpointLocalhost,
		},
		{
			name:     "invalid url",
			input:    "://invalid.url",
			expected: bunnystorage.EndpointFalkenstein,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := bunnystorage.Parse(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

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
			name:     "London",
			endpoint: bunnystorage.EndpointLondon,
			expected: "https://uk.storage.bunnycdn.com",
		},
		{
			name:     "Stockholm",
			endpoint: bunnystorage.EndpointStockholm,
			expected: "https://se.storage.bunnycdn.com",
		},
		{
			name:     "SaoPaulo",
			endpoint: bunnystorage.EndpointSaoPaulo,
			expected: "https://br.storage.bunnycdn.com",
		},
		{
			name:     "Johannesburg",
			endpoint: bunnystorage.EndpointJohannesburg,
			expected: "https://jh.storage.bunnycdn.com",
		},
		{
			name:     "Localhost",
			endpoint: bunnystorage.EndpointLocalhost,
			expected: "http://localhost:62769",
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
