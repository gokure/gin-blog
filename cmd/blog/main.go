package main

import (
	"github.com/gokure/gin-blog/internal/models"
	"github.com/gokure/gin-blog/internal/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func init() {
	models.Setup()
}

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"up"},
			Usage:   "Starts web server",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "migrate",
			Usage: "Initializes the database if needed",
			Action: func(c *cli.Context) error {
				return models.AutoMigrate()
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
