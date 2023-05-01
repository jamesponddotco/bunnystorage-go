# bunnystorage-go

[![Go Documentation](https://godocs.io/git.sr.ht/~jamesponddotco/bunnystorage-go?status.svg)](https://godocs.io/git.sr.ht/~jamesponddotco/bunnystorage-go)
[![Go Report Card](https://goreportcard.com/badge/git.sr.ht/~jamesponddotco/bunnystorage-go)](https://goreportcard.com/report/git.sr.ht/~jamesponddotco/bunnystorage-go)
[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/bunnystorage-go.svg)](https://builds.sr.ht/~jamesponddotco/bunnystorage-go?)

Package `bunnystorage` is a simple and easy-to-use package for
interacting with the [Bunny.net Edge Storage
API](https://docs.bunny.net/reference/storage-api). It provides a
convenient way to manage files in your storage zones.

## Installation

To install `bunnystorage`, run:

```console
go get git.sr.ht/~jamesponddotco/bunnystorage-go
```

## Usage

```go
package main

import (
	"context"
	"log"
	"os"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
)

func main() {
	readOnlyKey, ok := os.LookupEnv("BUNNYNET_READ_API_KEY")
	if !ok {
		log.Fatal("missing env var: BUNNYNET_READ_API_KEY")
	}

	readWriteKey, ok := os.LookupEnv("BUNNYNET_WRITE_API_KEY")
	if !ok {
		log.Fatal("missing env var: BUNNYNET_WRITE_API_KEY")
	}

	// Create new Config to be initialize a Client.
	cfg := &bunnystorage.Config{
		Application: &bunnystorage.Application{
			Name:    "My Application",
			Version: "1.0.0",
			Contact: "contact@example.com",
		},
		StorageZone: "my-storage-zone",
		Key:         readWriteKey,
		ReadOnlyKey: readOnlyKey,
		Endpoint:    bunnystorage.EndpointFalkenstein,
	}

	// Create a new Client instance with the given Config.
	client, err := bunnystorage.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// List files in the storage zone.
	files, _, err := client.List(context.Background(), "/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Printf("File: %s, Size: %d\n", file.ObjectName, file.Length)
	}
}
```

For more examples and usage details, please [check the Go reference
documentation](https://godocs.io/git.sr.ht/~jamesponddotco/bunnystorage-go).

## Contributing

Anyone can help make `bunnystorage` better. Check out [the contribution
guidelines](https://git.sr.ht/~jamesponddotco/bunnystorage-go/tree/trunk/item/CONTRIBUTING.md)
for more information.

## Resources

The following resources are available:

- [Package documentation](https://godocs.io/git.sr.ht/~jamesponddotco/bunnystorage-go).
- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/bunnystorage-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/bunnystorage-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/bunnystorage).

---

Released under the [MIT License](LICENSE.md).
