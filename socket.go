package socketio

import "time"

type Socket struct {
	Channel   string
	Session   *Session
	Transport *Transport
}

func NewSocket(url string, channel string) (*Socket, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(session, channel)
	if err != nil {
		return nil, err
	}

	socket := &Socket{channel, session, transport}
	socket.handshake()
	go socket.heartbeat()

	return socket, nil
}

func (socket *Socket) Send(msg string) error {
	return socket.Transport.Send(msg)
}

func (socket *Socket) Receive(msg *string) error {
	return socket.Transport.Receive(msg)
}

func (socket *Socket) handshake() error {
	return socket.Transport.Send("1::" + socket.Channel)
}

func (socket *Socket) heartbeat() {
	time.Sleep(time.Duration(socket.Session.HeartbeatTimeout-1) * time.Second)
	if err := socket.Transport.Send("2::"); err == nil {
		socket.heartbeat()
	}
}
