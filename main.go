package main

import (
	"github.com/labstack/echo/v4"
	_ "gitlab-webhook/jobs"
	"gitlab-webhook/routes"
)

var e *echo.Echo

func init() {
	e = echo.New()
}

func main() {
	routes.Serve(e)
	e.Logger.Fatal(e.Start(":4345"))
}
