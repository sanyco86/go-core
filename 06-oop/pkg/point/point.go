package point

import "errors"

// Point - Точка координат.
type Point struct {
	x, y float64
}

// New создаёт точку координат и возвращает указатель на нее и ошибку в случае если в конструктор передали отрицательное значение
func New(x, y float64) (*Point, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("координаты не могут быть меньше нуля")
	}
	p := &Point{x, y}
	return p, nil
}

// X возвращает значение x для точки
func (p *Point) X() float64 {
	return p.x
}

// Y возвращает значение y для точки
func (p *Point) Y() float64 {
	return p.y
}
