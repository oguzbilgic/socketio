package socketio

import (
	"code.google.com/p/go.net/websocket"
)

type Transport interface {
	Send(*IOMessage) error
	Receive() (*IOMessage, error)
}

type WSTransport struct {
	Conn *websocket.Conn
}

func NewWSTransport(session *Session, url, channel string) (*WSTransport, error) {
	// Connect through websocket
	ws, err := websocket.Dial("ws://"+url+"/socket.io/1/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &WSTransport{ws}, nil
}

func (wsTransport *WSTransport) Send(msg *IOMessage) error {
	return websocket.Message.Send(wsTransport.Conn, msg.String())
}

func (wsTransport *WSTransport) Receive() (*IOMessage, error) {
	var rawMsg string
	err := websocket.Message.Receive(wsTransport.Conn, &rawMsg)
	if err != nil {
		return nil, err
	}

	msg, err := NewIOMessage(rawMsg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
