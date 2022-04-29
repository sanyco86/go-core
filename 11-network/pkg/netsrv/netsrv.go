package netsrv

import (
	"bufio"
	"go-core/11-network/pkg/crawler"
	"log"
	"net"
	"strings"
)

func Handle(conn net.Conn, docs []crawler.Document) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		conn.Write([]byte("\nВведите запрос: "))
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		if len(msg) > 0 {
			conn.Write([]byte("Результат поиска:\n"))
			for _, doc := range docs {
				if strings.Contains(strings.ToLower(doc.Title), strings.ToLower(string(msg))) {
					result := doc.URL + doc.Title + "\n"
					log.Print(result)
					conn.Write([]byte(result))
				}
			}
		} else {
			conn.Write([]byte("Ничего не найдено!\n"))
			return
		}
	}
}
