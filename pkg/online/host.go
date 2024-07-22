package online

import (
	"fmt"
	"net/http"
	"sync"

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
	ch := make(chan rune)
	b := game.CreateBoard()
	var board game.Board = &b
	json := gui.RenderBoard(board)
	for {
		gui.Flush(34, 20, json, true)
		conn.WriteMessage(websocket.TextMessage, []byte(json))
		_, message, _ := conn.ReadMessage()
		ch <- rune(message[0])
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	_ws.Add(1)
	*s = 0
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()
	defer _ws.Done()

	startGame(conn)

	// The event loop
	// for {
	// 	fmt.Println("Waiting for message")
	// 	messageType, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		fmt.Println("Error during message reading:", err)
	// 		break
	// 	}
	// 	fmt.Printf("Received: %s", message)
	// 	err = conn.WriteMessage(messageType, message)
	// 	if err != nil {
	// 		fmt.Println("Error during message writing:", err)
	// 		break
	// 	}
	// }
}

func StartHostServer(state *int, ch chan int, ws *sync.WaitGroup) {
	s = state
	_ws = ws
	server := http.Server{Addr: ":5555"}
	server.ListenAndServe()
	*state = <-ch
	server.Close()
}
