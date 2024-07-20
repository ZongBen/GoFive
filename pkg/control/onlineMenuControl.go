package control

import (
	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/eiannone/keyboard"
)

func OnlineMenuCommandHandler(o menu.OnlineMenu, char rune, key keyboard.Key) int {
	selector := o.GetMenuSelect()
	switch char {
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
