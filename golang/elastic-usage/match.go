package main

import (
	"fmt"

	"github.com/aquasecurity/esquery"
	"github.com/elastic/go-elasticsearch/v7"
)

func MyMatchQuery(es *elasticsearch.Client) {

	qRes, _ := esquery.Query(
		esquery.
			Match("title", "人 的"),
	).Run(
		es,
		es.Search.WithIndex("subject"),
	)
	defer qRes.Body.Close()

	fmt.Printf("Match Body: %v\n", qRes.String())

}
