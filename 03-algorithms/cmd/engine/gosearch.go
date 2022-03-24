package main

import (
	"flag"
	"fmt"
	"go-core/pkg/crawler"
	"go-core/pkg/crawler/spider"
	"go-core/pkg/index/hash"
	"log"
)

const maxDepth = 3

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	q := flag.String("s", "", "Поиск")
	flag.Parse()
	if len(*q) == 0 {
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Начался поиск по: %s...\n", *q)
	indexDB := hash.New()
	records := scan(urls, *indexDB)
	ids := indexDB.Search(*q)
	for _, id := range ids {
		record := records[id]
		fmt.Printf("ID: %d Title: %s URL: %s\n", record.ID, record.Title, record.URL)
	}
}

func scan(urls []string, indexDB hash.Index) []crawler.Document {
	var result []crawler.Document
	var i = 0
	s := spider.New()
	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
		}
		for _, record := range pages {
			record.ID = i
			result = append(result, record)
			i++
		}
	}
	indexDB.Add(result)
	return result
}
