package online

import (
	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
)

func showDialog(b game.Board) int {
	state := -1
	for state == -1 {
		gui.Flush(126, 60, gui.RenderBoard(b), true)
		state = control.ExecuteCommand(b.GetDialog(), control.DialogCommandHandler)
	}
	return state
}
