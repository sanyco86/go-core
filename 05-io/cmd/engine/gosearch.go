package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go-core/05-io/pkg/crawler"
	"go-core/05-io/pkg/crawler/spider"
	"go-core/05-io/pkg/index/hash"
	"go-core/05-io/pkg/storage"
	"log"
	"os"
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
	if empty(fileName) {
		store(scanUrls(), fileName)
	}
	docs, _ := get(fileName)
	return docs
}

func store(docs []crawler.Document, fileName string) (bool, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return false, err
	}
	data, err := json.Marshal(docs)
	if err != nil {
		return false, err
	}
	defer f.Close()
	err = storage.Write(f, data)
	if err != nil {
		return false, err
	}
	return true, err
}

func get(fileName string) ([]crawler.Document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	b, err := storage.Read(f)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var docs []crawler.Document
	json.Unmarshal([]byte(b), &docs)
	return docs, nil
}

func empty(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return true
	}
	fileInfo, err := os.Lstat(fileName)
	if fileInfo.Size() == 0 {
		return true
	}
	if err != nil {
		return true
	}
	return false
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
