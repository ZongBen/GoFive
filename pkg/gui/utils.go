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

func Flush(content string) {
	wg := new(sync.WaitGroup)
	lines := strings.Split(content, "\n")
	for y, line := range lines {
		wg.Add(1)
		go func(y int, line string) {
			for x, char := range line {
				termbox.SetChar(x, y, char)
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
