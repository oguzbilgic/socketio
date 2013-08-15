package socketio

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

	// Send initial handshake
	if err := transport.Send("1::" + channel); err != nil {
		return nil, err
	}

	// Send heartbeats periodically in a seperate goroutine
	go func() {
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			if err := transport.Send("2::"); err != nil {
				break
			}
		}
	}()

	return &Connection{channel, session, transport}, nil
}

func (conn *Connection) Send(msg string) error {
	return conn.Transport.Send(msg)
}

func (conn *Connection) Receive(msg *string) error {
	return conn.Transport.Receive(msg)
}
