package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// GetFirstArgument returns first argument's value
func GetFirstArgument(c *cli.Context, name string, required bool) (string, error) {
	return GetNthArgument(1, c, name, required)
}

// GetNthArgument returns n-th argument's value
func GetNthArgument(number int, c *cli.Context, name string, required bool) (string, error) {
	argument := c.Args().Get(number - 1)
	if required && argument == "" {
		err := fmt.Errorf("Argument %q must be provided", name)
		return argument, err
	}

	return argument, nil
}
