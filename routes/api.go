package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab-webhook/controllers"
	"gitlab-webhook/middleware"
)

// Serve ..
func Serve(e *echo.Echo) {

	registerMiddleware(e)

	e.POST("/gitlab/:appname", controllers.ReceiveWebhook)
}

// registerMiddleware ..
func registerMiddleware(e *echo.Echo) {
	e.Use(middleware.Logging)
}
