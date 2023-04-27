package bunnystorage

import (
	"fmt"
	"sync"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/build"
	"git.sr.ht/~jamesponddotco/httpx-go"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xlog"
)

const (
	// ErrInvalidApplication is returned when an application is invalid.
	ErrInvalidApplication xerrors.Error = "invalid application"

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

	// ErrApplicationRequired is returned when a Config is created without an
	// application.
	ErrApplicationRequired xerrors.Error = "application required"

	// ErrApplicationNameRequired is returned when an application is created
	// without a name.
	ErrApplicationNameRequired xerrors.Error = "application name required"

	// ErrApplicationVersionRequired is returned when an application is created
	// without a version.
	ErrApplicationVersionRequired xerrors.Error = "application version required"

	// ErrApplicationContactRequired is returned when an application is created
	// without contact information.
	ErrApplicationContactRequired xerrors.Error = "application contact required"

	// ErrApplicationKeyRequired is returned when an application is created
	// without an API key.
	ErrApplicationKeyRequired xerrors.Error = "application key required"
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

// Logger defines the interface for logging. It is basically a thin wrapper
// around the standard logger which implements only a subset of the logger API.
type Logger interface {
	Printf(format string, v ...any)
}

// Operation represents an operation that can be performed on a Bunny.net
// Storage API.
type Operation int

// Application represents the application that is making requests to the API.
type Application struct {
	// Name is the name of the application.
	Name string

	// Version is the version of the application.
	Version string

	// Contact is the contact information for the application. Either an email
	// or an URL.
	Contact string
}

// DefaultApplication returns a new Application with default values.
func DefaultApplication() *Application {
	return &Application{
		Name:    build.Name,
		Version: build.Version,
		Contact: build.URL,
	}
}

// UserAgent returns the user agent string for the application.
func (a *Application) UserAgent() *httpx.UserAgent {
	if err := a.validate(); err != nil {
		return &httpx.UserAgent{}
	}

	return &httpx.UserAgent{
		Token:   a.Name,
		Version: a.Version,
		Comment: []string{
			a.Contact,
		},
	}
}

// Validate returns an error if the application is invalid.
func (a *Application) validate() error {
	if a.Name == "" {
		return fmt.Errorf("%w: %w", ErrInvalidApplication, ErrApplicationNameRequired)
	}

	if a.Version == "" {
		return fmt.Errorf("%w: %w", ErrInvalidApplication, ErrApplicationVersionRequired)
	}

	if a.Contact == "" {
		return fmt.Errorf("%w: %w", ErrInvalidApplication, ErrApplicationContactRequired)
	}

	return nil
}

// Config holds the basic configuration for the Bunny.net Storage API.
type Config struct {
	// Application is the application that is making requests to the API.
	Application *Application

	// Logger is the logger to use for logging requests when debugging.
	Logger Logger

	// StorageZone is the name of the storage zone to connect to.
	StorageZone string

	// Key is the API key used to authenticate with the API. The storage zone
	// password also doubles as your key.
	Key string

	// ReadOnlyKey is the read-only API key used to authenticate with the API.
	// This key is optional and only used for read-only operations.
	ReadOnlyKey string

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

	// Debug specifies whether or not to enable debug logging.
	//
	// This field is optional.
	Debug bool

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

	if c.Application == nil {
		c.Application = DefaultApplication()
	}

	if c.MaxRetries < 1 {
		c.MaxRetries = DefaultMaxRetries
	}

	if c.Timeout < 1 {
		c.Timeout = DefaultTimeout
	}

	if c.Logger == nil && c.Debug {
		c.Logger = xlog.DefaultZeroLogger
	}
}

// validate returns an error if the config is invalid.
func (c *Config) validate() error {
	if c.Application == nil {
		return ErrApplicationRequired
	}

	if err := c.Application.validate(); err != nil {
		return err
	}

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
