package server

import (
	"github.com/little-dude/tgen/schemas"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

// Serve waits for clients and starts a new session for each of them
func Serve() {
	InitLogging()
	listener, e := net.Listen("tcp", ":1234") // Listen for incoming connections
	if e != nil {
		Error.Println(e.Error())
		os.Exit(1)
	}
	defer listener.Close() // Close the listener when the application closes.
	Info.Println("Waiting for connections")

	controller := schemas.Controller_ServerToClient(&Controller{})
	for true {
		connection, e := listener.Accept()
		if e != nil {
			Error.Println(e.Error())
			os.Exit(1)
		}
		Trace.Println("New connection")
		defer connection.Close()
		go rpc.NewConn(rpc.StreamTransport(connection), rpc.MainInterface(controller.Client)).Wait()
	}
}
