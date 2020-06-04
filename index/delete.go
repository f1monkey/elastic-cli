package index

import (
	"context"

	elastic "github.com/olivere/elastic/v7"
)

func deleteIndex(client *elastic.Client, name string) error {
	_, err := client.DeleteIndex(name).Do(context.Background())

	return err
}
