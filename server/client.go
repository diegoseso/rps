package server

import (
	"net/http"
	"log"
	"time"
	"github.com/gorilla/websocket"
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"fmt"
	"github.com/diegoseso/parchis/models"
	_"encoding/json"
	"encoding/json"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
	username = "username"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) reader() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		routeIncomming(message)
		c.hub.broadcast <- message
	}
}

func (c *Client) writer() {
	ticker := time.NewTicker(pingPeriod)
	chatTicker := time.NewTicker(models.OnlinePlayersUpdate)
	defer func() {
		chatTicker.Stop()
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case <-chatTicker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			onlinePLayers, _ := json.Marshal(&models.OnlinePLayersMsg{ Type:"onlinePLayers", Data:&models.Players{Players:models.GetOnlinePlayers()}})
			w.Write(onlinePLayers)
		}

	}
}

func connect(hub *Hub, w http.ResponseWriter, r *http.Request) {

	u := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	   },
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func LoginHandler(w http.ResponseWriter, req *http.Request){
	req.ParseForm()
	spew.Dump(req.FormValue(username))
	if req.FormValue(username) == "" {
		fmt.Fprint(w, "{\"success\"=false, \"data\"=\"No username provided\"}")
		return
	}
	_, player := models.LoginUser(req.FormValue(username))
	fmt.Fprint(w, "{\"success\"=true, \"data\"=\"" + player.Username + "\"}")
}
