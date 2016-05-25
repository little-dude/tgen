package server

import (
	schema "github.com/little-dude/tgen/capnp"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

// Serve waits for clients and starts a new session for each of them
func Serve() {
	InitLogging()
	listener, e := net.Listen("tcp", ":1234") // Listen for incoming connections
	if e != nil {
		Error.Println("Error listening: ", e.Error())
		os.Exit(1)
	}
	defer listener.Close() // Close the listener when the application closes.
	Info.Println("Listening on TCP port 1234")

	controller := schema.Controller_ServerToClient(&Controller{})
	for true {
		connection, e := listener.Accept()
		Info.Println("New connection")
		if e != nil {
			Error.Println("Error accepting connection: ", e.Error())
			os.Exit(1)
		}
		defer connection.Close()
		go rpc.NewConn(rpc.StreamTransport(connection), rpc.MainInterface(controller.Client)).Wait()
	}
}
