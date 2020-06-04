package index

import (
	"context"

	elastic "github.com/olivere/elastic/v7"
)

type indexExistsChecker func(client *elastic.Client, name string) (bool, error)

func checkIfIndexExists(client *elastic.Client, name string) (bool, error) {
	result, err := client.IndexExists(name).Do(context.Background())

	return result, err
}
