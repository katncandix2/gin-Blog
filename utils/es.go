package utils

import (
	"fmt"

	es "github.com/olivere/elastic"
)

func EsClient() *es.Client {
	client, err := es.NewClient(es.SetSniff(false), es.SetURL("http://localhost:9200/"))

	if err != nil {
		fmt.Println("connect es error", err.Error())
		return nil
	}

	return client
}
