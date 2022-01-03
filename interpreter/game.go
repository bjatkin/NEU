package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 64
	screenHeight = 64
)

type Game struct {
	Buff        image.Image
	Interpreter *Interp
}

func NewGame() *Game {
	game := &Game{
		Interpreter: &Interp{},
	}

	game.Buff = &ScreenBuff{
		Ref: game.Interpreter.Memory[:1024],
	}

	game.Interpreter.Debugger = NewDebugger(game.Interpreter)

	return game
}

func (g *Game) Update() error {
	g.Interpreter.Debugger.frameCount++
	return g.Interpreter.Run()
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw := ebiten.NewImageFromImage(g.Buff)
	screen.DrawImage(draw, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

type ScreenBuff struct {
	Ref []byte
}

func (s *ScreenBuff) ColorModel() color.Model {
	return color.RGBAModel
}

func (s *ScreenBuff) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 64, Y: 64},
	}
}

func (s *ScreenBuff) At(x, y int) color.Color {
	i := y*64 + x
	frac := i % 4
	offset := i / 4
	b := s.Ref[offset]
	var pal byte
	switch frac {
	case 0:
		pal = b & 0b00000011
	case 1:
		pal = (b >> 2) & 0b00000011
	case 2:
		pal = (b >> 4) & 0b00000011
	case 3:
		pal = (b >> 6) & 0b00000011
	}
	return BuffColor{Pal: int(pal)}
}

type BuffColor struct {
	Pal int
}

func (c BuffColor) RGBA() (r, g, b, a uint32) {
	switch c.Pal {
	case 0:
		return 239 * 255, 249 * 255, 214 * 255, 0xffff
	case 1:
		return 186 * 255, 80 * 255, 68 * 255, 0xffff
	case 2:
		return 122 * 255, 28 * 255, 75 * 255, 0xffff
	case 3:
		return 27 * 255, 3 * 255, 38 * 255, 0xffff
	default:
		fmt.Printf("refusing to handle color index %d(%b)\n", c.Pal, c.Pal)
		panic(1)
	}
}
