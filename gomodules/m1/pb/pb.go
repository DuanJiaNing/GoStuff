package pb

import (
	"github.com/labstack/echo"
	"net/http"
)

func Hi(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
