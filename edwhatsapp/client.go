package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
	writeWait      = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: false,
}

type Message struct {
	Nickname string `json:"nickname,omitempty"`
	Content  string `json:"content,omitempty"`
}
type Client struct {
	nickname     string
	hub          *Hub
	conn         *websocket.Conn
	queueMessage chan Message
}

func (c *Client) writeWS() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, isOpen := <-c.queueMessage:
			if !isOpen {
				return
			}
			c.conn.SetReadDeadline(time.Now().Add(writeWait))

			if err := c.conn.WriteJSON(message); err != nil {
				log.Println("can`t get the websocket connection: ", err)
				return
			}
		}
	}
}

func (c *Client) readWS() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		message := Message{}
		if err := c.conn.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("can`t get the websocket connection: ", err)

			}
			return
		}
		c.hub.broadcast <- message
	}
}

func handleWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query()["nickname"]
	if len(nickname) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("can`t get the websocket connection: ", err)
		return
	}

	client := &Client{
		nickname:     nickname[0],
		hub:          hub,
		conn:         conn,
		queueMessage: make(chan Message, 2),
	}

	client.hub.register <- client
	go client.writeWS()
	go client.readWS()

}
