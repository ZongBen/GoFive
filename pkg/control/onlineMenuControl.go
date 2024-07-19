package control

import (
	"github.com/ZongBen/GoFive/pkg/menu"
)

func OnlineMenuCommandHandler(o menu.OnlineMenu, key rune) int {
	selector := o.GetMenuSelect()
	switch key {
	case 'w':
		if selector > 0 {
			selector--
		} else {
			selector = 2
		}
	case 's':
		if selector < 2 {
			selector++
		} else {
			selector = 0
		}
	case 'e':
		switch selector {
		case 0:
			return menu.JOIN
		case 1:
			return menu.HOST
		case 2:
			return menu.ONLINE_BACK
		}
	}
	o.SetMenuSelect(selector)
	return -1
}
