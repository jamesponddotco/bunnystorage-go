// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run() int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "storage-zone",
			Aliases:  []string{"z"},
			Usage:    "storage zone",
			Required: true,
			EnvVars:  []string{"BUNNY_STORAGE_ZONE"},
		},
		&cli.StringFlag{
			Name:     "key",
			Aliases:  []string{"k"},
			Usage:    "api key",
			Required: true,
			EnvVars:  []string{"BUNNY_KEY"},
		},
		&cli.StringFlag{
			Name:    "endpoint",
			Aliases: []string{"e"},
			Usage:   "api endpoint",
			Value:   "https://storage.bunnycdn.com",
			EnvVars: []string{"BUNNY_ENDPOINT"},
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"V"},
			Usage:   "enable verbose output",
			Value:   false,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list files and directories on storage zone",
			Action:  ListAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "directory path to your files",
					Value:   "/",
				},
				&cli.BoolFlag{
					Name:    "json",
					Aliases: []string{"j"},
					Usage:   "enable json output",
					Value:   false,
				},
			},
		},
		{
			Name:    "upload",
			Aliases: []string{"cp"},
			Usage:   "upload file to storage zone",
			Action:  UploadAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "directory path to your files",
					Value:   "/",
				},
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "file to upload",
					Required: true,
				},
			},
		},
		{
			Name:    "download",
			Aliases: []string{"dl"},
			Usage:   "download file from storage zone",
			Action:  DownloadAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "directory path to your files",
					Value:   "/",
				},
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "file to download",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "output",
					Aliases:  []string{"o"},
					Usage:    "output file",
					Required: true,
				},
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{"F"},
					Usage:   "force overwrite",
					Value:   false,
				},
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove file from storage zone",
			Action:  DeleteAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "directory path to your files",
					Value:   "/",
				},
				&cli.StringFlag{
					Name:     "file",
					Aliases:  []string{"f"},
					Usage:    "file to remove",
					Required: true,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
