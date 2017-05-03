package services

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/TakuSemba/camembert/models"
)

func GetJsonPrefectures(c echo.Context) error {

	prefectures, _ := models.GetPrefectures()

	return c.JSON(http.StatusOK, prefectures)
}