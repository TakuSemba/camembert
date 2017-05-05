package models

import "gopkg.in/mgo.v2/bson"

type Prefecture struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Romaji   string      `json:"romaji"`
}

func GetPrefectures() (prefectures []Prefecture, err error) {
	session := GetSession()
	defer session.Close()

	err = session.Prefectures().Find(bson.M{}).Sort("_id").All(&prefectures)
	return
}