package main

import (
	"github.com/ZongBen/GoFive/pkg/display"

	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/game"
)

func main() {
	b := game.CreateBoard()
	display.RenderBoard(&b)
	for !b.Finish {
		key := control.ReadKey()
		control.Command(&b, key)
		display.RenderBoard(&b)
	}
}
