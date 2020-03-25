package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/olivere/elastic"
)

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

var (
	indexName = "subject1"
	servers   = []string{"http://localhost:9200/"}
)

func main() {
	subjects := []Subject{
		Subject{
			ID:     1,
			Title:  "肖恩克的救赎",
			Genres: []string{"犯罪", "剧情"},
		},
		Subject{
			ID:     2,
			Title:  "千与千寻",
			Genres: []string{"剧情", "喜剧", "爱情", "战争"},
		},
		Subject{
			ID:     3,
			Title:  "这个杀手太冷",
			Genres: []string{"剧情", "动作", "犯罪"},
		},
		Subject{
			ID:     4,
			Title:  "阿甘正传",
			Genres: []string{"剧情", "爱情"},
		},
	}
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(servers...))
	bulkRequest := client.Bulk()
	for _, subject := range subjects {
		doc := elastic.NewBulkIndexRequest().Index(indexName).Id(strconv.Itoa(subject.ID)).Doc(subject)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(ctx)

	if err != nil {
		panic(err)

	}
	failed := response.Failed()
	l := len(failed)
	if l > 0 {
		fmt.Printf("Error(%d), %v", l, response.Errors)
	}

}
