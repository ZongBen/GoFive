package menu

const (
	LOCAL_PLAY = iota
	ONLINE_PLAY
	EXIT
)

type HomeMenu interface {
	GetMenuSelect() int
	SetMenuSelect(int)
	Quit() bool
	IsQuit() bool
}

type homeMenu struct {
	menuSelect int
	quit       bool
}

func CreateHomeMenu() homeMenu {
	return homeMenu{menuSelect: LOCAL_PLAY}
}

func (h *homeMenu) GetMenuSelect() int {
	return h.menuSelect
}

func (h *homeMenu) SetMenuSelect(menuSelect int) {
	h.menuSelect = menuSelect
}

func (h *homeMenu) IsQuit() bool {
	return h.quit
}

func (h *homeMenu) Quit() bool {
	h.quit = true
	return h.quit
}
