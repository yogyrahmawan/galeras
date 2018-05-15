package websocket

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/yogyrahmawan/galeras/app"

	"log"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *app.MonitorResp, 5)

var upgrader = websocket.Upgrader{}

// Start starting websocket
func Start(publicPath string) {
	fs := http.FileServer(http.Dir(publicPath))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)

	go periodicChecking()
	go handleMessages()

	log.Println("websocket server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	clients[ws] = true

	for {
		// starting all nodes
		log.Println("socket connected")
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println(string(msg))
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func periodicChecking() {
	for {
		firstResp, err := app.MonitorNode("root", "root", "galera-node-1", "SHOW STATUS LIKE 'wsrep_cluster_size';")
		if err != nil {
			log.Printf("first monitoring, err %v", err)
		}

		if firstResp == nil {
			log.Printf("first rep is nil")
			firstResp = &app.MonitorResp{
				Name:  "wsrep_cluster_size",
				Value: "0",
			}
		}

		log.Printf("got message")
		broadcast <- firstResp
		time.Sleep(10 * time.Second)
	}
}
