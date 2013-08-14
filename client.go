package socketio

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"strings"
)

func Subscribe(ch chan<- string, url, channel string) {
	session := NewSession(url)
	conn := NewConnection(session, channel)

	// Receive loop
	var rawJsonMsg string
	for {
		// Receive the message through websocket
		if err := websocket.Message.Receive(conn.Ws, &rawJsonMsg); err != nil {
			log.Fatal(err)
		}

		// Remove the socketio message headers
		rawJsonMsg := strings.TrimLeftFunc(rawJsonMsg, func(r rune) bool {
			if r == '{' {
				return false
			}
			return true
		})

		// ignore emtpy data and handshakes
		if len(rawJsonMsg) > 2 {
			ch <- rawJsonMsg
		}
	}
}
