package socketio

import (
	"errors"
	"strconv"
	"strings"
)

type IOMessage struct {
	Type     int
	Id       string
	Endpoint *Endpoint
	Data     string
}

func NewIOMessage(rawMsg string) (*IOMessage, error) {
	if len(rawMsg) == 0 {
		return nil, errors.New("Empty message")
	}

	msgType, err := strconv.Atoi(string(rawMsg[0]))
	if err != nil {
		return nil, err
	}

	switch msgType {
	case 3, 4, 5:
		parts := strings.SplitN(rawMsg, ":", 4)
		id := parts[1]
		return &IOMessage{msgType, id, nil, parts[3]}, nil
	default:
		return &IOMessage{Type: msgType}, nil
	}
}

func (m IOMessage) String() string {
	raw := strconv.Itoa(m.Type)

	raw += ":" + m.Id

	raw += ":"
	if m.Endpoint != nil {
		raw += m.Endpoint.String()
	}

	if m.Data != "" {
		raw += ":" + m.Data
	}

	return raw
}

func NewDisconnect() *IOMessage {
	return &IOMessage{Type: 0}
}

func NewConnect(endpoint *Endpoint) *IOMessage {
	return &IOMessage{Type: 1, Endpoint: endpoint}
}

func NewHeartbeat() *IOMessage {
	return &IOMessage{Type: 2}
}

func NewMessage(endpoint *Endpoint, msg string) *IOMessage {
	return &IOMessage{Type: 3, Endpoint: endpoint, Data: msg}
}

func NewJSONMessage(endpoint *Endpoint, data string) *IOMessage {
	return &IOMessage{Type: 4, Endpoint: endpoint, Data: data}
}

func NewEvent(endpoint *Endpoint, name string, args string) *IOMessage {
	return &IOMessage{Type: 5, Endpoint: endpoint, Data: args}
}

func NewACK(data string) *IOMessage {
	return &IOMessage{Type: 6, Data: data}
}

func NewError(endpoint *Endpoint, reason string, advice string) *IOMessage {
	return &IOMessage{Type: 7, Endpoint: endpoint, Data: reason + advice}
}

func NewNoop() *IOMessage {
	return &IOMessage{Type: 8}
}
