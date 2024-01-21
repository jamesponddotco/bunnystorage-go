package bunnystorage

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/build"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
)

const (
	// ErrInvalidConfig is returned when Config is invalid.
	ErrInvalidConfig xerrors.Error = "invalid config"

	// ErrInvalidEndpoint is returned when an endpoint is invalid.
	ErrInvalidEndpoint xerrors.Error = "invalid endpoint"

	// ErrStorageZoneRequired is returned when a Config is created without a
	// storage zone.
	ErrStorageZoneRequired xerrors.Error = "storage zone required"

	// ErrStorageZoneNameRequired is returned when a storage zone is created
	// without a name.
	ErrStorageZoneNameRequired xerrors.Error = "storage zone name required"

	// ErrStorageZoneKeyRequired is returned when a storage zone is created
	// without an API key.
	ErrStorageZoneKeyRequired xerrors.Error = "storage zone key required"

	// ErrEndpointRequired is returned when a Config is created without an
	// endpoint.
	ErrEndpointRequired xerrors.Error = "endpoint required"
)

// Default values for the Config struct.
const (
	DefaultMaxRetries int           = 3
	DefaultTimeout    time.Duration = 60 * time.Second
)

const (
	OperationRead Operation = iota
	OperationWrite
)

// Operation represents an operation that can be performed on a Bunny.net
// Storage API.
type Operation int

// Config holds the basic configuration for the Bunny.net Storage API.
type Config struct {
	// Logger is the structured logger to use for logging information about API
	// requests and responses.
	Logger *slog.Logger

	// StorageZone is the name of the storage zone to connect to.
	StorageZone string

	// Key is the API key used to authenticate with the API. The storage zone
	// password also doubles as your key.
	Key string

	// ReadOnlyKey is the read-only API key used to authenticate with the API.
	// This key is optional and only used for read-only operations.
	ReadOnlyKey string

	// UserAgent is the user agent to use when making HTTP requests to the API.
	UserAgent string

	// Endpoint is the endpoint to use for the API.
	Endpoint Endpoint

	// MaxRetries specifies the maximum number of times to retry a request if it
	// fails due to rate limiting.
	//
	// This field is optional.
	MaxRetries int

	// Timeout is the time limit for requests made by the client to the  API.
	//
	// This field is optional.
	Timeout time.Duration

	// mu protects Config initialization.
	mu sync.Mutex
}

// AccessKey returns the API key to use for the given operation.
func (c *Config) AccessKey(op Operation) string {
	if op == OperationRead && c.ReadOnlyKey != "" {
		return c.ReadOnlyKey
	}

	if c.Key != "" {
		return c.Key
	}

	return ""
}

// init initializes missing Config fields with their default values.
func (c *Config) init() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.UserAgent == "" {
		c.UserAgent = build.UserAgent
	}

	if c.MaxRetries < 1 {
		c.MaxRetries = DefaultMaxRetries
	}

	if c.Timeout < 1 {
		c.Timeout = DefaultTimeout
	}
}

// validate returns an error if the config is invalid.
func (c *Config) validate() error {
	if c.StorageZone == "" {
		return ErrStorageZoneRequired
	}

	if c.Key == "" {
		return ErrStorageZoneKeyRequired
	}

	if c.Endpoint == 0 {
		return ErrEndpointRequired
	}

	if !c.Endpoint.IsValid() {
		return fmt.Errorf("%w: %d", ErrInvalidEndpoint, c.Endpoint)
	}

	return nil
}
