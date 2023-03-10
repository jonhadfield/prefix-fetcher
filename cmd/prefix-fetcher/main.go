package main

import (
	"fmt"
	"github.com/jonhadfield/prefix-fetcher/pflog"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var version, versionOutput, tag, sha, buildDate string

func main() {
	pflog.SetLogLevel()

	if tag != "" && buildDate != "" {
		versionOutput = fmt.Sprintf("[%s-%s] %s UTC", tag, sha, buildDate)
	} else {
		versionOutput = version
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Name = "prefix-fetcher"
	app.Version = versionOutput
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "Jon Hadfield",
			Email: "jon@lessknown.co.uk",
		},
	}
	app.HelpName = ""
	app.Description = "prefix-fetcher is a tool to download and display network prefixes from various service providers."
	app.Usage = "prefix-fetcher [global options] provider [command options]"
	app.Flags = []cli.Flag{}
	app.Commands = []*cli.Command{
		abuseipdbCmd(),
		awsCmd(),
		azureCmd(),
		digitaloceanCmd(),
		gcpCmd(),
		geoipCmd(),
	}

	if err := app.Run(os.Args); err != nil {

		fmt.Printf("\nerror: %s\n", err.Error())

	}
}
