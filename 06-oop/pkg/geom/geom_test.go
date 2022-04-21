package geom

import (
	"go-core/06-oop/pkg/point"
	"testing"
)

func TestDistance(t *testing.T) {
	p1, _ := point.New(7., 5.)
	p2, _ := point.New(4., 1.)
	got := Distance(p1, p2)
	want := 5.
	if got != want {
		t.Fatalf("получили %+v, ожидалось %+v", got, want)
	}
}
