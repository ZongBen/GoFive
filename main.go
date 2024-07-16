package main

import (
	// "github.com/ZongBen/GoFive/pkg/gui"
	"fmt"

	"github.com/ZongBen/GoFive/pkg/gui/scenes"
	// "github.com/ZongBen/GoFive/pkg/control"
	// "github.com/ZongBen/GoFive/pkg/game"
)

func main() {
	// b := game.CreateBoard()
	// gui.RenderBoard(&b)
	// for !b.Finish {
	// 	key := control.ReadKey()
	// 	control.Command(&b, key)
	// 	gui.RenderBoard(&b)
	// }

	fmt.Println(scenes.RenderHome())
}
