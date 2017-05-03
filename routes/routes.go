package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/TakuSemba/camembert/services"
)

func DefineRoutes(e *echo.Echo) {

	api := e.Group("/v1")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api.GET("/prefectures", services.GetJsonPrefectures)

}