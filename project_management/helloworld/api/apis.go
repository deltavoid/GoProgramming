package api

import (
	"net/http"
	"github.com/labstack/echo"
)

// HelloWorld comment
func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
