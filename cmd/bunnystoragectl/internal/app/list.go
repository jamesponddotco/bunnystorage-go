package app

import (
	"encoding/json"
	"fmt"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/cmd/bunnystoragectl/internal/meta"
	"github.com/urfave/cli/v2"
)

// ListAction is the action for the list command.
func ListAction(c *cli.Context) error {
	client, err := meta.Client(c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	files, _, err := client.List(c.Context, c.String("path"))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if c.Bool("json") {
		data, _ := json.MarshalIndent(files, "", "    ")

		fmt.Fprintln(c.App.Writer, string(data))
	} else {
		fmt.Fprintf(c.App.Writer, "total %d\n", len(files))

		var (
			maxDateLen = 12
			maxSizeLen = 1
			maxNameLen = 40
		)

		// Initialize a slice to store the formatted strings.
		formattedStrings := make([]string, 0, len(files))

		for _, file := range files {
			// Parse the date from RFC3339 format
			t, err := time.Parse("2006-01-02T15:04:05.999", file.LastChanged)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			var (
				formattedDate = t.Format("Jan 02 15:04")
				formattedSize = fmt.Sprintf("%d", file.Length)
			)

			var (
				dateLen = len(formattedDate)
				sizeLen = len(formattedSize)
				nameLen = len(file.ObjectName)
			)

			if dateLen > maxDateLen {
				maxDateLen = dateLen
			}
			if sizeLen > maxSizeLen {
				maxSizeLen = sizeLen
			}
			if nameLen > maxNameLen {
				maxNameLen = nameLen
			}

			formattedStrings = append(
				formattedStrings,
				fmt.Sprintf(
					"%-*s %-*s %-*s",
					maxSizeLen,
					formattedSize,
					maxDateLen,
					formattedDate,
					maxNameLen,
					file.ObjectName,
				),
			)
		}

		for _, str := range formattedStrings {
			fmt.Fprintln(c.App.Writer, str)
		}
	}

	return nil
}
