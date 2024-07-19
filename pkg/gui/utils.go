package gui

import (
	"strings"
	"sync"

	"github.com/nsf/termbox-go"
)

func init() {
	if !termbox.IsInit {
		termbox.Init()
	}
}

func Flush(offset_x, offset_y int, content string, alignCenter bool) {
	if alignCenter {
		offset_x, offset_y = getOffsetCenter(offset_x, offset_y)
	}
	Clear()
	wg := new(sync.WaitGroup)
	lines := strings.Split(content, "\n")
	for y, line := range lines {
		wg.Add(1)
		go func(y int, line string) {
			for x, char := range line {
				termbox.SetChar(x+offset_x, y+offset_y, char)
			}
			wg.Done()
		}(y, line)
	}
	wg.Wait()
	termbox.Flush()
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func getOffsetCenter(x, y int) (int, int) {
	t_x, t_y := termbox.Size()
	return (t_x - x) / 2, (t_y - y) / 2
}

func Close() {
	termbox.Close()
}
