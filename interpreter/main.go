package main

import (
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()

	code, err := ioutil.ReadFile("hello_world.n")
	if err != nil {
		log.Fatalln(err)
	}

	err = game.Interpreter.LoadCode(code)
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
