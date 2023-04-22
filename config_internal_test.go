package bunnystorage

import "testing"

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		config    *Config
		expectErr bool
	}{
		{
			name: "Valid config",
			config: &Config{
				Application: DefaultApplication(),
				StorageZone: "storage-zone",
				Key:         "api-key",
				Endpoint:    EndpointFalkenstein,
			},
			expectErr: false,
		},
		{
			name: "Missing application",
			config: &Config{
				StorageZone: "storage-zone",
				Key:         "api-key",
				Endpoint:    EndpointFalkenstein,
			},
			expectErr: true,
		},
		{
			name: "Invalid application",
			config: &Config{
				Application: &Application{
					Name:    "",
					Version: "1.0.0",
					Contact: "test@example.com",
				},
				StorageZone: "storage-zone",
				Key:         "api-key",
				Endpoint:    EndpointFalkenstein,
			},
			expectErr: true,
		},
		{
			name: "Missing storage zone",
			config: &Config{
				Application: DefaultApplication(),
				Key:         "api-key",
				Endpoint:    EndpointFalkenstein,
			},
			expectErr: true,
		},
		{
			name: "Missing API key",
			config: &Config{
				Application: DefaultApplication(),
				StorageZone: "storage-zone",
				Endpoint:    EndpointFalkenstein,
			},
			expectErr: true,
		},
		{
			name: "Missing endpoint",
			config: &Config{
				Application: DefaultApplication(),
				StorageZone: "storage-zone",
				Key:         "api-key",
			},
			expectErr: true,
		},
		{
			name: "Invalid endpoint URL",
			config: &Config{
				Application: DefaultApplication(),
				StorageZone: "storage-zone",
				Key:         "api-key",
				Endpoint:    Endpoint(999),
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.config.validate()
			if tt.expectErr && err == nil {
				t.Errorf("Expected an error, but got nil")
			} else if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}
