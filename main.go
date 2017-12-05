package main

import (
	"github.com/pipizhang/korkort/korkort"
	"github.com/urfave/cli"
	"os"
)

const (
	APP_NAME    = "KORKORT"
	APP_VERSION = "0.1"
)

func main() {

	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = ""
	app.Version = APP_VERSION

	flagConfig := cli.StringFlag{
		Name:  "config, c",
		Usage: "Load configuratin file from `FILE`",
	}

	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Usage:  "Initialize Application",
			Action: korkort.Setup,
			Flags: []cli.Flag{
				flagConfig,
			},
		},
		{
			Name:   "scrape",
			Usage:  "Scrape data",
			Action: korkort.Scrape,
			Flags: []cli.Flag{
				flagConfig,
			},
		},
	}

	app.Run(os.Args)

}
