package main

import (
	"flag"
	"fmt"
	"go-core/02-syntax/pkg/crawler"
	"go-core/02-syntax/pkg/crawler/spider"
	"log"
	"strings"
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
	pages := scan(urls)
	for _, p := range pages {
		if strings.Contains(strings.ToLower(p.Title), strings.ToLower(*q)) {
			fmt.Printf("`%s` Найдено на: %s\n", p.Title, p.URL)
		}
	}
}

func scan(urls []string) []crawler.Document {
	var result []crawler.Document
	s := spider.New()
	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
		}
		result = append(result, pages...)
	}
	return result
}
