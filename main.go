package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/yourname/mygame/ui"
)

func main() {
	game := ui.NewGame()

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("My Game")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
