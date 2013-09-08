package socketio

import (
	"errors"
	"strconv"
	"strings"
)

type Message struct {
	Type     int
	ID       string
	Endpoint *Endpoint
	Data     string
}

func ParseMessage(rawMsg string) (*Message, error) {
	parts := strings.SplitN(rawMsg, ":", 4)

	if len(parts) < 3 {
		return nil, errors.New("Empty message")
	}

	msgType, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	id := parts[1]
	endpoint := ParseEndpoint(parts[2])

	data := ""
	if len(parts) == 4 {
		data = parts[3]
	}

	return &Message{msgType, id, endpoint, data}, nil
}

func (m Message) String() string {
	raw := strconv.Itoa(m.Type)

	raw += ":" + m.ID

	raw += ":"
	if m.Endpoint != nil {
		raw += m.Endpoint.String()
	}

	if m.Data != "" {
		raw += ":" + m.Data
	}

	return raw
}

func NewDisconnect() *Message {
	return &Message{Type: 0}
}

func NewConnect(endpoint *Endpoint) *Message {
	return &Message{Type: 1, Endpoint: endpoint}
}

func NewHeartbeat() *Message {
	return &Message{Type: 2}
}

func NewMessageMsg(endpoint *Endpoint, msg string) *Message {
	return &Message{Type: 3, Endpoint: endpoint, Data: msg}
}

func NewJSONMessage(endpoint *Endpoint, data string) *Message {
	return &Message{Type: 4, Endpoint: endpoint, Data: data}
}

func NewEvent(endpoint *Endpoint, name string, args string) *Message {
	return &Message{Type: 5, Endpoint: endpoint, Data: args}
}

func NewACK(data string) *Message {
	return &Message{Type: 6, Data: data}
}

func NewError(endpoint *Endpoint, reason string, advice string) *Message {
	return &Message{Type: 7, Endpoint: endpoint, Data: reason + advice}
}

func NewNoop() *Message {
	return &Message{Type: 8}
}
