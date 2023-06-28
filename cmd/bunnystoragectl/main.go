package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/app"
)

func main() {
	os.Exit(app.Run())
}
