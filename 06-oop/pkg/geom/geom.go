package geom

import (
	"go-core/pkg/point"
	"math"
)

// Distance считает и возвращает дистанцию между двумя точками
func Distance(p1, p2 *point.Point) float64 {
	x := p1.X() - p2.X()
	y := p1.Y() - p2.Y()
	return math.Sqrt(x*x + y*y)
}
