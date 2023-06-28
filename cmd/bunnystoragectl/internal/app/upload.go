package app

import (
	"fmt"

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

	_, err = client.Upload(c.Context, c.String("path"), c.String("file"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Upload complete.\n")
	}

	return nil
}
