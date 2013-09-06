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
	transport.Send(connectMsg)

	// Heartbeat goroutine
	go func() {
		heartbeatMsg := NewHeartbeat()
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			_ = transport.Send(heartbeatMsg)
		}
	}()

	return &Socket{url, channel, session, transport}, nil
}

func (socket *Socket) Receive() (*IOMessage, error) {
Begining:
	msg, err := socket.Transport.Receive()
	if err != nil {
		return nil, err
	}

	switch msg.Type {
	case 3, 5:
		return msg, nil
	default:
		goto Begining
	}
}
