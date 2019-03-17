package main

import (
	"context"
	"fmt"
	"reflect"
	"simpleBlog/utils"

	"github.com/olivere/elastic"
)

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

func main() {
	es := utils.EsClient()

	// Use the IndexExists service to check if a specified index exists.
	exists, err := es.IndexExists("twitter").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		mapping := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"doc":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
            "retweets":{
                "type":"long"
            },
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}
`
		createIndex, err := es.CreateIndex("twitter").Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := es.Search().
		Index("twitter").        // search in index "twitter"
		Query(termQuery).        // specify the query
		Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		fmt.Println("err here", err.Error())
		panic(err)
	}

	fmt.Println("success")
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		t := item.(Tweet)
		// fmt.Println(t, t.ArticleId)
		fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	}

	// // searchResult is of type SearchResult and returns hits, suggestions,
	// // and all kinds of other information from Elasticsearch.
	// // fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

}
