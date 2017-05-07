package services

import (
	"github.com/labstack/echo"
	"net/http"
	proto1 "github.com/golang/protobuf/proto"
	"github.com/TakuSemba/camembert/models"
	"github.com/TakuSemba/camembert/conv"
)

func GetJsonMembers(c echo.Context) error {

	members, _ := models.GetMembers()

	data := models.Members{
		Members: members,
	}

	return c.JSON(http.StatusOK, data)
}

func GetBinaryMembers(c echo.Context) error {

	members, _ := models.GetMembers()

	res := conv.ToGetMembersResponse(members)
	data, _ := proto1.Marshal(&res)

	return c.Blob(http.StatusOK, "application/protobuf", data)
}