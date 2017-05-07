package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/TakuSemba/camembert/services"
)

func DefineRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	v1 := e.Group("/v1")
	v1.GET("/prefectures", services.GetJsonPrefectures)
	v1.GET("/members", services.GetJsonMembers)

	v2 := e.Group("/v2")
	v2.GET("/prefectures", services.GetJsonMembers)
	v2.GET("/members", services.GetBinaryMembers)

}