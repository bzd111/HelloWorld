package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aquasecurity/esquery"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Failed creating client: %s", err)
	}

	rangeQuery := esquery.Range("publish")
	rangeQuery.Gt("2014-01-01T")
	qRes, _ := esquery.Query(
		rangeQuery,
	).Run(
		es,
		es.Search.WithContext(context.TODO()),
		es.Search.WithIndex("book"),
	)

	defer qRes.Body.Close()
	r, _ := json.Marshal(qRes.String())
	fmt.Printf("result: %v", string(r))
	fmt.Println(qRes.String())
}
