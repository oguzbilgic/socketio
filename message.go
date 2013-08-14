package socketio

import (
	"strconv"
)

type IOEndpoint struct {
	Path  string
	Query string
}

func NewEndpoint(path, query string) *IOEndpoint {
	return &IOEndpoint{Path: path, Query: query}
}

func (e IOEndpoint) String() string {
	if e.Query != "" {
		return e.Path + "?" + e.Query
	}
	return e.Path
}

type IOMessage struct {
	Type     int
	Id       int
	Endpoint *IOEndpoint
	Data     string
}

func (m IOMessage) String() string {
	raw := strconv.Itoa(m.Type) + ":"

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

func Parse(raw string) *IOMessage {
	return &IOMessage{}
}
