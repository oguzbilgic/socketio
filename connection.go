package socketio

import (
	"code.google.com/p/go.net/websocket"
	"time"
)

type Connection struct {
	Session *Session
	Channel string
	Ws      *websocket.Conn
}

func NewConnection(url string, channel string) (*Connection, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	// Connect through websocket
	ws, err := websocket.Dial("ws://"+url+"/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	// Send initial handshake
	if err := websocket.Message.Send(ws, "1::"+channel); err != nil {
		return nil, err
	}

	// Send heartbeats periodically in a seperate goroutine
	go func() {
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			if err := websocket.Message.Send(ws, "2::"); err != nil {
				//return nil, err
			}
		}
	}()

	return &Connection{session, channel, ws}, nil
}

func (conn *Connection) Send(msg string) error {
	return websocket.Message.Send(conn.Ws, msg)
}

func (conn *Connection) Receive(msg *string) error {
	return websocket.Message.Receive(conn.Ws, msg)
}
