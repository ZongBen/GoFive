package gui

import "github.com/ZongBen/tanvas"

func renderSelector(s tanvas.Section) {
	s.SetChar(0, 0, '┏')
	s.SetChar(0, 6, '┓')

	s.SetChar(1, 0, '┃')
	s.SetChar(1, 6, '┃')

	s.SetChar(2, 0, '┗')
	s.SetChar(2, 6, '┛')
}
