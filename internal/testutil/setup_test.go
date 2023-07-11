package testutil_test

import (
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/testutil"
)

func TestSetupClient(t *testing.T) {
	tests := []struct {
		name          string
		zone          string
		readKey       string
		writeKey      string
		expectedError bool
	}{
		{
			name:          "Valid",
			zone:          "dummy_zone",
			readKey:       "dummy_read_key",
			writeKey:      "dummy_write_key",
			expectedError: false,
		},
		{
			name:          "Missing BUNNY_STORAGE_ZONE",
			zone:          "",
			readKey:       "dummy_read_key",
			writeKey:      "dummy_write_key",
			expectedError: true,
		},
		{
			name:          "Missing BUNNY_READ_API_KEY",
			zone:          "dummy_zone",
			readKey:       "",
			writeKey:      "dummy_write_key",
			expectedError: true,
		},
		{
			name:          "Missing BUNNY_WRITE_API_KEY",
			zone:          "dummy_zone",
			readKey:       "dummy_read_key",
			writeKey:      "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BUNNY_STORAGE_ZONE", tt.zone)
			t.Setenv("BUNNY_READ_API_KEY", tt.readKey)
			t.Setenv("BUNNY_WRITE_API_KEY", tt.writeKey)

			client, err := testutil.SetupClient()
			if tt.expectedError && err == nil {
				t.Errorf("expected error, got nil")

				return
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectedError && client == nil {
				t.Errorf("expected non-nil client")
			}
		})
	}
}

func TestSetupMockClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		wantPanic bool
	}{
		{
			name:      "setup_success",
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("SetupMockClient() panic = %v, wantPanic = %v", r != nil, tt.wantPanic)
				}
			}()

			client := testutil.SetupMockClient(t)

			if client == nil {
				t.Errorf("SetupMockClient() client = %v, want client not nil", client)
			}
		})
	}
}

func TestSetupFile(t *testing.T) {
	t.Parallel()

	t.Run("creates file with content", func(t *testing.T) {
		t.Parallel()

		name, size, err := testutil.SetupFile(t)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		content, err := os.ReadFile(name)
		if err != nil {
			t.Fatalf("failed to read test file: %v", err)
		}

		expectedContent := "Hello, tester!"
		if string(content) != expectedContent {
			t.Errorf("expected content to be %q, got %q", expectedContent, content)
		}

		expectedSize := int64(len(expectedContent))
		if size != expectedSize {
			t.Errorf("expected size to be %d, got %d", expectedSize, size)
		}

		err = os.Remove(name)
		if err != nil {
			t.Fatalf("failed to delete test file: %v", err)
		}
	})
}
