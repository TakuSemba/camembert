package services

import (
	"github.com/labstack/echo"
	"net/http"
	proto1 "github.com/golang/protobuf/proto"
	"github.com/TakuSemba/camembert/models"
	"github.com/TakuSemba/camembert/conv"
)

func GetJsonPrefectures(c echo.Context) error {

	prefectures, _ := models.GetPrefectures()

	data := models.Prefectures{
		Prefectures: prefectures,
	}

	return c.JSON(http.StatusOK, data)
}

func GetBinaryPrefectures(c echo.Context) error {

	prefectures, _ := models.GetPrefectures()

	res := conv.ToGetPrefecturesResponse(prefectures)
	data, _ := proto1.Marshal(&res)

	return c.Blob(http.StatusOK, "application/protobuf", data)
}