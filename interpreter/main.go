package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := NewGame()
	if len(os.Args) != 2 {
		log.Fatalln("incorect args, need a .n file to run")
	}

	code, err := ioutil.ReadFile(os.Args[1])
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
