package main

import "go-core/10-concurrency/pkg/pingpong"

func main() {
	game := pingpong.New()
	game.Start()
}
