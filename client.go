package socketio

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"strings"
)

func socketIOMarshall(v interface{}) (msg []byte, payloadType byte, err error) {
	switch data := v.(type) {
	case string:
		return []byte(data), websocket.TextFrame, nil
	}
	return nil, websocket.UnknownFrame, websocket.ErrNotSupported
}

func socketIOUnmarshall(msg []byte, payloadType byte, v interface{}) (err error) {
	switch data := v.(type) {
	case *string:
		str := strings.TrimLeftFunc(string(msg), func(r rune) bool {
			if r == '{' {
				return false
			}
			return true
		})
		*data = string(str)
		return nil
	}
	return websocket.ErrNotSupported
}

func Subscribe(ch chan<- string, url, channel string) {
	session := NewSession(url)
	conn := NewConnection(session, channel)

	// Receive loop
	var rawJsonMsg string
	var SocketIOCodec = websocket.Codec{socketIOMarshall, socketIOUnmarshall}
	for {
		// Receive the message through websocket and remove socketio headers
		if err := SocketIOCodec.Receive(conn.Ws, &rawJsonMsg); err != nil {
			log.Fatal(err)
		}

		// ignore emtpy data and handshakes
		if len(rawJsonMsg) > 2 {
			ch <- rawJsonMsg
		}
	}
}
