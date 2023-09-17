package main

import (
	"game-of-life/game"
)

func main() {
	g := game.Init(100)

	g.Start()

	for g.Ticks <= 100 {

	}

	g.Stop()
}
