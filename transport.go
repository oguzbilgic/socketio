package socketio

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"time"
)

// Transport is an interface for sending and receiving raw messages from
// the socket.io server.
type Transport interface {
	Send(string) error
	Receive() (string, error)
}

// NewTransport returns an implemented transport which is also supported
// by the socket.io server.
func NewTransport(session *Session, url string) (Transport, error) {
	if session.SupportProtocol("websocket") {
		return NewWSTransport(session, url)
	}

	return nil, errors.New("none of the implemented protocols are supported by the server ")
}

// WSTransport implements Transport interface for websocket protocol.
type WSTransport struct {
	Conn *websocket.Conn
	readTimeout time.Duration
}

func NewWSTransport(session *Session, url string) (*WSTransport, error) {
	urlParser, err := NewUrlParser(url)
	if err != nil {
		return nil, err
	}
	ws, err := websocket.Dial(urlParser.Websocket(session.ID), "", "http://localhost/")
	if err != nil {
		return nil, err
	}
	//expect to receive message once in each HartbeatTimeout 
	readTimeout := session.HeartbeatTimeout + time.Second
	return &WSTransport{ws, readTimeout}, nil
}

func (wsTransport *WSTransport) Send(rawMsg string) error {
	return websocket.Message.Send(wsTransport.Conn, rawMsg)
}

func (wsTransport *WSTransport) Receive() (string, error) {
	var rawMsg string
	wsTransport.Conn.SetReadDeadline(time.Now().Add(wsTransport.readTimeout))
	err := websocket.Message.Receive(wsTransport.Conn, &rawMsg)
	if err != nil {
		return "", err
	}

	return rawMsg, nil
}
