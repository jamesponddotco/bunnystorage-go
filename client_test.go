package bunnystorage_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func TestClient_NewClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		giveConfig *bunnystorage.Config
		wantErr    bool
	}{
		{
			name: "success with default Application",
			giveConfig: &bunnystorage.Config{
				Application: bunnystorage.DefaultApplication(),
				StorageZone: "my-storage-zone",
				Key:         "my-key",
				ReadOnlyKey: "my-read-only-key",
				Endpoint:    bunnystorage.EndpointFalkenstein,
				MaxRetries:  bunnystorage.DefaultMaxRetries,
				Timeout:     bunnystorage.DefaultTimeout,
			},
			wantErr: false,
		},
		{
			name: "success with custom Application and Config",
			giveConfig: &bunnystorage.Config{
				Application: nil,
				StorageZone: "my-storage-zone",
				Key:         "my-key",
				ReadOnlyKey: "my-read-only-key",
				Endpoint:    bunnystorage.EndpointFalkenstein,
			},
			wantErr: false,
		},
		{
			name: "error with default Application",
			giveConfig: &bunnystorage.Config{
				Application: bunnystorage.DefaultApplication(),
			},
			wantErr: true,
		},
		{
			name:       "error with nil Config",
			giveConfig: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := bunnystorage.NewClient(tt.giveConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
