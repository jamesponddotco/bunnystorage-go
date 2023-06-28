package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrFileExists is the error returned when the file already exists.
	ErrFileExists xerrors.Error = "file already exists"

	// ErrUnexpectedStatusCode is the error returned when the status code is not 200.
	ErrUnexpectedStatusCode xerrors.Error = "unexpected status code"
)

// DownloadAction is the action for the download command.
func DownloadAction(c *cli.Context) error {
	client, err := meta.Client(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Downloading %q to %q...\n", c.String("path"), c.String("file"))
	}

	object, resp, err := client.Download(context.Background(), c.String("path"), c.String("file"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if resp.Status != 200 {
		return fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.Status)
	}

	var (
		output = c.String("output")
		dir    = filepath.Dir(output)
	)

	if err = os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("%w", err)
	}

	if _, err = os.Stat(output); !os.IsNotExist(err) && !c.Bool("force") {
		return fmt.Errorf("%w: %s", ErrFileExists, output)
	}

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()

	if _, err = file.Write(object); err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("verbose") {
		fmt.Fprintf(c.App.Writer, "Downloaded %q to %q.\n", c.String("path"), c.String("file"))
	}

	return nil
}
