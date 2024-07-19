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
	menuCanvas := tanvas.CreateCanvas(22, 4, 1)
	menuCanvas.SetOffset(7, 0)
	_menuCanvas = &menuCanvas

	menuSection := menuCanvas.CreateSection(0, 0, 22, 4, 0)
	_menuSection = &menuSection

	title = renderTitle()
}

func RenderHome(homeMenu menu.HomeMenu) string {
	_menuSection.Clear()
	result := title + renderHomeMenu(homeMenu)
	return result
}

func RenderOnline(onlineMenu menu.OnlineMenu) string {
	_menuSection.Clear()
	result := title + renderOnlineMenu(onlineMenu)
	return result
}

func renderTitle() string {
	titleCanvas := tanvas.CreateCanvas(35, 7, 1)
	titleSection := titleCanvas.CreateSection(0, 0, 35, 7, 0)
	_titleSection := &titleSection

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
	return titleCanvas.Project()
}

func renderHomeMenu(m menu.HomeMenu) string {
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
	return _menuCanvas.Project()
}

func renderOnlineMenu(m menu.OnlineMenu) string {
	menu :=
		`
  1. Join Game
  2. Host Game
  3. Back     `
	lines := strings.Split(menu, "\n")
	for i, line := range lines {
		if i == m.GetMenuSelect()+1 {
			line += " <=="
		} else {
			line += "    "
		}
		_menuSection.SetRow(0, i, line)
	}
	return _menuCanvas.Project()
}
