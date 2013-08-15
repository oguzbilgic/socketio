package socketio

import (
	"strconv"
)

type IOMessage struct {
	Type     int
	Id       int
	Endpoint *IOEndpoint
	Data     string
}

func (m IOMessage) Marshall() string {
	raw := strconv.Itoa(m.Type)

	raw += ":"
	if m.Id != 0 {
		raw += strconv.Itoa(m.Id)
	}

	raw += ":"
	if m.Endpoint != nil {
		raw += m.Endpoint.String()
	}

	if m.Data != "" {
		raw += ":" + m.Data
	}

	return raw
}

func Unmarshall(raw string) *IOMessage {
	return &IOMessage{Type: 0}
}

// type Disconnect IOMessage
// type Connect IOMessage
// type Heartbeat IOMessage
// type Message IOMessage
// type JSON IOMessage
// type Event IOMessage
// type ACK IOMessage
// type Error IOMessage
// type Noop IOMessage

func NewDisconnect() *IOMessage {
	return &IOMessage{Type: 0}
}

func NewConnect(endpoint *IOEndpoint) *IOMessage {
	return &IOMessage{Type: 1, Endpoint: endpoint}
}

func NewHeartbeat() *IOMessage {
	return &IOMessage{Type: 2}
}

func NewMessage(endpoint *IOEndpoint, data string) *IOMessage {
	return &IOMessage{Type: 3, Endpoint: endpoint, Data: data}
}

func NewJSONMessage(endpoint *IOEndpoint, data string) *IOMessage {
	return &IOMessage{Type: 4, Endpoint: endpoint, Data: data}
}

func NewEvent(endpoint *IOEndpoint, data string) *IOMessage {
	return &IOMessage{Type: 5, Endpoint: endpoint, Data: data}
}

func NewACK(endpoint *IOEndpoint, data string) *IOMessage {
	return &IOMessage{Type: 6, Endpoint: endpoint, Data: data}
}

func NewError(endpoint *IOEndpoint, data string) *IOMessage {
	return &IOMessage{Type: 7, Endpoint: endpoint, Data: data}
}

func NewNoop() *IOMessage {
	return &IOMessage{Type: 8}
}
