package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

const (
	gameHeight = 800
	gameWidth  = 1280
)

var op = &ebiten.DrawImageOptions{}

func update(screen *ebiten.Image) error {
	fmt.Println(ebiten.CurrentFPS())

	if ebiten.IsRunningSlowly() {
		return nil
	}

	op.GeoM.Reset()

	return nil
}

func main() {
	ebiten.Run(update, gameWidth, gameHeight, 1, "Ebiten Starter")
}
