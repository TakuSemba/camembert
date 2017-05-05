package services

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/TakuSemba/camembert/models"
	"github.com/TakuSemba/camembert/proto"
)

func GetJsonPrefectures(c echo.Context) error {

	prefectures, _ := models.GetPrefectures()

	return c.JSON(http.StatusOK, prefectures)
}

func GetBinaryPrefectures(c echo.Context) error {

	prefectures, _ := models.GetPrefectures()

	response := []models.Prefecture{}

	for _, prefecture := range prefectures {
		prefectureProto := proto.Prefecture{
			Id:        prefecture.Id,
			Name:      prefecture.Name,
			Romaji:   prefecture.Romaji,
		}
		response = append(response, &prefectureProto)
	}


	return c.Blob(http.StatusOK, "application/protobuf", response)
}