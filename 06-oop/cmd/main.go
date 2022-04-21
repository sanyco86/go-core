package main

import (
	"fmt"
	"go-core/06-oop/pkg/geom"
	"go-core/06-oop/pkg/point"
)

func main() {
	p1, err := point.New(1., 2.)
	if err != nil {
		fmt.Println(err)
		return
	}
	p2, err := point.New(4., 6.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(geom.Distance(p1, p2))
}
