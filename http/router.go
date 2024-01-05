package http

import "github.com/labstack/echo/v4"

func NewRouter(e *echo.Echo, config *Config) *echo.Group {
	return e.Group(config.HTTP.Router.Prefix)
}
