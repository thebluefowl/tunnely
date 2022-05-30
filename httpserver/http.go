package httpserver

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/thebluefowl/tunnely/store"
)

var tunnelHost *echo.Echo
var controlHost *echo.Echo

type Server struct {
	Host  string
	Port  int
	Store store.TunnelStore
}

func (s *Server) address() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Server) getTunnelHost() *echo.Echo {
	tunnelHost = echo.New()
	tunnelHost.Any("/*", s.tunnelHandler)
	return tunnelHost
}

func (s *Server) getControlHost() *echo.Echo {
	controlHost = echo.New()
	controlHost.Any("/*", s.controlHandler)
	return controlHost
}

func (s *Server) Start() {
	s.Store = store.GetTunnelStore()

	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		hostName := req.Host

		var host *echo.Echo
		if hostName == s.Host {
			host = s.getControlHost()
		} else {
			host = s.getTunnelHost()
		}

		host.ServeHTTP(res, req)

		return
	})

	e.Logger.Fatal(e.Start(s.address()))
}
