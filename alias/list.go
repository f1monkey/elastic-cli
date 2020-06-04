package alias

import (
	"context"

	client "github.com/f1monkey/elastic-cli/client"
	elastic "github.com/olivere/elastic/v7"
)

func getAliasList() (*elastic.AliasesResult, error) {
	es, err := client.NewClient()
	if err != nil {
		return nil, err
	}

	return es.Aliases().Do(context.Background())
}
