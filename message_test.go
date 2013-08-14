package socketio

import (
	"testing"
)

func TestDisconnect(t *testing.T) {
	m := NewDisconnect()
	if m.Marshall() != "0::" {
		t.Errorf("Disconnect message string")
	}
}

func TestConnect(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewConnect(endpoint)
	if m.Marshall() != "1::/path?Key=Value" {
		t.Errorf("Connect message string")
	}
}

func TestHeartbeat(t *testing.T) {
	m := NewHeartbeat()
	if m.Marshall() != "2::" {
		t.Errorf("Connect message string")
	}
}

func TestMessage(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewMessage(endpoint, "This is a test message")
	if m.Marshall() != "3::/path?Key=Value:This is a test message" {
		t.Errorf("Connect message string")
	}
}
