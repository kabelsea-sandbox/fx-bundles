package controllers

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func NewMetrics(e *echo.Echo) {
	e.GET("/metrics", echoprometheus.NewHandler())
}
