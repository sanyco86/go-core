package index

// Обратный индекс отсканированных документов.
import "go-core/11-network/pkg/crawler"

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
