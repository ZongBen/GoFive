package control

import (
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/eiannone/keyboard"
)

func ReadKey() rune {
	key, _, _ := keyboard.GetSingleKey()
	return key
}

func Command(b *game.Board, key rune) {
	x, y := b.GetSelectorPosition()
	switch key {
	case 'w':
		if y > 0 {
			y--
		}
	case 's':
		if y < b.Max_y-1 {
			y++
		}
	case 'a':
		if x > 0 {
			x--
		}
	case 'd':
		if x < b.Max_x-1 {
			x++
		}
	case 'e':
		b.SetPiece(x, y, b.GetPiece())
		b.Finish = b.CheckWin()
		b.ChangeTurn()
	case 'q':
		b.Quit()
	}
	b.SetSelectorPosition(x, y)
}
