package index

import (
	"fmt"

	client "github.com/f1monkey/elastic-cli/client"
	command "github.com/f1monkey/elastic-cli/command"
	elastic "github.com/olivere/elastic/v7"
	"github.com/urfave/cli/v2"
)

// GetIndexCommands returns list of commands to manipulate ES indexes
func GetIndexCommands() *cli.Command {
	result := cli.Command{
		Name:  "index",
		Usage: "Index manipulation commands",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List existing indexes",
				Action: func(c *cli.Context) error {
					client, err := client.NewClient()
					if err != nil {
						return err
					}

					return printIndexList(client)
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

					client, err := client.NewClient()
					if err != nil {
						return err
					}

					result, err := checkIfIndexExists(client, indexName)
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

					client, err := client.NewClient()
					if err != nil {
						return err
					}

					err = createIndex(client, indexName)
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

					client, err := client.NewClient()
					client.Reindex()
					if err != nil {
						return err
					}

					err = deleteIndex(client, indexName)
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

func printIndexList(client *elastic.Client) error {
	names, err := getIndexList(client)
	if err != nil {
		return err
	}

	fmt.Println("Index list:")
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
