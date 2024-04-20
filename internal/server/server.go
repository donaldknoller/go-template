package server

import (
	"fmt"
	"go-template/internal/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type ServerInstance struct {
	e *echo.Echo
}

func NewServer() *ServerInstance {
	e := echo.New()
	s := &ServerInstance{e: e}
	s.applyRoutes()
	return s

}

func (s *ServerInstance) applyRoutes() {
	s.e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func (s *ServerInstance) Start() {
	listner := fmt.Sprintf(":%s", viper.GetString(config.SERVER_PORT))
	s.e.Start(listner)
}
