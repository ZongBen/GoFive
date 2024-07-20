package control

import (
	"github.com/ZongBen/GoFive/pkg/dialog"
	"github.com/eiannone/keyboard"
)

func DialogCommandHandler(d dialog.Dialog, char rune, key keyboard.Key) int {
	state := d.GetState()
	if char == 'a' || char == 'd' {
		if state == dialog.AGAIN {
			d.SetState(dialog.QUIT)
		} else {
			d.SetState(dialog.AGAIN)
		}
	} else if char == 'e' {
		return state
	}
	return -1
}
