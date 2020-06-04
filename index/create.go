package index

import (
	"context"

	elastic "github.com/olivere/elastic/v7"
)

func createIndex(client *elastic.Client, name string) error {
	_, err := client.CreateIndex(name).Do(context.Background())

	return err
}
