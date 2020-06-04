package index

import (
	"fmt"

	command "github.com/f1monkey/elastic-cli/command"
	"github.com/urfave/cli/v2"
)

// GetCommands returns list of commands to manipulate ES indexes
func GetCommands() *cli.Command {
	result := cli.Command{
		Name:  "index",
		Usage: "Index manipulation commands",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List existing indexes",
				Action: func(c *cli.Context) error {
					return printIndexList()
				},
			},
			{
				Name:  "exists",
				Usage: "Check if index exists",
				Action: func(c *cli.Context) error {
					indexName, err := command.GetFirstArgument(c, "index name", true)
					if err != nil {
						return err
					}

					result, err := checkIfIndexExists(indexName)
					if err != nil {
						return err
					}

					if result {
						fmt.Printf("Index %q exists\n", indexName)
					} else {
						fmt.Printf("Index %q does not exist\n", indexName)
					}

					return nil
				},
			},
			{
				Name:  "create",
				Usage: "Create index",
				Action: func(c *cli.Context) error {
					indexName, err := command.GetFirstArgument(c, "index name", true)
					if err != nil {
						return err
					}

					err = createIndex(indexName)
					if err != nil {
						return err
					}

					fmt.Printf("Index %q created\n", indexName)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete index",
				Action: func(c *cli.Context) error {
					indexName, err := command.GetFirstArgument(c, "index name", true)
					if err != nil {
						return err
					}

					err = deleteIndex(indexName)
					if err != nil {
						return err
					}

					fmt.Printf("Index %q deleted\n", indexName)
					return nil
				},
			},
		},
	}

	return &result
}

func printIndexList() error {
	names, err := getIndexList()
	if err != nil {
		return err
	}

	fmt.Println("Index list:")
	for _, name := range names {
		fmt.Printf("\t- %s\n", name)
	}

	return nil
}
