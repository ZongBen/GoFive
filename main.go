package main

import (
	"fmt"

	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
	"github.com/ZongBen/GoFive/pkg/menu"
)

var _homeMenu menu.IHomeMenu

func init() {
	homeMenu := menu.CreateHomeMenu()
	_homeMenu = &homeMenu
}

func main() {

	for !_homeMenu.IsQuit() {
		gui.Clear()
		fmt.Println(gui.RenderHome(_homeMenu))
		command := control.ExecuteCommand(_homeMenu, control.HomeMenuCommandHandler)
		switch command {
		case control.LOCAL_GAME:
			StartLocalGame()
		case control.ONLINE_GAME:
			fmt.Println("Online Game")
		case control.EXIT:
			_homeMenu.Quit()
		}
	}
}

func StartLocalGame() {
	var _gameBoard game.Board
	b := game.CreateBoard()
	_gameBoard = &b
	for !b.IsFinish() {
		gui.Clear()
		gui.RenderBoard(_gameBoard)
		control.ExecuteCommand(_gameBoard, control.GameCommandHandler)
	}
}
