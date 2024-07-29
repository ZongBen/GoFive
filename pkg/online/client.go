package online

import (
	"fmt"

	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/dialog"
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
	"github.com/gorilla/websocket"
)

func ConnectToHost(ip string) {
	url := "ws://" + ip + ":5555/ws"
	connection, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Error during connection:", err)
		return
	}
	defer connection.Close()

	startClientGame(connection)
}

func startClientGame(connection *websocket.Conn) {
	b := game.CreateBoard()
	var _gameBoard game.Board = &b
	gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
	for !b.IsFinish() {
		if b.GetWinner() != 0 {
			d := b.GetDialog()
			_, message, _ := connection.ReadMessage()
			dialogResult := control.DialogCommandHandler(d, rune(message[0]), 0)
			if dialogResult == dialog.QUIT {
				b.Quit()
			} else if dialogResult == dialog.AGAIN {
				startClientGame(connection)
			}
		} else {
			if b.GetTurn() == false {
				input := control.ExecuteCommand(_gameBoard, control.GameCommandHandler)
				connection.WriteMessage(websocket.TextMessage, []byte(string(input)))
			} else {
				_, message, _ := connection.ReadMessage()
				control.GameCommandHandler(_gameBoard, rune(message[0]), 0)
			}
		}
		gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
	}
}
