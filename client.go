package socketio

import (
	"code.google.com/p/go.net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func socketIOMarshall(v interface{}) (msg []byte, payloadType byte, err error) {
	switch data := v.(type) {
	case string:
		return []byte(data), websocket.TextFrame, nil
	case []byte:
		return data, websocket.BinaryFrame, nil
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

/*

Socket-io Protocol

GET /socket.io/1 HTTP/1.1
Host: socketio.mtgox.com
Connection: keep-alive
Origin: null

GET /socket.io/1/websocket/[SESSION-ID] HTTP/1.1
Pragma: no-cache
Origin: null
Host: socketio.mtgox.com
Upgrade: websocket
Cache-Control: no-cache
Connection: Upgrade

Websocket Protocol

GET /mtgox?Currency=USD HTTP/1.1
Host: mtgox.mtgox.com
Upgrade: websocket
Connection: Upgrade
Connection: keep-alive
Origin: http://localhost/

*/
func Subscribe(ch chan<- string, url, channel string) {
	// Handshake Request
	resp, _ := http.Get("http://" + url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyParts := strings.Split(string(body), ":")
	// Agreed configs
	sessionId := bodyParts[0]
	heartbeatTimeout, _ := strconv.Atoi(bodyParts[1])
	//connectionTimeout, _ := strconv.Atoi(bodyParts[2])
	supportedProtocols := strings.Split(string(bodyParts[3]), ",")

	// Fail if websocket is not supported
	for i := 0; i < len(supportedProtocols); i++ {
		if supportedProtocols[i] == "websocket" {
			break
		} else if i == len(supportedProtocols)-1 {
			log.Fatal("Websocket is not supported")
		}
	}

	// Connect
	ws, err := websocket.Dial("ws://"+url+"/websocket/"+sessionId, "", "http://localhost/")
	if err != nil {
		log.Fatal(err)
	}

	// Initial handshake
	if err := websocket.Message.Send(ws, "1::"+channel); err != nil {
		log.Fatal(err)
	}

	// Send heartbeat in agreed timeout perios
	go func() {
		for {
			time.Sleep(time.Duration(heartbeatTimeout-1) * time.Second)
			if err := websocket.Message.Send(ws, "2::"); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Receive loop
	var msg string
	var SocketIOCodec = websocket.Codec{socketIOMarshall, socketIOUnmarshall}
	for {
		// Remove socketio headers
		if err := SocketIOCodec.Receive(ws, &msg); err != nil {
			log.Fatal(err)
		}

		// ignore emtpy data and handshakes
		if len(msg) > 2 {
			ch <- msg
		}
	}
}
