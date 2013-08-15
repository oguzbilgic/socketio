package socketio

import "time"

type Connection struct {
	Channel   string
	Session   *Session
	Transport *Transport
}

func NewConnection(url string, channel string) (*Connection, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(session, channel)
	if err != nil {
		return nil, err
	}

	conn := &Connection{channel, session, transport}
	conn.handshake()
	go conn.heartbeat()

	return conn, nil
}

func (conn *Connection) Send(msg string) error {
	return conn.Transport.Send(msg)
}

func (conn *Connection) Receive(msg *string) error {
	return conn.Transport.Receive(msg)
}

func (conn *Connection) handshake() error {
	return conn.Transport.Send("1::" + conn.Channel)
}

func (conn *Connection) heartbeat() {
	time.Sleep(time.Duration(conn.Session.HeartbeatTimeout-1) * time.Second)
	if err := conn.Transport.Send("2::"); err == nil {
		conn.heartbeat()
	}
}
