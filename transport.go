package socketio

import (
	"code.google.com/p/go.net/websocket"
	"errors"
	"io"
	"time"
)

// Transport is an interface for sending and receiving raw messages from
// the socket.io server.
type transport interface {
	Send(string) error
	Receive() (string, error)
	io.Closer
}

// NewTransport returns an implemented transport which is also supported
// by the socket.io server.
func newTransport(session *Session, url string) (transport, error) {
	if session.SupportProtocol("websocket") {
		return newWsTransport(session, url)
	}

	return nil, errors.New("none of the implemented protocols are supported by the server ")
}

// WSTransport implements Transport interface for websocket protocol.
type wsTransport struct {
	Conn        *websocket.Conn
	readTimeout time.Duration
}

func newWsTransport(session *Session, url string) (*wsTransport, error) {
	urlParser, err := newURLParser(url)
	if err != nil {
		return nil, err
	}
	ws, err := websocket.Dial(urlParser.websocket(session.ID), "", "http://localhost/")
	if err != nil {
		return nil, err
	}
	//expect to receive message once in each HartbeatTimeout
	readTimeout := session.HeartbeatTimeout + time.Second
	return &wsTransport{ws, readTimeout}, nil
}

func (wsTransport *wsTransport) Send(rawMsg string) error {
	return websocket.Message.Send(wsTransport.Conn, rawMsg)
}

func (wsTransport *wsTransport) Receive() (string, error) {
	var rawMsg string
	wsTransport.Conn.SetReadDeadline(time.Now().Add(wsTransport.readTimeout))
	err := websocket.Message.Receive(wsTransport.Conn, &rawMsg)
	if err != nil {
		return "", err
	}

	return rawMsg, nil
}

func (wsTransport *wsTransport) Close() error {
	return wsTransport.Conn.Close()
}
