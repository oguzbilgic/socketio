# Socket.IO

Package socketio implements a client for SocketIO protocol in Go language as specified in 
[socket.io-spec](https://github.com/LearnBoost/socket.io-spec)

## Usage

```go
package main

import (
	"fmt"
	"github.com/oguzbilgic/socketio"
)

func main() {
	// Open a new client connection to the given socket.io server
	// Connect to the given channel on the socket.io server
	socket, err := socketio.DialAndConnect("socketio-server.com:80", "/channel", "key=value")
	if err != nil {
		panic(err)
	}

	for {
		// Receive socketio.Message from the server
		msg, err := socket.Receive()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Type: %v, ID: '%s', Endpoint: '%s', Data: '%s' \n", msg.Type, msg.ID, msg.Endpoint, msg.Data)
	}
}
```

## Documentation 

http://godoc.org/github.com/oguzbilgic/socketio

## License

The MIT License (MIT)
