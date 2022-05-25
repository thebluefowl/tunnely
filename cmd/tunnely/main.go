package main

import "github.com/thebluefowl/tunnely/tcpserver"

func main() {

	// s := server.Server{
	// 	Port: 6969,
	// 	Host: "tunnely.xyz:6969",
	// }
	// s.Start()

	ts := tcpserver.NewTCPServer(4200)
	if err := ts.Listen(); err != nil {
		panic(err)
	}
}
