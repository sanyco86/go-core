package main

import "go-core/pkg/pingpong"

func main() {
	game := pingpong.New()
	game.Start()
}
