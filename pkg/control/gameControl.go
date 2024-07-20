package control

import (
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/eiannone/keyboard"
)

func GameCommandHandler(b game.Board, char rune, key keyboard.Key) int {
	x, y := b.GetSelectorPosition()
	switch char {
	case 'w':
		if y > 0 {
			y--
		}
	case 's':
		if y < b.GetHeight()-1 {
			y++
		}
	case 'a':
		if x > 0 {
			x--
		}
	case 'd':
		if x < b.GetWidth()-1 {
			x++
		}
	case 'e':
		if b.GetPoint(x, y).State != game.EMPTY {
			break
		}
		b.SetPiece(x, y, b.GetPiece())
		b.CheckWin()
		b.ChangeTurn()
	case 'q':
		b.Quit()
	}
	b.SetSelectorPosition(x, y)
	return 0
}
