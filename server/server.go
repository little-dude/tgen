package server

import (
	schema "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server/log"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

// Serve waits for clients and starts a new session for each of them
func Serve() {
	log.InitLogging()
	listener, e := net.Listen("tcp", ":1234") // Listen for incoming connections
	if e != nil {
		log.Error.Println("Error listening: ", e.Error())
		os.Exit(1)
	}
	defer listener.Close() // Close the listener when the application closes.
	log.Info.Println("Listening on TCP port 1234")

	controller := schema.Controller_ServerToClient(&Controller{})
	for true {
		connection, e := listener.Accept()
		log.Info.Println("New connection")
		if e != nil {
			log.Error.Println("Error accepting connection: ", e.Error())
			os.Exit(1)
		}
		defer connection.Close()
		go rpc.NewConn(rpc.StreamTransport(connection), rpc.MainInterface(controller.Client)).Wait()
	}
}
