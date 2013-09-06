// Package socketio implements a client for SocketIO protocol
// as specified in https://github.com/LearnBoost/socket.io-spec
package socketio

import (
	"encoding/json"
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
	Transport *Transport

	OnMessageFuncs []OnMessageFunc
	OnEventFuncs   []OnEventFunc
	OnJSONFuncs    []OnJSONFunc
}

type OnDisconnectFunc func(channel string)

type OnConnectFunc func(endpoint Endpoint)

type OnHeartbeatFunc func()

type OnMessageFunc func(msg string)

type OnEventFunc func(event string, args interface{})

type OnJSONFunc func(jsonStr string)

type OnACKFunc func(id string, data string)

type OnErrorFunc func(channel, reason, advice string)

type OnNoopFunc func()

func NewSocket(url string, channel string) (*Socket, error) {
	session, err := NewSession(url)
	if err != nil {
		return nil, err
	}

	transport, err := NewTransport(session, url, channel)
	if err != nil {
		return nil, err
	}

	return &Socket{url, channel, session, transport, nil, nil, nil}, nil
}

func (socket *Socket) Dial() {
	// Handshake
	socket.Transport.Send("1::" + socket.Channel)

	// Heartbeat goroutine
	go func() {
		for {
			time.Sleep(time.Duration(socket.Session.HeartbeatTimeout-1) * time.Second)
			_ = socket.Transport.Send("2::")
		}
	}()

	// Message receiving loop
	for {
		msg, err := socket.Transport.Receive()
		if err != nil {
			println(err.Error())
			continue
		}

		switch msg.Type {
		case 3:
			for _, onMessageFunc := range socket.OnMessageFuncs {
				go onMessageFunc(msg.Data)
			}
		case 5:
			var event Event
			err := json.Unmarshal([]byte(msg.Data), &event)
			if err != nil {
				panic(err)
			}

			for _, onEventFunc := range socket.OnEventFuncs {
				go onEventFunc(event.Name, event.Args)
			}
		default:
			println(msg.Type)
		}
	}
}

func (socket *Socket) OnMessage(omf OnMessageFunc) {
	socket.OnMessageFuncs = append(socket.OnMessageFuncs, omf)
}

func (socket *Socket) OnJSON(ojf OnJSONFunc) {
	socket.OnJSONFuncs = append(socket.OnJSONFuncs, ojf)
}

func (socket *Socket) OnEvent(oef OnEventFunc) {
	socket.OnEventFuncs = append(socket.OnEventFuncs, oef)
}

func (socket *Socket) Send(msg string) error {
	return socket.Transport.Send(msg)
}
