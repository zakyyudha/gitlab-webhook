package middleware

import (
	"github.com/labstack/echo/v4"
	"gitlab-webhook/utils"
)

// Logging ..
func Logging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.LogEntry(c).Info("incoming request")
		return next(c)
	}
}
