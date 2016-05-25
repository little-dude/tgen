package main

import (
	// "fmt"
	// schemas "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server"
	// "golang.org/x/net/context"
	// "net"
	// "zombiezen.com/go/capnproto2/rpc"
)

// func client(ctx context.Context, c net.Conn) error {
// 	// Create a connection that we can use to get the HashFactory.
// 	conn := rpc.NewConn(rpc.StreamTransport(c))
// 	defer conn.Close()
// 	// Get the "bootstrap" interface.  This is the capability set with
// 	// rpc.MainInterface on the remote side.
// 	controller := schemas.Controller{Client: conn.Bootstrap(ctx)}
//
// 	// Now we can call methods on controller, and they will be sent over c.
// 	s, _ := controller.GetPorts(
// 		ctx,
// 		func(params schemas.Controller_getPorts_Params) error {
// 			return nil
// 		}).Struct()
// 	p, _ := s.Ports()
// 	for i := 0; i < p.Len(); i++ {
// 		ptr, _ := p.PtrAt(i)
// 		itf := ptr.Interface().Client()
// 		port := schemas.Port{Client: itf}
// 		c, _ := port.GetConfig(
// 			ctx,
// 			func(params schemas.Port_getConfig_Params) error {
// 				return nil
// 			}).Struct()
// 		config, _ := c.Config()
// 		fmt.Println(config.Name())
// 	}
// 	return nil
// }
//
func main() {
	// go server.Serve()
	// conn, _ := net.Dial("tcp", "localhost:1234")
	// defer conn.Close()
	// client(context.Background(), conn)
	server.Serve()
}
