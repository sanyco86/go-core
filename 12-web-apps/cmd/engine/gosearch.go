package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go-core/12-web-apps/pkg/crawler"
	"go-core/12-web-apps/pkg/crawler/spider"
	"go-core/12-web-apps/pkg/index/hash"
	"go-core/12-web-apps/pkg/storage"
	"go-core/12-web-apps/pkg/webapp"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

const maxDepth = 3
const fileName = "storage.json"
const address = "0.0.0.0:8080"

var urls = []string{"https://go.dev", "https://golang.org"}

type crawlerDocs struct {
	docs []crawler.Document
}

func main() {
	c := crawlerDocs{}
	log.Print("Scanning pages......")
	c.docs = scan()
	log.Print("Starting application......")
	r := webapp.Router()
	r.HandleFunc("/docs", c.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index/{id}", c.indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/search/{query}", c.searchHandler).Methods(http.MethodGet)
	webapp.ListenAndServe(address, r)
}

func (c *crawlerDocs) docsHandler(w http.ResponseWriter, r *http.Request) {
	if len(c.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var ul string
	for _, doc := range c.docs {
		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (c *crawlerDocs) indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	doc, err := search(id, c.docs)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ul := "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"" + doc.URL + "\">" + doc.Title + "</a></p>"

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
}

func (c *crawlerDocs) searchHandler(w http.ResponseWriter, r *http.Request) {
	if len(c.docs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	store := hash.New()
	store.Add(c.docs)
	vars := mux.Vars(r)
	ids := store.Search(vars["query"])

	if len(ids) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var ul string
	for _, id := range ids {
		doc, err := search(id, c.docs)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ul += "<p>" + fmt.Sprint(doc.ID, ": ") + "<a href=\"/index/" + fmt.Sprint(doc.ID) + "\">" + doc.Title + "</a></p>"
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><div>%v</div></body></html>", ul)
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
