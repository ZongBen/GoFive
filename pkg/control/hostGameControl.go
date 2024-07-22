package control

import (
	"github.com/eiannone/keyboard"
)

const (
	CANCEL_HOST = iota
)

type Input struct {
	State *int
	Ch    chan int
}

func HostGameCommandHandler(state Input, _ rune, key keyboard.Key) int {
	switch key {
	case keyboard.KeyEsc:
		*state.State = CANCEL_HOST
		state.Ch <- CANCEL_HOST
		return CANCEL_HOST
	}
	return -1
}
