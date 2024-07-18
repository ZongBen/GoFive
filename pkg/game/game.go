package game

import "github.com/ZongBen/GoFive/pkg/dialog"

const (
	EMPTY = iota
	BLACK
	WHITE
)

type Board interface {
	GetWidth() int
	GetHeight() int
	GetSelectorPosition() (int, int)
	SetSelectorPosition(x, y int)
	SetPiece(x, y int, piece piece)
	Quit()
	GetPiece() piece
	IsFinish() bool
	GetPoint(x, y int) piece
	ChangeTurn()
	CheckWin()
	GetTurn() bool
	GetWinner() int
	GetDialog() dialog.IDialog
}

type board struct {
	turn             bool // true: Black, false: White
	point            [18][18]piece
	max_x            int
	max_y            int
	selectorPosition [2]int
	finish           bool
	winner           int
	dialog           dialog.IDialog
}

type piece struct {
	State int
}

func (b *board) GetWinner() int {
	return b.winner
}

func (b *board) GetDialog() dialog.IDialog {
	return b.dialog
}

func CreateBoard() board {
	d := dialog.CreateDialog()
	return board{max_x: 18, max_y: 18, selectorPosition: [2]int{9, 9}, dialog: &d}
}

func (b *board) GetTurn() bool {
	return b.turn
}

func (b *board) GetWidth() int {
	return b.max_x
}

func (b *board) GetHeight() int {
	return b.max_y
}

func (b *board) GetSelectorPosition() (int, int) {
	return b.selectorPosition[0], b.selectorPosition[1]
}

func (b *board) SetSelectorPosition(x, y int) {
	b.selectorPosition[0] = x
	b.selectorPosition[1] = y
}

func (b *board) SetPiece(x, y int, piece piece) {
	b.point[y][x] = piece
}

func (b *board) Quit() {
	b.finish = true
}

func (b *board) GetPiece() piece {
	state := EMPTY
	if b.turn {
		state = BLACK
	} else {
		state = WHITE
	}
	return piece{state}
}

func (b *board) IsFinish() bool {
	return b.finish
}

func (b *board) GetPoint(x, y int) piece {
	return b.point[y][x]
}

func (b *board) ChangeTurn() {
	b.turn = !b.turn
}

func (b *board) CheckWin() {
	for y := 0; y < b.max_y; y++ {
		for x := 0; x < b.max_x; x++ {
			if b.point[y][x].State == EMPTY {
				continue
			}
			if checkRight(b, x, y) || checkDown(b, x, y) || checkRightDown(b, x, y) || checkLeftDown(b, x, y) {
				b.winner = b.point[y][x].State
				return
			}
		}
	}
}

func checkRight(b *board, x, y int) bool {
	state := b.point[y][x].State
	for i := 1; i < 5; i++ {
		if x+i > b.max_x {
			return false
		}
		if b.point[y][x+i].State != state {
			return false
		}
	}
	return true
}

func checkDown(b *board, x, y int) bool {
	state := b.point[y][x].State
	for i := 1; i < 5; i++ {
		if y+i > b.max_y {
			return false
		}
		if b.point[y+i][x].State != state {
			return false
		}
	}
	return true
}

func checkRightDown(b *board, x, y int) bool {
	state := b.point[y][x].State
	for i := 1; i < 5; i++ {
		if x+i > b.max_x || y+i > b.max_y {
			return false
		}
		if b.point[y+i][x+i].State != state {
			return false
		}
	}
	return true
}

func checkLeftDown(b *board, x, y int) bool {
	state := b.point[y][x].State
	for i := 1; i < 5; i++ {
		if x-i < 0 || y+i > b.max_y {
			return false
		}
		if b.point[y+i][x-i].State != state {
			return false
		}
	}
	return true
}
