package main

import (
	"flag"
	"fmt"
	"github.com/little-dude/tgen/server"
	"os"
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

var version string
var build string

func main() {
	if versionFlag == true {
		fmt.Printf("Version: %s\nBuild: %s\n", version, build)
		return
	}
	server.Serve()
}
