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

// ParseMessages parses the given raw message in to Message.
func parseMessage(rawMsg string) (*Message, error) {
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

// String returns the string represenation of the Message.
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

// NewDisconnect returns a new disconnect Message.
func NewDisconnect() *Message {
	return &Message{Type: 0}
}

// NewConnect returns a new connect Message for the given endpoint.
func NewConnect(endpoint *Endpoint) *Message {
	return &Message{Type: 1, Endpoint: endpoint}
}

// NewHeartbeat returns a new heartbeat Message.
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

// NewError returns a new error Message for the given endpoint with a
// reason and an advice.
func NewError(endpoint *Endpoint, reason string, advice string) *Message {
	return &Message{Type: 7, Endpoint: endpoint, Data: reason + "+" + advice}
}

// NewNoop returns a new no-op Message.
func NewNoop() *Message {
	return &Message{Type: 8}
}
