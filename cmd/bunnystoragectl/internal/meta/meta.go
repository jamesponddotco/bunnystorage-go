// Package meta provides build information and other metadata for the application.
package meta

import (
	"fmt"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"github.com/urfave/cli/v2"
)

const (
	// Name is the name of the application.
	Name = "bunnystoragectl"

	// Version is the version of the application.
	Version = "0.1.0"

	// Description is the description of the application.
	Description = "A command line interface for BunnyStorage"

	// Website is the website of the application.
	Website = "https://sr.ht/~jamesponddotco/bunnystorage-go/"
)

// Client returns a client for the BunnyStorage API that the application uses to
// interact with the service.
func Client(c *cli.Context) (*bunnystorage.Client, error) {
	cfg := &bunnystorage.Config{
		Application: &bunnystorage.Application{
			Name:    Name,
			Version: Version,
			Contact: Website,
		},
		StorageZone: c.String("storage-zone"),
		Key:         c.String("key"),
		Endpoint:    bunnystorage.Parse(c.String("endpoint")),
		Timeout:     c.Duration("timeout"),
	}

	client, err := bunnystorage.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return client, nil
}
