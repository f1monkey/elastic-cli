package main

import (
	"log"
	"os"

	alias "github.com/f1monkey/elastic-cli/alias"
	index "github.com/f1monkey/elastic-cli/index"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "elastic-cli",
		Usage: "Set of CLI tools for ElasticSearch",
		Commands: []*cli.Command{
			alias.GetCommands(),
			index.GetCommands(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
