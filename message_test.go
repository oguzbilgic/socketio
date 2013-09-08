package socketio

import (
	"testing"
)

func TestNewMessageDissconnect(t *testing.T) {
	connectMsg, _ := ParseMessage("0::")
	if connectMsg.String() != "0::" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Type != 0 {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Id != "" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Endpoint.String() != "" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Data != "" {
		t.Errorf("Disconnect message string")
	}
}

func TestNewMessageConnect(t *testing.T) {
	connectMsg, _ := ParseMessage("1::")
	if connectMsg.String() != "1::" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Type != 1 {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Id != "" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Endpoint.String() != "" {
		t.Errorf("Disconnect message string")
	}

	if connectMsg.Data != "" {
		t.Errorf("Disconnect message string")
	}
}

func TestDisconnect(t *testing.T) {
	m := NewDisconnect()
	if m.String() != "0::" {
		t.Errorf("Disconnect message string")
	}
}

func TestConnect(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewConnect(endpoint)
	if m.String() != "1::/path?Key=Value" {
		t.Errorf("Connect message string")
	}
}

func TestHeartbeat(t *testing.T) {
	m := NewHeartbeat()
	if m.String() != "2::" {
		t.Errorf("Connect message string")
	}
}

func TestMessage(t *testing.T) {
	endpoint := NewEndpoint("/path", "Key=Value")
	m := NewMessageMsg(endpoint, "This is a test message")
	if m.String() != "3::/path?Key=Value:This is a test message" {
		t.Errorf("Connect message string")
	}
}
