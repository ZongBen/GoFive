package online

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	s        *int
	_ws      *sync.WaitGroup
)

func init() {
	http.HandleFunc("/ws", socketHandler)
}

func startGame(conn *websocket.Conn) {
	b := game.CreateBoard()
	var _gameBoard game.Board = &b
	gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
	for {
		if b.GetTurn() == true {
			input := control.ExecuteCommand(_gameBoard, control.GameCommandHandler)
			conn.WriteMessage(websocket.TextMessage, []byte(string(input)))
		} else {
			_, message, _ := conn.ReadMessage()
			control.GameCommandHandler(_gameBoard, rune(message[0]), 0)
		}
		gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	_ws.Add(1)
	*s = 0

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()
	defer _ws.Done()

	startGame(conn)
}

func StartHostServer(state *int, ch chan int, ws *sync.WaitGroup) {
	s = state
	_ws = ws
	server := http.Server{Addr: ":5555"}
	server.ListenAndServe()
	*state = <-ch
	server.Close()
}
