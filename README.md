# camembert
API server with protocol buffers


## prerequisite

・docker for mac

・mongoDB

・Go


## how to start locally

1. use local ip address constant

##### models/db.go
```golang

func GetSession() (session Session) {
	var err error
	var s *mgo.Session

	s, err = mgo.Dial(Localhost) // <- 127.0.0.1

	session = Session{s}
	if err != nil {
		panic(err)
	}
	return
}

```

2. start mongo db

```
sudo mongod --dbpath /var/lib/mongodb --logpath /var/log/mongodb.log
```

3. start go server

```
go run main.go
```



## how to start with docker

1. use ip address for mongo container constant

##### models/db.go
```golang

func GetSession() (session Session) {
	var err error
	var s *mgo.Session

	s, err = mgo.Dial(IP_ADDRESS) // <- ip address for mongo container

	session = Session{s}
	if err != nil {
		panic(err)
	}
	return
}

```

2. start docker

```
docker-compose build
docker-compose up
```


