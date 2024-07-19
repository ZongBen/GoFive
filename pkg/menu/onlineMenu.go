package menu

const (
	JOIN = iota
	HOST
	ONLINE_BACK
)

type OnlineMenu interface {
	GetMenuSelect() int
	SetMenuSelect(int)
}

type onlineMenu struct {
	menuState int
}

func CreateOnlineMenu() onlineMenu {
	return onlineMenu{menuState: JOIN}
}

func (o *onlineMenu) GetMenuSelect() int {
	return o.menuState
}

func (o *onlineMenu) SetMenuSelect(state int) {
	o.menuState = state
}
