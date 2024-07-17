package gui

import (
	"strings"

	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/ZongBen/tanvas"
)

var _menuCanvas tanvas.Canvas
var _menuSection tanvas.Section
var title string

func init() {
	title = renderTitle()
	menuCanvas := tanvas.CreateCanvas(44, 4, 1)
	_menuCanvas = &menuCanvas

	menuSection := menuCanvas.CreateSection(8, 0, 36, 4, 0)
	_menuSection = &menuSection
}

func RenderHome(homeMenu menu.IHomeMenu) string {
	title := renderTitle()
	result := title + renderMenu(homeMenu)
	return result
}

func renderTitle() string {
	titleCanvas := tanvas.CreateCanvas(40, 7, 1)
	titleSection := titleCanvas.CreateSection(0, 0, 40, 7, 0)
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
		titleSection.SetRow(6, i, line)
	}
	return titleCanvas.Render()
}

func renderMenu(m menu.IHomeMenu) string {
	_menuSection.Clear()
	menu :=
		`
  1. Local  Game
  2. Online Game
  3. Exit       `
	lines := strings.Split(menu, "\n")
	for i, line := range lines {
		if i == m.GetMenuSelect()+1 {
			line += " <=="
		}
		_menuSection.SetRow(6, i, line)
	}
	return _menuCanvas.Render()
}
