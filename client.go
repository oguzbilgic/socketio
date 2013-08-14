package socketio

import (
	"log"
	"strings"
)

func Subscribe(ch chan<- string, url, channel string) {
	conn := NewConnection(url, channel)

	// Receive loop
	var rawJsonMsg string
	for {
		// Receive the message through websocket
		if err := conn.Receive(&rawJsonMsg); err != nil {
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
