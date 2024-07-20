package control

import (
	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/eiannone/keyboard"
)

func HomeMenuCommandHandler(m menu.HomeMenu, char rune, key keyboard.Key) int {
	selector := m.GetMenuSelect()
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
			return menu.LOCAL_PLAY
		case 1:
			return menu.ONLINE_PLAY
		case 2:
			return menu.EXIT
		}
	}
	m.SetMenuSelect(selector)
	return -1
}
