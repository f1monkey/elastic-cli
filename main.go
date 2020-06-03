package main

import (
	"log"
	"os"

	index "github.com/f1monkey/elastic-cli/index"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "elastic-cli",
		Usage: "Set of CLI tools for ElasticSearch",
		Commands: []*cli.Command{
			index.GetIndexCommands(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
