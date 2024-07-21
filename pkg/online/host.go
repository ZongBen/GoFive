package online

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		fmt.Println("Waiting for message")
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error during message reading:", err)
			break
		}
		fmt.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error during message writing:", err)
			break
		}
	}
}

func StartWebsocketServer() {
	http.HandleFunc("/ws", socketHandler)
	fmt.Println("Server started at port 5555")
	http.ListenAndServe(":5555", nil)
}
