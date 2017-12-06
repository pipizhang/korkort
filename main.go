package main

import (
	"github.com/pipizhang/korkort/korkort"
	"github.com/urfave/cli"
	"os"
)

const (
	APP_NAME    = "KORKORT"
	APP_VERSION = "1.0"
)

func main() {

	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = ""
	app.Version = APP_VERSION

	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Usage:  "Initialize database",
			Action: korkort.Setup,
		},
		{
			Name:   "scrape",
			Usage:  "Scrape data",
			Action: korkort.Scrape,
		},
	}

	app.Run(os.Args)

}
