package gui

import (
	"strings"

	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/ZongBen/tanvas"
)

var _menuCanvas tanvas.Canvas
var _menuSection tanvas.Section

var _titleCanvas tanvas.Canvas
var _titleSection tanvas.Section

func init() {
	menuCanvas := tanvas.CreateCanvas(22, 4, 1)
	_menuCanvas = &menuCanvas

	menuSection := menuCanvas.CreateSection(0, 0, 22, 4, 0)
	_menuSection = &menuSection

	titleCanvas := tanvas.CreateCanvas(35, 7, 1)
	_titleCanvas = &titleCanvas

	titleSection := titleCanvas.CreateSection(0, 0, 35, 7, 0)
	_titleSection = &titleSection
}

func RenderHome(homeMenu menu.IHomeMenu) string {
	titleOffset_x, title_Offset_y := CenterOffset(35, 15)
	_titleCanvas.SetOffset(titleOffset_x, title_Offset_y)
	result := renderTitle() + renderMenu(homeMenu)
	return result
}

func renderTitle() string {
	title :=
		`
   _____       _____  _           
  / ____|     |  ___|(_)          
 | |  __  ___ | |__   ___   _____ 
 | | |_ |/ _ \|  __| | \ \ / / _ \
 | |__| | (_) | |    | |\ V /  __/
  \_____|\___/|_|    |_| \_/ \___|
  `
	title = strings.Trim(title, "\n")
	lines := strings.Split(title, "\n")
	for i, line := range lines {
		_titleSection.SetRow(0, i, line)
	}
	return _titleCanvas.Render()
}

func renderMenu(m menu.IHomeMenu) string {
	menuOffset_x, _ := CenterOffset(20, 4)
	_menuCanvas.SetOffset(menuOffset_x, 0)
	menu :=
		`
  1. Local  Game
  2. Online Game
  3. Exit       `
	lines := strings.Split(menu, "\n")
	for i, line := range lines {
		if i == m.GetMenuSelect()+1 {
			line += " <=="
		} else {
			line += "    "
		}
		_menuSection.SetRow(0, i, line)
	}
	return _menuCanvas.Render()
}
