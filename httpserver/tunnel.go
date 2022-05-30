package httpserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *Server) tunnelHandler(c echo.Context) error {
	hostName := c.Request().Host
	domain := strings.ReplaceAll(hostName, fmt.Sprintf(".%s", s.Host), "")
	tunnel := s.Store.GetTunnelByDomain(domain)
	if tunnel == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	if err := c.Request().Write(tunnel.Conn); err != nil {
		return c.JSON(http.StatusBadGateway, nil)
	}

	return nil
}
