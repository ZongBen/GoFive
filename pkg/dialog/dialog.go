package dialog

const (
	AGAIN = iota
	QUIT
)

type Dialog interface {
	GetState() int
	SetState(int)
}

type dialog struct {
	state int
}

func CreateDialog() dialog {
	return dialog{state: 0}
}

func (d *dialog) GetState() int {
	return d.state
}

func (d *dialog) SetState(state int) {
	d.state = state
}
