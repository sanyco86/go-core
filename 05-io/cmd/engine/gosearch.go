package main

import (
	"errors"
	"flag"
	"fmt"
	"go-core/pkg/crawler"
	"go-core/pkg/crawler/spider"
	"go-core/pkg/index/hash"
	"go-core/pkg/storage"
	"log"
	"sort"
)

const maxDepth = 3
const fileName = "storage.json"

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
	docs := scan()
	store.Add(docs)
	ids := store.Search(*q)
	for _, id := range ids {
		doc, err := search(id, docs)
		if err != nil {
			log.Print(err)
			break
		}
		fmt.Printf("ID: %d Title: %s URL: %s\n", doc.ID, doc.Title, doc.URL)
	}
	fmt.Println("Поиск окончен")
}

func search(id int, docs []crawler.Document) (crawler.Document, error) {
	index := sort.Search(len(docs), func(index int) bool { return docs[index].ID >= id })
	if index >= len(docs) || docs[index].ID != id {
		doc := crawler.Document{}
		err := errors.New("поиск не дал результатов")
		return doc, err
	}
	return docs[index], nil
}

func scan() []crawler.Document {
	if storage.Empty(fileName) {
		storage.Write(scanUrls(), fileName)
	}
	docs, _ := storage.Read(fileName)
	return docs
}

func scanUrls() []crawler.Document {
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
