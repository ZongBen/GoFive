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
	Clear()
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
	termbox.Sync()
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func GetTSize() (int, int) {
	return termbox.Size()
}

func CenterOffset(contentWidth int, contentHeight int) (int, int) {
	screenWidth, screenHeight := GetTSize()
	return max(0, (screenWidth/2)-(contentWidth/2)), max(0, (screenHeight/2)-(contentHeight/2))
}

func Close() {
	termbox.Close()
}
