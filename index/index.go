package index

import (
	"context"

	elastic "github.com/olivere/elastic/v7"
)

type indexNamesGetter interface {
	IndexNames() ([]string, error)
}

type indexExistsChecker func(client *elastic.Client, name string) (bool, error)

func getIndexList(client indexNamesGetter) ([]string, error) {
	names, err := client.IndexNames()
	if err != nil {
		return nil, err
	}

	return names, nil
}

func checkIfIndexExists(client *elastic.Client, name string) (bool, error) {
	result, err := client.IndexExists(name).Do(context.Background())

	return result, err
}

func createIndex(client *elastic.Client, name string) error {
	_, err := client.CreateIndex(name).Do(context.Background())

	return err
}

func deleteIndex(client *elastic.Client, name string) error {
	_, err := client.DeleteIndex(name).Do(context.Background())

	return err
}
