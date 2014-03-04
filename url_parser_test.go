package socketio

import (
	"testing"
)

func testHandshakeUrls(t *testing.T, rawUrls []string, expected string) {
	for _, raw := range rawUrls {
		u, err := newURLParser(raw)
		if err != nil {
			t.Errorf("NewUrl error:  %s", err)
		}
		if u.handshake() != expected {
			t.Errorf("Wrong handshake formatted url, expected: %s, actual: %s", expected, u.handshake())
		}
	}
}

func testWebsocketUrls(t *testing.T, rawUrls []string, expected string) {
	for _, raw := range rawUrls {
		u, err := newURLParser(raw)
		if err != nil {
			t.Errorf("NewUrl error:  %s", err)
		}
		ws := u.websocket("session_id")
		if ws != expected {
			t.Errorf("Wrong websocket formatted url, expected: %s, actual: %s", expected, ws)
		}
	}
}

func TestHandshakeUrl(t *testing.T) {
	testHandshakeUrls(t,
		[]string{"server.com", "http://server.com"},
		"http://server.com/socket.io/1")

	testHandshakeUrls(t,
		[]string{"server.com/path", "http://server.com/path"},
		"http://server.com/path/socket.io/1")

	testHandshakeUrls(t,
		[]string{"https://server.com"},
		"https://server.com/socket.io/1")
}

func TestWebsocketUrl(t *testing.T) {
	testWebsocketUrls(t,
		[]string{"server.com", "http://server.com"},
		"ws://server.com/socket.io/1/websocket/session_id")

	testWebsocketUrls(t,
		[]string{"server.com/path", "http://server.com/path"},
		"ws://server.com/path/socket.io/1/websocket/session_id")

	testWebsocketUrls(t,
		[]string{"https://server.com"},
		"wss://server.com/socket.io/1/websocket/session_id")
}
