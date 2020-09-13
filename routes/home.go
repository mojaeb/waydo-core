package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func Home (context echo.Context) error {
	return context.String(http.StatusOK, "Welcome to WayDo application\n you can use document in this url")
}