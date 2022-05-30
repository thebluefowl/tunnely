package main

import (
	"fmt"

	"github.com/thebluefowl/tunnely/httpserver"
	"github.com/thebluefowl/tunnely/store"
	"github.com/thebluefowl/tunnely/tcpserver"
)

func main() {

	store := store.GetTunnelStore()

	go func() {
		ts := tcpserver.NewTCPServer(4200, store)
		if err := ts.Listen(); err != nil {
			fmt.Println(err)
		}
	}()

	s := httpserver.Server{
		Port: 6969,
		Host: "tunnely.xyz:6969",
	}
	s.Start()

}
