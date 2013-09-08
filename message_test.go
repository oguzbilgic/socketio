package socketio

import (
	"testing"
)

func TestParseMessageInvalidType(t *testing.T) {
	_, err := ParseMessage("a::")
	if err == nil {
		t.Errorf("Invaild message type was not detected")
	}
}

func TestParseMessageShort(t *testing.T) {
	_, err := ParseMessage("0:")
	if err == nil {
		t.Errorf("Invaild message was not detected")
	}
}

func TestParseMessageData(t *testing.T) {
	msg, _ := ParseMessage("4:::This is data")
	if msg.Data != "This is data" {
		t.Errorf("Message data was not parsed correctly")
	}
}

func TestNewDisconnect(t *testing.T) {
	m := NewDisconnect()
	if m.String() != "0::" {
		t.Errorf("Disconnect message string")
	}
}

func TestNewConnect(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewConnect(endpoint)
	if m.String() != "1::/path?Key=Value" {
		t.Errorf("Connect message string")
	}
}

func TestNewHeartbeat(t *testing.T) {
	m := NewHeartbeat()
	if m.String() != "2::" {
		t.Errorf("Connect message string")
	}
}

func TestNewMessage(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewMessageMsg(endpoint, "This is a test message")
	if m.String() != "3::/path?Key=Value:This is a test message" {
		t.Errorf("Connect message string")
	}
}

func TestNewJSONMessage(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewJSONMessage(endpoint, "This is JSON data")
	if m.String() != "4::/path?Key=Value:This is JSON data" {
		t.Errorf("Error NewJSONMessage")
	}
}

func TestNewEvent(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewEvent(endpoint, "name", "args")
	if m.String() != "5::/path?Key=Value:args" {
		t.Errorf("Error NewEvent()")
	}
}

func TestNewACK(t *testing.T) {
	m := NewACK("data")
	if m.String() != "6:::data" {
		t.Errorf("Error NewACK")
	}
}

func TestNewError(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewError(endpoint, "reason", "advice")
	if m.String() != "7::/path?Key=Value:reason+advice" {
		t.Errorf("Error NewError")
	}
}

func TestNewNoop(t *testing.T) {
	m := NewNoop()
	if m.String() != "8::" {
		t.Errorf("Error NewNoop")
	}
}
