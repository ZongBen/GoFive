package control

import (
	"github.com/eiannone/keyboard"
)

const (
	CANCEL_HOST = iota
)

func HostGameCommandHandler(ch chan int, _ rune, key keyboard.Key) int {
	_, ok := <-ch
	switch key {
	case keyboard.KeyEsc:
		if ok {
			close(ch)
		}
		return CANCEL_HOST
	}
	if !ok {
		return 0
	}
	return -1
}
