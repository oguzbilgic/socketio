package socketio

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Session struct {
	Url                string
	Id                 string
	HeartbeatTimeout   int
	ConnectionTimeout  int
	SupportedProtocols []string
}

func NewSession(url string) *Session {
	s := new(Session)

	// Initiate the session via http request
	response, err := http.Get("http://" + url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()

	// Extract the session configs from the response
	sessionVars := strings.Split(string(body), ":")
	s.Url = url
	s.Id = sessionVars[0]
	s.HeartbeatTimeout, _ = strconv.Atoi(sessionVars[1])
	s.ConnectionTimeout, _ = strconv.Atoi(sessionVars[2])
	s.SupportedProtocols = strings.Split(string(sessionVars[3]), ",")

	// Fail if websocket is not supported by SocketIO server
	for i, protocol := range s.SupportedProtocols {
		if protocol == "websocket" {
			break
		}

		if i == len(s.SupportedProtocols)-1 {
			log.Fatal("Websocket is not supported")
		}
	}

	return s
}
