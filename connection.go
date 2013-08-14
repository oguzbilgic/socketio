package socketio

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"time"
)

type Connection struct {
	Session *Session
	Channel string
	Ws      *websocket.Conn
}

func NewConnection(session *Session, channel string) *Connection {

	// Connect through websocket
	ws, err := websocket.Dial("ws://"+session.Url+"/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}

	// Send initial handshake
	if err := websocket.Message.Send(ws, "1::"+channel); err != nil {
		log.Fatal(err)
	}

	// Send heartbeats periodically in a seperate goroutine
	go func() {
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			if err := websocket.Message.Send(ws, "2::"); err != nil {
				log.Fatal(err)
			}
		}
	}()

	return &Connection{session, channel, ws}
}
