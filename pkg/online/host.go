package online

import (
	"fmt"
	"net/http"

	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
	hostCH   chan<- int
)

func init() {
	http.HandleFunc("/ws", socketHandler)
}

func startGame(conn *websocket.Conn) {
	b := game.CreateBoard()
	var board game.Board = &b
	json := gui.RenderBoard(board)
	fmt.Println("Sending: ", json)
	for {
		_ = conn.WriteMessage(websocket.TextMessage, []byte(json))
		_, message, _ := conn.ReadMessage()
		fmt.Println("Received: ", message)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

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

func StartHostServer(ch chan int) {
	hostCH = ch
	server := http.Server{Addr: ":5555"}

	go func() {
		server.ListenAndServe()
	}()

	_, ok := <-ch
	if !ok {
		server.Close()
	}
}
