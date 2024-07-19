package main

import (
	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/dialog"
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
		gui.Flush(34, 20, gui.RenderHome(_homeMenu), true)
		command := control.ExecuteCommand(_homeMenu, control.HomeMenuCommandHandler)
		switch command {
		case control.LOCAL_GAME:
			StartLocalGame()
		case control.ONLINE_GAME:
		case control.EXIT:
			_homeMenu.Quit()
		}
	}
	gui.Clear()
	gui.Close()
}

func StartLocalGame() {
	var _gameBoard game.Board
	b := game.CreateBoard()
	_gameBoard = &b
	for !_gameBoard.IsFinish() {
		if _gameBoard.GetWinner() != 0 {
			result := showDialog(_gameBoard)
			if result == dialog.AGAIN {
				StartLocalGame()
				break
			} else if result == dialog.QUIT {
				_gameBoard.Quit()
				break
			}
		}
		gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
		control.ExecuteCommand(_gameBoard, control.GameCommandHandler)
	}
	gui.Clear()
}

func showDialog(b game.Board) int {
	state := -1
	for state == -1 {
		gui.Flush(126, 60, gui.RenderBoard(b), true)
		state = control.ExecuteCommand(b.GetDialog(), control.DialogCommandHandler)
	}
	return state
}
