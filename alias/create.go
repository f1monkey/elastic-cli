package alias

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
)

func createAlias(index string, alias string) error {
	es, err := client.NewClient()
	if err != nil {
		return err
	}

	_, err = es.Alias().Add(index, alias).Do(context.Background())

	return err
}
