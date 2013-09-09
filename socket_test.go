package socketio

import (
	"fmt"
)

func ExampleDial() {
	socket, err := Dial("localhost:12345")
	if err != nil {
		panic(err)
	}

	for {
		msg, err := socket.Receive()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg)
	}
}

func ExampleDialAndConnect() {
	socket, err := DialAndConnect("localhost:12345", "/channel", "key=value")
	if err != nil {
		panic(err)
	}

	for {
		msg, err := socket.Receive()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg)
	}
}
