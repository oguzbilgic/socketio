package socketio

import (
	"code.google.com/p/go.net/websocket"
	"errors"
)

type Transport interface {
	Send(string) error
	Receive() (string, error)
}

func NewTransport(session *Session, url string) (Transport, error) {
	if session.SupportProtocol("websocket") {
		return NewWSTransport(session, url)
	}

	return nil, errors.New("none of the implemented protocols are supported by the server ")
}

type WSTransport struct {
	Conn *websocket.Conn
}

func NewWSTransport(session *Session, url string) (*WSTransport, error) {
	ws, err := websocket.Dial("ws://"+url+"/socket.io/1/websocket/"+session.ID, "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	return &WSTransport{ws}, nil
}

func (wsTransport *WSTransport) Send(rawMsg string) error {
	return websocket.Message.Send(wsTransport.Conn, rawMsg)
}

func (wsTransport *WSTransport) Receive() (string, error) {
	var rawMsg string
	err := websocket.Message.Receive(wsTransport.Conn, &rawMsg)
	if err != nil {
		return "", err
	}

	return rawMsg, nil
}
