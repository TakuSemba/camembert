package models

import "gopkg.in/mgo.v2/bson"

type Member struct {
	Id       	int         `json:"id"           bson:"index"`
	Age      	int	    `json:"age"          bson:"age"`
	EyeColor 	string      `json:"eye_color"    bson:"eye_color"`
	Name   		string      `json:"name"         bson:"name"`
	Gender   	string      `json:"gender"       bson:"gender"`
	Company   	string      `json:"company"      bson:"company"`
	Email   	string      `json:"email"        bson:"email"`
	Phone   	string      `json:"phone"        bson:"phone"`
	Address   	string      `json:"address"      bson:"address"`
	About   	string      `json:"about"        bson:"about"`
	Greeting   	string      `json:"greeting"     bson:"greeting"`
}

type Members struct {
	Members 	[]Member   `json:"members"`
}

func GetMembers() (members []Member, err error) {
	session := GetSession()
	defer session.Close()

	err = session.Members().Find(bson.M{}).Sort("_id").All(&members)
	return
}
