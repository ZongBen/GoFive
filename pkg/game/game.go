package game

const (
	Empty = iota
	Black
	White
)

type Board struct {
	Turn             bool // true: Black, false: White
	Point            [18][18]piece
	Max_x            int
	Max_y            int
	selectorPosition [1][2]int
	Finish           bool
}

type piece struct {
	State int
}

func CreateBoard() Board {
	return Board{Max_x: 18, Max_y: 18, selectorPosition: [1][2]int{{9, 9}}}
}

func (b *Board) GetSelectorPosition() (int, int) {
	return b.selectorPosition[0][0], b.selectorPosition[0][1]
}

func (b *Board) SetSelectorPosition(x, y int) {
	b.selectorPosition[0][0] = x
	b.selectorPosition[0][1] = y
}

func (b *Board) SetPiece(x, y int, piece piece) {
	b.Point[x][y] = piece
}

func (b *Board) Quit() {
	b.Finish = true
}

func (b *Board) GetPiece() piece {
	state := Empty
	if b.Turn {
		state = Black
	} else {
		state = White
	}
	return piece{state}
}

func (b *Board) ChangeTurn() {
	b.Turn = !b.Turn
}

func (b *Board) CheckWin() bool {
	for x := 0; x < b.Max_x; x++ {
		for y := 0; y < b.Max_y; y++ {
			if b.Point[x][y].State == Empty {
				continue
			}
			if checkRight(b, x, y) || checkDown(b, x, y) || checkRightDown(b, x, y) || checkLeftDown(b, x, y) {
				return true
			}
		}
	}
	return false
}

func checkRight(b *Board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if x+i > b.Max_x {
			return false
		}
		if b.Point[x+i][y].State != state {
			return false
		}
	}
	return true
}

func checkDown(b *Board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if y+i > b.Max_y {
			return false
		}
		if b.Point[x][y+i].State != state {
			return false
		}
	}
	return true
}

func checkRightDown(b *Board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if x+i > b.Max_x || y+i > b.Max_y {
			return false
		}
		if b.Point[x+i][y+i].State != state {
			return false
		}
	}
	return true
}

func checkLeftDown(b *Board, x, y int) bool {
	state := b.Point[x][y].State
	for i := 1; i < 5; i++ {
		if x-i < 0 || y+i > b.Max_y {
			return false
		}
		if b.Point[x-i][y+i].State != state {
			return false
		}
	}
	return true
}
