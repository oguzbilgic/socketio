package socketio

import (
	"code.google.com/p/go.net/websocket"
)

type Transport interface {
	send(*Message) error
	receive() (*Message, error)
}

type WSTransport struct {
	Conn *websocket.Conn
}

func NewWSTransport(session *Session, url string) (*WSTransport, error) {
	// Connect through websocket
	ws, err := websocket.Dial("ws://"+url+"/socket.io/1/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &WSTransport{ws}, nil
}

func (wsTransport *WSTransport) send(msg *Message) error {
	return websocket.Message.Send(wsTransport.Conn, msg.String())
}

func (wsTransport *WSTransport) receive() (*Message, error) {
	var rawMsg string
	err := websocket.Message.Receive(wsTransport.Conn, &rawMsg)
	if err != nil {
		return nil, err
	}

	msg, err := ParseMessage(rawMsg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
