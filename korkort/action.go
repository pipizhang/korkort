package korkort

import (
	"fmt"
	"github.com/urfave/cli"
)

func Setup(c *cli.Context) {
	InitConfig()

	DB := GetDB()
	defer DB.Close()

	DB.AutoMigrate(&Question{}, &Choice{})
}

func Scrape(c *cli.Context) {
	fmt.Println("scrape .....")
}
