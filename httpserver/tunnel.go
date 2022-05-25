package httpserver

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *Server) tunnelHandler(c echo.Context) error {
	hostName := c.Request().Host
	domain := strings.ReplaceAll(hostName, fmt.Sprintf(".%s", s.Host), "")

	fmt.Println(domain)
	return nil
}
