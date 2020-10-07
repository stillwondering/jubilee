package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stillwondering/jubilee"
	"github.com/stillwondering/jubilee/cmd/jubilee/csv"

	"github.com/urfave/cli/v2"
)

type RecordProvider interface {
	Fetch() ([]jubilee.BirthdayRecord, error)
}

func main() {
	app := &cli.App{
		Name:  "jubilee",
		Usage: "never again forget an anniversary",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Get the input from `FILE`",
			},
			&cli.IntFlag{
				Name:    "min-age",
				Aliases: []string{"m"},
				Usage:   "Do not consider people younger than `AGE`",
				Value:   60,
			},
			&cli.IntFlag{
				Name:    "year",
				Aliases: []string{"y"},
				Usage:   "Run the processor as if the current year was `YEAR`",
				Value:   time.Now().Year(),
			},
			&cli.IntSliceFlag{
				Name:    "anniversaries",
				Aliases: []string{"a"},
				Usage:   "Define the anniversaries to consider",
				Value:   cli.NewIntSlice(jubilee.TenYearsAnniversary, jubilee.FiveYearsAnniversary),
			},
		},
		Action: makeEntrypoint(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func makeEntrypoint() cli.ActionFunc {
	return func(c *cli.Context) error {
		year := c.Int("year")
		minAge := c.Int("min-age")
		anniversaries := c.IntSlice("anniversaries")

		var input *os.File
		if c.String("file") == "" {
			// We're supposed to run in a pipe
			if !isInputFromPipe() {
				return fmt.Errorf("we are supposed to run from a pipe")
			}

			input = os.Stdin
		} else {
			file, err := os.Open(c.String("file"))
			if err != nil {
				return fmt.Errorf("open file %s: %w", c.String("file"), err)
			}

			input = file
		}

		provider := csv.NewProvider(input)

		records, err := provider.Fetch()
		if err != nil {
			return err
		}

		processor := jubilee.NewProcessor([]jubilee.FilterFunc{
			jubilee.AgeFilter(year, minAge),
			jubilee.AnniversariesFilter(year, anniversaries),
		})

		filtered := processor.Process(records)

		for _, record := range filtered {
			ageToBe := c.Int("year") - record.DateOfBirth.Year()
			fmt.Printf("%s will be %d on %s\n", record.Name, ageToBe, record.DateOfBirth.Format(jubilee.DateFormat))
		}

		return nil
	}
}

func isInputFromPipe() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return fileInfo.Mode()&os.ModeCharDevice == 0
}
