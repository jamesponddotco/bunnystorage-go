package testutil

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
)

const (
	// ErrMissingEnvVar is returned when an environment variable is missing.
	ErrMissingEnvVar xerrors.Error = "missing environment variable"

	// ErrTestClient is returned when the test client fails to initialize.
	ErrTestClient xerrors.Error = "failed to initialize test client"
)

// MockServerAddr is the address of the mock server.
const MockServerAddr string = "localhost:62769"

// SetupClient sets up a test client for integration tests.
//
// The following environment variables are required:
// - BUNNY_STORAGE_ZONE
// - BUNNY_READ_API_KEY
// - BUNNY_WRITE_API_KEY
//
// The test will fail if any of them are empty or not set.
func SetupClient() (client *bunnystorage.Client, err error) {
	zone, ok := os.LookupEnv("BUNNY_STORAGE_ZONE")
	if !ok || zone == "" {
		return nil, fmt.Errorf("%w: BUNNY_STORAGE_ZONE", ErrMissingEnvVar)
	}

	readKey, ok := os.LookupEnv("BUNNY_READ_API_KEY")
	if !ok || readKey == "" {
		return nil, fmt.Errorf("%w: BUNNY_READ_API_KEY", ErrMissingEnvVar)
	}

	writeKey, ok := os.LookupEnv("BUNNY_WRITE_API_KEY")
	if !ok || writeKey == "" {
		return nil, fmt.Errorf("%w: BUNNY_WRITE_API_KEY", ErrMissingEnvVar)
	}

	cfg := &bunnystorage.Config{
		StorageZone: zone,
		Key:         writeKey,
		ReadOnlyKey: readKey,
		Endpoint:    bunnystorage.EndpointFalkenstein,
		Debug:       true,
	}

	client, err = bunnystorage.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrTestClient, err)
	}

	return client, nil
}

// SetupMockClient sets up a mock client for mocking tests.
func SetupMockClient(t *testing.T) *bunnystorage.Client {
	t.Helper()

	cfg := &bunnystorage.Config{
		StorageZone: "mock",
		Key:         "mock",
		ReadOnlyKey: "mock",
		Endpoint:    bunnystorage.EndpointLocalhost,
		Debug:       true,
	}

	client, err := bunnystorage.NewClient(cfg)
	if err != nil {
		t.Fatalf("failed to initialize test client: %v", err)
	}

	return client
}

// SetupMockServer sets up a mock server for mocking tests.
func SetupMockServer(t *testing.T) (mux *http.ServeMux, teardown func()) {
	t.Helper()

	mux = http.NewServeMux()
	srv := httptest.NewUnstartedServer(mux)

	var err error

	srv.Listener, err = net.Listen("tcp", MockServerAddr)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	srv.Start()

	return mux, srv.Close
}

// SetupFile sets up a simple text file for use in integration tests.
func SetupFile(t *testing.T) (name string, size int64, err error) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		return "", 0, fmt.Errorf("%w", err)
	}

	if _, err = tempFile.WriteString("Hello, tester!"); err != nil {
		return "", 0, fmt.Errorf("%w", err)
	}

	fileInfo, err := tempFile.Stat()
	if err != nil {
		return "", 0, fmt.Errorf("%w", err)
	}

	return tempFile.Name(), fileInfo.Size(), nil
}
