package index

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
)

func createIndex(name string) error {
	es, err := client.NewClient()
	if err != nil {
		return err
	}

	_, err = es.CreateIndex(name).Do(context.Background())

	return err
}
