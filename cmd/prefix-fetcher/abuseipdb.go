package main

import (
	"github.com/jonhadfield/prefix-fetcher/abuseipdb"
	"github.com/urfave/cli/v2"
)

func abuseipdbCmd() *cli.Command {
	return &cli.Command{
		Name:  "abuseipdb",
		Usage: "fetch abuseipdb prefixes",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key",
				Usage: "api key", Aliases: []string{"k"}, Required: true,
			},
			&cli.IntFlag{
				Name:  "confidence",
				Usage: "minimum confidence percentage score to return", Value: 75, Aliases: []string{"c"},
			},
			&cli.Int64Flag{
				Name:  "limit",
				Usage: "maximum number of results to return", Value: 1000, Aliases: []string{"l"},
			},
			&cli.StringFlag{
				Name:  "path",
				Usage: "where to save the file", Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "output as yaml or json", Aliases: []string{"f"},
			},
		},
		Action: func(c *cli.Context) error {
			a := abuseipdb.New()
			a.Limit = c.Int64("limit")
			a.APIKey = c.String("key")
			a.ConfidenceMinimum = c.Int("confidence")
			data, _, _, err := a.FetchData()
			if err != nil {
				return err
			}

			path := c.String("path")
			if path != "" {
				return saveFile(saveFileInput{
					provider:        "abuseipdb",
					data:            data,
					path:            path,
					defaultFileName: "blacklist",
				})
			}

			return nil
		},
	}
}
