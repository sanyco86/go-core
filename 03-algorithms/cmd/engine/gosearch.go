package main

import (
	"errors"
	"flag"
	"fmt"
	"go-core/pkg/crawler"
	"go-core/pkg/crawler/spider"
	"go-core/pkg/index/hash"
	"log"
	"sort"
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
	store := hash.New()
	docs := scan(urls)
	store.Add(docs)
	ids := store.Search(*q)
	for _, id := range ids {
		doc, err := Search(id, docs)
		if err != nil {
			log.Print(err)
			break
		}
		fmt.Printf("ID: %d Title: %s URL: %s\n", doc.ID, doc.Title, doc.URL)
	}
	fmt.Println("Поиск окончен")
}

func Search(id int, docs []crawler.Document) (crawler.Document, error) {
	index := sort.Search(len(docs), func(index int) bool { return docs[index].ID >= id })
	if index >= len(docs) || docs[index].ID != id {
		doc := crawler.Document{}
		err := errors.New("поиск не дал результатов")
		return doc, err
	}
	return docs[index], nil
}

func scan(urls []string) []crawler.Document {
	var docs []crawler.Document
	var i = 0
	s := spider.New()
	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
			continue
		}
		for _, record := range pages {
			record.ID = i
			docs = append(docs, record)
			i++
		}
	}
	return docs
}
