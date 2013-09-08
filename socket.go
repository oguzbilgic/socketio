// Package socketio implements a client for SocketIO protocol
// as specified in https://github.com/LearnBoost/socket.io-spec
package socketio

import (
	"time"
)

type Socket struct {
	URL       string
	Session   *Session
	Transport Transport
}

func DialAndConnect(url string, channel string, query string) (*Socket, error) {
	socket, err := Dial(url)
	if err != nil {
		return nil, err
	}

	endpoint := NewEndpoint(channel, query)
	connectMsg := NewConnect(endpoint)
	socket.Send(connectMsg)

	return socket, nil
}

func Dial(url string) (*Socket, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(session, url)
	if err != nil {
		return nil, err
	}

	// Heartbeat goroutine
	go func() {
		heartbeatMsg := NewHeartbeat()
		for {
			time.Sleep(session.HeartbeatTimeout - time.Second)
			transport.Send(heartbeatMsg.String())
		}
	}()

	return &Socket{url, session, transport}, nil
}

func (socket *Socket) Receive() (*Message, error) {
	rawMsg, err := socket.Transport.Receive()
	if err != nil {
		return nil, err
	}

	msg, err := ParseMessage(rawMsg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (socket *Socket) Send(msg *Message) error {
	return socket.Transport.Send(msg.String())
}
