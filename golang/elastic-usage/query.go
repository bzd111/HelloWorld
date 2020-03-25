package main

import (
	"fmt"

	"github.com/aquasecurity/esquery"
	"github.com/elastic/go-elasticsearch/v7"
)

func MyBoolQuery(es *elasticsearch.Client) {
	qRes, _ := esquery.Query(
		esquery.
			Bool().
			Filter(esquery.Term("title", "王")),
		// Filter(esquery.Term("genres", []string{"冒险"})),
	).Run(
		es,
		es.Search.WithIndex("subject"),
	)
	defer qRes.Body.Close()

	fmt.Printf("Query Body: %v\n", qRes.String())
}
