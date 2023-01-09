package main

import (
	"fmt"
	"github.com/jonhadfield/prefix-fetcher/maxmind/geoip"
	"github.com/urfave/cli/v2"
)

func geoipCmd() *cli.Command {
	return &cli.Command{
		Name:  "geoip",
		Usage: "fetch maxmind geoip databases",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key",
				Usage: "license key", Aliases: []string{"k"}, Required: true,
			},
			&cli.StringFlag{
				Name:  "path",
				Usage: "where to save the file", Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "database format csv or mmdb", Value: "csv", Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:  "edition",
				Usage: "GeoLite2 or GeoIP2", Value: "GeoLite2", Aliases: []string{"e"},
			},
			&cli.BoolFlag{
				Name:  "extract",
				Usage: "extract compressed downloads", Value: true,
			},
		},
		Action: func(c *cli.Context) error {
			a := geoip.New()
			a.LicenseKey = c.String("key")
			a.Edition = c.String("edition")
			a.DBFormat = c.String("format")
			a.Root = c.String("path")
			a.Extract = c.Bool("extract")
			out, err := a.FetchFiles()
			if err != nil {
				return err
			}

			if a.Extract {
				fmt.Println("geoip data written to:")
				fmt.Printf("country - %s\n", out.CountryDataPath)
				fmt.Printf("city - %s\n", out.CityDataPath)
				fmt.Printf("asn - %s\n", out.ASNDataPath)

				return nil
			}

			fmt.Println("geoip compressed files written to:")
			fmt.Printf("country - %s\n", out.CountryCompressed)
			fmt.Printf("city - %s\n", out.CityCompressed)
			fmt.Printf("asn - %s\n", out.ASNCompressed)

			return nil
		},
	}
}