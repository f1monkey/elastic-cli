package alias

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
)

func deleteAlias(index string, alias string) error {
	es, err := client.NewClient()
	if err != nil {
		return err
	}

	_, err = es.Alias().Remove(index, alias).Do(context.Background())

	return err
}
