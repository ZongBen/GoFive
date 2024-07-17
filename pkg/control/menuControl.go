package control

import (
	"github.com/ZongBen/GoFive/pkg/menu"
)

const (
	NOTHING = iota
	LOCAL_GAME
	ONLINE_GAME
	EXIT
)

func HomeMenuCommandHandler(m menu.IHomeMenu, key rune) int {
	selector := m.GetMenuSelect()
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
			return LOCAL_GAME
		case 1:
			return ONLINE_GAME
		case 2:
			return EXIT
		}
	}
	m.SetMenuSelect(selector)
	return 0
}
