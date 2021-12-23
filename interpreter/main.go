package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	err := game.Interpreter.LoadCode(
		[]byte{0x10, 0x2, 0x10, 0xf, 0x0},
	)
	if err != nil {
		log.Fatalln(err)
	}

	// setup the screen
	ebiten.SetWindowSize(screenWidth*8, screenHeight*8)
	ebiten.SetWindowTitle("VM")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
