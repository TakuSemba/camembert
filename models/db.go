package models

import (
	"gopkg.in/mgo.v2"
)

const (
	IP_ADDRESS = "172.17.0.2"
	Localhost = "127.0.0.1"
	DB = "camembert"
	PREFECTURES = "prefectures"
)

type (
	Session struct {
		*mgo.Session
	}
)

func GetSession() (session Session) {
	var err error
	var s *mgo.Session

	s, err = mgo.Dial(IP_ADDRESS)

	session = Session{s}
	if err != nil {
		panic(err)
	}
	return
}

func (s *Session) Prefectures() (c *mgo.Collection){
	c = s.DB(DB).C(PREFECTURES)
	return
}
