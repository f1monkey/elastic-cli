package index

type indexNamesGetter interface {
	IndexNames() ([]string, error)
}

func getIndexList(client indexNamesGetter) ([]string, error) {
	names, err := client.IndexNames()
	if err != nil {
		return nil, err
	}

	return names, nil
}
