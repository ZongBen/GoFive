package online

import (
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			fmt.Println("Error in receive:", err)
			return
		}
		fmt.Printf("Received: %s\n", msg)
	}
}

func ConnectToHost() {
	connection, _, err := websocket.DefaultDialer.Dial("ws://localhost:5555/ws", nil)
	if err != nil {
		fmt.Println("Error during connection:", err)
		return
	}
	defer connection.Close()

	done = make(chan interface{})
	go receiveHandler(connection)

	for {
		select {
		case <-time.After(1 * time.Second):
			err := connection.WriteMessage(websocket.TextMessage, []byte("Hello from client"))
			if err != nil {
				fmt.Println("Error during message sending:", err)
				return
			}
		case <-done:
			return
		}
	}
}
