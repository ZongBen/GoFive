package gui

import (
	"sync"

	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/tanvas"
)

var boardCanvas tanvas.Canvas
var dialogSection tanvas.Section

var _instructionCanvas tanvas.Canvas
var _instructionSection tanvas.Section

var offset_x int

func init() {
	c := tanvas.CreateCanvas(18*7, 18*3, 3)
	boardCanvas = &c

	d := c.CreateSection(47, 18, 34, 8, 3)
	d.SetDisplay(false)
	dialogSection = &d

	instructionCanvas := tanvas.CreateCanvas(67, 5, 1)
	_instructionCanvas = &instructionCanvas

	instructionSection := instructionCanvas.CreateSection(0, 0, 67, 5, 1)
	_instructionSection = &instructionSection
}

func renderInstructions(turn bool) string {
	if turn {
		_instructionSection.SetRow(0, 3, "Black's Turn.")
	} else {
		_instructionSection.SetRow(0, 3, "White's Turn.")
	}

	_instructionSection.SetRow(0, 0, "Welcome to GoFive!")
	_instructionSection.SetRow(0, 1, "Use 'w', 'a', 's', 'd' to move the cursor and 'e' to place a piece.")
	_instructionSection.SetRow(0, 2, "Press 'q' to quit.")
	return _instructionCanvas.Project()
}

func RenderBoard(b game.Board) string {
	wg := new(sync.WaitGroup)
	select_x, select_y := b.GetSelectorPosition()

	if b.GetWinner() != game.EMPTY {
		renderDialog(b.GetDialog(), b.GetWinner(), dialogSection)
		dialogSection.SetDisplay(true)
	} else {
		dialogSection.SetDisplay(false)
	}

	for y := 0; y < b.GetHeight(); y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < b.GetWidth(); x++ {
				wg.Add(1)
				go func(x, y int) {
					s := boardCanvas.CreateSection(x*7, y*3, 7, 3, 1)
					s1 := boardCanvas.CreateSection(x*7, y*3, 7, 3, 2)

					renderPiece(b.GetPoint(x, y).State, &s)

					if x == select_x && y == select_y {
						renderSelector(&s1)
					} else {
						s1.Clear()
					}

					wg.Done()
				}(x, y)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
	return renderInstructions(b.GetTurn()) + boardCanvas.Project()
}
