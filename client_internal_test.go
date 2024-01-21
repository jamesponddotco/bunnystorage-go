package bunnystorage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	t.Parallel()

	cfg := &Config{
		StorageZone: "my-storage-zone",
		Key:         "my-key",
		ReadOnlyKey: "my-read-only-key",
		Endpoint:    EndpointFalkenstein,
		Debug:       true,
	}

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}

	tests := []struct {
		name      string
		ctx       context.Context
		method    string
		uri       string
		headers   map[string]string
		body      io.Reader
		expectErr bool
	}{
		{
			name:   "GET request",
			ctx:    context.Background(),
			method: http.MethodGet,
			uri:    "https://example.com",
			headers: map[string]string{
				"Accept": "application/json",
			},
			body: http.NoBody,
		},
		{
			name:   "POST request",
			ctx:    context.Background(),
			method: http.MethodPost,
			uri:    "https://example.com",
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			body: bytes.NewBuffer([]byte(`{"key": "value"}`)),
		},
		{
			name:      "Invalid context",
			ctx:       nil,
			method:    http.MethodGet,
			uri:       "https://example.com",
			headers:   nil,
			body:      http.NoBody,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req, err := client.request(tt.ctx, tt.method, tt.uri, tt.headers, tt.body)

			if tt.expectErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("request() failed: %v", err)
			}

			if req.Method != tt.method {
				t.Errorf("expected method %s, got %s", tt.method, req.Method)
			}

			if req.URL.String() != tt.uri {
				t.Errorf("expected uri %s, got %s", tt.uri, req.URL.String())
			}

			for k, v := range tt.headers {
				if req.Header.Get(k) != v {
					t.Errorf("expected header %s with value %s, got %s", k, v, req.Header.Get(k))
				}
			}

			if req.Header.Get("User-Agent") != cfg.UserAgent {
				t.Errorf("expected User-Agent header %s, got %s", cfg.UserAgent, req.Header.Get("User-Agent"))
			}
		})
	}
}
