package app

import (
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
	"github.com/urfave/cli/v2"
)

// UploadAction is the action for the upload command.
func UploadAction(c *cli.Context) error {
	client, err := meta.Client(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Uploading %q to %q...\n", c.String("file"), c.String("path"))
	}

	file, err := os.Open(c.String("file"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	filename := filepath.Base(c.String("file"))

	checksum, err := bunnystorage.ComputeSHA256(file)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err = file.Seek(0, 0); err != nil {
		return fmt.Errorf("%w", err)
	}

	_, err = client.Upload(c.Context, c.String("path"), filename, checksum, file)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Upload complete.\n")
	}

	return nil
}
