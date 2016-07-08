package main

import (
	"flag"
	"fmt"
	"github.com/little-dude/tgen/server/log"
	"github.com/little-dude/tgen/server/schemas"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

var versionFlag bool
var portFlag int

func init() {
	const (
		version      = false
		versionUsage = "print version and build information"
		port         = 1234
		portUsage    = "port to listen to"
	)
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("    %s [-h] [-p PORT]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&versionFlag, "version", version, versionUsage)
	flag.BoolVar(&versionFlag, "v", version, versionUsage)
	flag.IntVar(&portFlag, "port", port, portUsage)
	flag.IntVar(&portFlag, "p", port, portUsage)
	flag.Parse()
}

// Serve waits for clients and starts a new session for each of them
func Serve() {
	log.InitLogging()
	listener, e := net.Listen("tcp", ":1234") // Listen for incoming connections
	if e != nil {
		log.Error.Println(e.Error())
		os.Exit(1)
	}
	defer listener.Close() // Close the listener when the application closes.
	log.Info.Println("Waiting for connections")

	controller := schemas.Controller_ServerToClient(NewController())
	for true {
		connection, e := listener.Accept()
		if e != nil {
			log.Error.Println(e.Error())
			os.Exit(1)
		}
		log.Trace.Println("New connection")
		defer connection.Close()
		go rpc.NewConn(rpc.StreamTransport(connection), rpc.MainInterface(controller.Client)).Wait()
	}
}

var version string
var build string

func main() {
	if versionFlag == true {
		fmt.Printf("Version: %s\nBuild: %s\n", version, build)
		return
	}
	Serve()
}
