package socketio

import (
	"code.google.com/p/go.net/websocket"
)

type Transport struct {
	Conn *websocket.Conn
}

func NewTransport(session *Session, url, channel string) (*Transport, error) {
	// Connect through websocket
	ws, err := websocket.Dial("ws://"+url+"/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &Transport{ws}, nil
}

func (transport *Transport) Send(msg string) error {
	return websocket.Message.Send(transport.Conn, msg)
}

func (transport *Transport) Receive() (*IOMessage, error) {
	var rawMsg string
	websocket.Message.Receive(transport.Conn, &rawMsg)

	msg, err := NewIOMessage(rawMsg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
