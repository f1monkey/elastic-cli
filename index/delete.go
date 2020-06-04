package index

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
)

func deleteIndex(name string) error {
	es, err := client.NewClient()
	if err != nil {
		return err
	}

	_, err = es.DeleteIndex(name).Do(context.Background())

	return err
}
