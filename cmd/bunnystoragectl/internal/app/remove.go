package app

import (
	"fmt"
	"net/http"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

// ErrFileNotFound is returned when the file is not found.
const ErrFileNotFound xerrors.Error = "failed to delete file"

// DeleteAction is the action for the delete command.
func DeleteAction(c *cli.Context) error {
	client, err := meta.Client(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Deleting file %q from path %q...\n", c.String("file"), c.String("path"))
	}

	resp, err := client.Delete(c.Context, c.String("path"), c.String("file"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if resp.Status != http.StatusOK {
		return fmt.Errorf("%w: %d", ErrFileNotFound, resp.Status)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Deleted file %q from path %q.\n", c.String("file"), c.String("path"))
	}

	return nil
}
