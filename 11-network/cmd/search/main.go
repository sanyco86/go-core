package main

import (
	"go-core/11-network/pkg/crawler"
	"go-core/11-network/pkg/crawler/spider"
	"go-core/11-network/pkg/netsrv"
	"log"
	"net"
	"sync"
)

const network, address = "tcp4", "0.0.0.0:8000"
const maxDepth = 3

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	var docs []crawler.Document
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		docs = scan(urls, maxDepth)
		wg.Done()
	}()

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if len(docs) == 0 {
			conn.Write([]byte("Введите запрос\n"))
			wg.Wait()
		}
		go netsrv.Handle(conn, docs)
	}
}

func scan(urls []string, depth int) (docs []crawler.Document) {
	scn := spider.New()
	for _, url := range urls {
		result, err := scn.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, result...)
	}
	return docs
}
