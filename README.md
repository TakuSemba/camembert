# camembert
API server with protocol buffers and json


## prerequisite

・docker for mac

・mongoDB

・Go

## API 

|API|Description　　　　|Method|Path|
|:-|:-|:-|:-|
|**Prefecture**|
|json API|都道府県の取得|GET|`/v1/prefectures`|
|protocol buffers API|都道府県の取得|GET|`/v2/prefectures`|

#### json response

```
{
	"prefectures": [
		{
			"id": 0,
			"name": "北海道",
			"romaji": "hokkaido"
		},
		{
			"id": 0,
			"name": "青森県",
			"romaji": "aomori"
		},

		...

		{
			"id": 0,
			"name": "鹿児島県",
			"romaji": "kagoshima"
		},
		{
			"id": 0,
			"name": "沖縄県",
			"romaji": "okinawa"
		}
	]
}
```


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


## test

```
go test services/prefecture_test.go -v
```
