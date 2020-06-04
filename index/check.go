package index

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
)

type indexExistsChecker func(name string) (bool, error)

func checkIfIndexExists(name string) (bool, error) {
	es, err := client.NewClient()
	if err != nil {
		return false, err
	}

	result, err := es.IndexExists(name).Do(context.Background())

	return result, err
}
