package gui

import (
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/tanvas"
)

func renderPiece(piece int, s tanvas.Section) {
	switch piece {
	case game.EMPTY:
		renderEmpty(s)
	case game.BLACK:
		renderBlack(s)
	case game.WHITE:
		renderWhite(s)
	}
}

func renderWhite(s tanvas.Section) {
	s.SetRow(2, 0, "OOO")
	s.SetRow(1, 1, "OOOOO")
	s.SetRow(2, 2, "OOO")
}

func renderBlack(s tanvas.Section) {
	s.SetRow(2, 0, "***")
	s.SetRow(1, 1, "*****")
	s.SetRow(2, 2, "***")
}

func renderEmpty(s tanvas.Section) {
	s.SetRow(0, 0, "       ")
	s.SetRow(0, 1, "   +   ")
	s.SetRow(0, 2, "       ")
}
