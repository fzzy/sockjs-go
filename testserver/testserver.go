package main

import (
	"fmt"
	"github.com/fzzbt/sockjs-go/sockjs"
	"net/http"
)

func echoHandler(s sockjs.Session) {
	fmt.Println("session opened")
	
	for {
		m, err := s.Receive()
		if err != nil {
			println("ERR:",err.Error())
			break
		}
		fmt.Println("Received:", string(m))
		s.Send(m)
	}
	fmt.Println("session closing")
}

func main() {
	server := sockjs.NewServer(http.DefaultServeMux)
	defer server.Close()

	dwsconf := sockjs.NewConfig()
	dwsconf.Websocket = false

	http.Handle("/static", http.FileServer(http.Dir("./static")))
	server.Handle("/echo", echoHandler, sockjs.NewConfig())
	server.Handle("/disabled_websocket_echo", echoHandler, dwsconf)
	server.Handle("/close",
		func(s sockjs.Session) { s.Close() },
		sockjs.NewConfig())
	err := http.ListenAndServe(":8081", server)
	if err != nil {
		fmt.Println(err)
	}
}
