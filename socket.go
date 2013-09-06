// Package socketio implements a client for SocketIO protocol
// as specified in https://github.com/LearnBoost/socket.io-spec
package socketio

import (
	"time"
)

type Event struct {
	Name string
	Args interface{}
}

type Socket struct {
	Url       string
	Channel   string
	Session   *Session
	Transport Transport
}

func Dial(url string, channel string) (*Socket, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	transport, err := NewWSTransport(session, url, channel)
	if err != nil {
		return nil, err
	}

	// Connect
	endpoint := NewEndpoint(channel, "")
	connectMsg := NewConnect(endpoint)
	transport.send(connectMsg)

	// Heartbeat goroutine
	go func() {
		heartbeatMsg := NewHeartbeat()
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			_ = transport.send(heartbeatMsg)
		}
	}()

	return &Socket{url, channel, session, transport}, nil
}

func (socket *Socket) Receive() (*IOMessage, error) {
	return socket.Transport.receive()
}

func (socket *Socket) Send(msg *IOMessage) error {
	return socket.Transport.send(msg)
}
