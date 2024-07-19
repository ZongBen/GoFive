package gui

import (
	"github.com/ZongBen/GoFive/pkg/dialog"
	"github.com/ZongBen/tanvas"
)

func renderDialog(d dialog.Dialog, winner int, s tanvas.Section) {
	turn := ""
	if winner == 1 {
		turn = "Black"
	} else {
		turn = "White"
	}

	if d.GetState() == dialog.AGAIN {
		s.SetRow(0, 4, "|    ---------                   |")
		s.SetRow(0, 5, "|    | Again |        Quit       |")
		s.SetRow(0, 6, "|    ---------                   |")
	} else if d.GetState() == dialog.QUIT {
		s.SetRow(0, 4, "|                   --------     |")
		s.SetRow(0, 5, "|      Again        | Quit |     |")
		s.SetRow(0, 6, "|                   --------     |")
	}

	s.SetRow(0, 0, "==================================")
	s.SetRow(0, 1, "|                                |")
	s.SetRow(0, 2, "|           "+turn+" Win            |")
	s.SetRow(0, 3, "|                                |")
	s.SetRow(0, 7, "==================================")
}
