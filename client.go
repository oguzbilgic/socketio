package socketio

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"strings"
	"time"
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

	// Connect through websocket
	ws, err := websocket.Dial("ws://"+url+"/websocket/"+session.Id, "", "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}

	// Send initial handshake
	if err := websocket.Message.Send(ws, "1::"+channel); err != nil {
		log.Fatal(err)
	}

	// Send heartbeats periodically in a seperate goroutine
	go func() {
		for {
			time.Sleep(time.Duration(session.HeartbeatTimeout-1) * time.Second)
			if err := websocket.Message.Send(ws, "2::"); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Receive loop
	var rawJsonMsg string
	var SocketIOCodec = websocket.Codec{socketIOMarshall, socketIOUnmarshall}
	for {
		// Receive the message through websocket and remove socketio headers
		if err := SocketIOCodec.Receive(ws, &rawJsonMsg); err != nil {
			log.Fatal(err)
		}

		// ignore emtpy data and handshakes
		if len(rawJsonMsg) > 2 {
			ch <- rawJsonMsg
		}
	}
}
