package socketio

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	ID                 string
	HeartbeatTimeout   time.Duration
	ConnectionTimeout  time.Duration
	SupportedProtocols []string
}

func NewSession(url string) (*Session, error) {
	response, err := http.Get("http://" + url + "/socket.io/1")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	sessionVars := strings.Split(string(body), ":")
	if len(sessionVars) != 4 {
		return nil, errors.New("Session variables is not 4")
	}

	id := sessionVars[0]

	heartbeatTimeoutSec, _ := strconv.Atoi(sessionVars[1])
	connectionTimeoutSec, _ := strconv.Atoi(sessionVars[2])

	heartbeatTimeout := time.Duration(heartbeatTimeoutSec) * time.Second
	connectionTimeout := time.Duration(connectionTimeoutSec) * time.Second

	supportedProtocols := strings.Split(string(sessionVars[3]), ",")

	return &Session{id, heartbeatTimeout, connectionTimeout, supportedProtocols}, nil
}
