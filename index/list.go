package index

import (
	client "github.com/f1monkey/elastic-cli/client"
)

func getIndexList() ([]string, error) {
	es, err := client.NewClient()
	if err != nil {
		return nil, err
	}

	names, err := es.IndexNames()
	if err != nil {
		return nil, err
	}

	return names, nil
}
