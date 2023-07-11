package testutil_test

import (
	"bytes"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/testutil"
)

func TestReadFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fileName  string
		wantError bool
	}{
		{
			name:      "existing_file",
			fileName:  "testdata/file.txt",
			wantError: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := testutil.ReadFile(t, tt.fileName)

			if !tt.wantError {
				want, err := os.ReadFile(tt.fileName)
				if err != nil {
					t.Fatalf("Failed to read control file: %v", err)
				}

				if !bytes.Equal(got, want) {
					t.Errorf("ReadFile() = %v, want %v", got, want)
				}
			}
		})
	}
}
