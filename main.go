package main

import (
	"fmt"

	"github.com/abradley2/pew-pew-pew-lasors/lib"
	"github.com/hajimehoshi/ebiten"
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
	ebiten.Run(update, lib.GameWidth, lib.GameHeight, 0.25, "Pew Pew Pew")
}
