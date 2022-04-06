package main

import (
	"fmt"
	"go-core/pkg/geom"
	"go-core/pkg/point"
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
