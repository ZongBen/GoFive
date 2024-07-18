package control

import (
	"github.com/ZongBen/GoFive/pkg/dialog"
)

func DialogCommandHandler(d dialog.IDialog, key rune) int {
	state := d.GetState()
	if key == 'a' || key == 'd' {
		if state == dialog.AGAIN {
			d.SetState(dialog.QUIT)
		} else {
			d.SetState(dialog.AGAIN)
		}
	} else if key == 'e' {
		return state
	}
	return -1
}
