package alias

import (
	"fmt"

	"github.com/f1monkey/elastic-cli/command"
	"github.com/urfave/cli/v2"
)

// GetCommands returns list of commands to manipulate ES index aliases
func GetCommands() *cli.Command {
	result := cli.Command{
		Name:  "alias",
		Usage: "Index alias manipulation commands",
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List existing aliases by index",
				Action: func(c *cli.Context) error {
					return printAllIndexAliasList()
				},
			},
			{
				Name:  "create",
				Usage: "Create alias",
				Action: func(c *cli.Context) error {
					indexName, err := command.GetFirstArgument(c, "index name", true)
					if err != nil {
						return err
					}

					aliasName, err := command.GetNthArgument(2, c, "alias name", true)
					if err != nil {
						return err
					}

					err = createAlias(indexName, aliasName)
					if err != nil {
						return err
					}

					fmt.Printf("Index %q alias %q created\n", indexName, aliasName)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Delete alias",
				Action: func(c *cli.Context) error {
					indexName, err := command.GetFirstArgument(c, "index name", true)
					if err != nil {
						return err
					}

					aliasName, err := command.GetNthArgument(2, c, "alias name", true)
					if err != nil {
						return err
					}
					err = deleteAlias(indexName, aliasName)

					if err != nil {
						return err
					}

					fmt.Printf("Index %q alias %q deleted\n", indexName, aliasName)
					return nil
				},
			},
		},
	}

	return &result
}

func printAllIndexAliasList() error {
	names, err := getAliasList()
	if err != nil {
		return err
	}

	fmt.Println("Index alias list:")
	for indexName, aliases := range names.Indices {
		fmt.Printf("\tIndex %s:\n", indexName)
		for _, alias := range aliases.Aliases {
			fmt.Printf("\t\t - name: %s, is_write_index: %t\n", alias.AliasName, alias.IsWriteIndex)
		}
	}

	return nil
}
