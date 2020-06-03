package client

import (
	"fmt"
	"os"

	elastic "github.com/olivere/elastic/v7"
)

const esURLVariable string = "ELASTICSEARCH_URL"

// NewClient Create elasticsearch client instance
func NewClient() (*elastic.Client, error) {
	elasticsearchURL, exists := os.LookupEnv(esURLVariable)
	if !exists {
		err := fmt.Errorf("Env variable %q must be defined", esURLVariable)

		return nil, err
	}

	return elastic.NewClient(elastic.SetURL(elasticsearchURL))
}
