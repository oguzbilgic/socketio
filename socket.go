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

// Dial opens a new client connection to the socket.io server then connects
// to the given channel.
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

// Dial opens a new client connection to the socket.io server using one of
// the implemented and supported Transports.
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
			err := transport.Send(heartbeatMsg.String())
			if err != nil {
				return
			}
		}
	}()

	return &Socket{url, session, transport}, nil
}

// Receive receives the raw message from the underlying transport and
// converts it to the Message type.
func (socket *Socket) Receive() (*Message, error) {
	rawMsg, err := socket.Transport.Receive()
	if err != nil {
		return nil, err
	}

	return ParseMessage(rawMsg)
}

// Send sends the given Message to the socket.io server using it's
// underlying transport.
func (socket *Socket) Send(msg *Message) error {
	return socket.Transport.Send(msg.String())
}

// Close underlying transport
func (socket *Socket) Close() error {
	return socket.Transport.Close()
}
