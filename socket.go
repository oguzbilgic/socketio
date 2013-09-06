// Package socketio implements a client for SocketIO protocol
// as specified in https://github.com/LearnBoost/socket.io-spec
package socketio

import (
	"encoding/json"
	"strings"
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
		var msg string
		socket.Receive(&msg)

		msgCode := string(msg[0])

		switch msgCode {
		case "3":
			msgParts := strings.SplitN(msg, ":", 4)
			msgData := msgParts[3]

			for _, onMessageFunc := range socket.OnMessageFuncs {
				go onMessageFunc(msgData)
			}
		case "5":
			msgParts := strings.SplitN(msg, ":", 4)
			msgData := msgParts[3]

			var event Event
			err := json.Unmarshal([]byte(msgData), &event)
			if err != nil {
				panic(err)
			}

			for _, onEventFunc := range socket.OnEventFuncs {
				go onEventFunc(event.Name, event.Args)
			}
		default:
			println(msg)
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

func (socket *Socket) Receive(msg *string) error {
	return socket.Transport.Receive(msg)
}
