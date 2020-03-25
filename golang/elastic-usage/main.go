package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/elastic/go-elasticsearch/v7"
	// myquery "elastic-usage/query"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal("Faild creating client: %s", err)
	}
	fmt.Println(reflect.TypeOf(es))

	// bool query
	MyBoolQuery(es)
	// match query
	MyMatchQuery(es)
}
