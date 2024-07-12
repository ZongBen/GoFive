package board

const (
	Empty = iota
	Black
	White
)

type board struct {
	Turn  bool // true: Black, false: White
	Point [64][64]piece
	max_x int
	max_y int
}

type piece struct {
	State int
}

func CreateBoard() board {
	return board{max_x: 64, max_y: 64}
}

func (b *board) SetPiece(x, y int, piece piece) {
	b.Point[x][y] = piece
}

func (b *board) GetPiece() piece {
	state := Empty
	if b.Turn {
		state = Black
	} else {
		state = White
	}
	return piece{state}
}

func (b *board) ChangeTurn() {
	b.Turn = !b.Turn
}

func (b *board) CheckWin() bool {
	for x := 0; x < b.max_x; x++ {
		for y := 0; y < b.max_y; y++ {
			if b.Point[x][y].State == Empty {
				continue
			}
			if checkRight(b, x, y) || checkDown(b, x, y) || checkRightDown(b, x, y) {
				return true
			}
		}
	}
	return false
}

func checkRight(b *board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if x+i > b.max_x {
			return false
		}
		if b.Point[x+i][y].State != state {
			return false
		}
	}
	return true
}

func checkDown(b *board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if y+i > b.max_y {
			return false
		}
		if b.Point[x][y+i].State != state {
			return false
		}
	}
	return true
}

func checkRightDown(b *board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if x+i > b.max_x || y+i > b.max_y {
			return false
		}
		if b.Point[x+i][y+i].State != state {
			return false
		}
	}
	return true
}
