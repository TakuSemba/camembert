# camembert
API server with protocol buffers and json


## prerequisite

・docker for mac

・mongoDB

・Go

## API 

|API|Description　　　　|Method|Path|
|:-|:-|:-|:-|
|json API|都道府県の取得|GET|`/v1/prefectures`|
|protocol buffers API|都道府県の取得|GET|`/v2/prefectures`|
|json API|人物リストの取得|GET|`/v1/members`|
|protocol buffers API|人物リストの取得|GET|`/v2/members`|

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


[
  {
    "index": 1,
    "age": 25,
    "eye_color": "blue",
    "name": "Rebekah Horne",
    "gender": "female",
    "company": "REMOTION",
    "email": "rebekahhorne@remotion.com",
    "phone": "+1 (883) 507-2758",
    "address": "375 Moultrie Street, Kohatk, Nevada, 5200",
    "about": "Commodo ullamco minim Lorem non voluptate aliqua dolor dolor minim ea enim eiusmod aliquip consequat. Voluptate aute consequat ipsum officia culpa occaecat occaecat qui enim sit cillum id deserunt ut. Occaecat est magna esse dolor veniam mollit exercitation. Sunt sit incididunt amet velit veniam veniam ea ipsum amet aliqua consequat. Deserunt aute velit aliquip ex exercitation labore laboris in ut duis consectetur aliquip amet. Duis adipisicing laboris nulla mollit irure deserunt voluptate est. Ullamco enim exercitation ipsum Lorem.\r\n",
    "greeting": "Hello, Rebekah Horne"
  },
  {
    "index": 2,
    "age": 40,
    "eye_color": "brown",
    "name": "Gay Espinoza",
    "gender": "female",
    "company": "TRIPSCH",
    "email": "gayespinoza@tripsch.com",
    "phone": "+1 (809) 551-3700",
    "address": "846 Bassett Avenue, Mammoth, Indiana, 912",
    "about": "Consequat ut aute mollit consequat aute et enim ipsum. Consectetur est sunt magna aliquip consequat qui aliquip nisi magna est. Voluptate aliquip labore aliqua ipsum quis et velit qui minim. Deserunt officia officia voluptate occaecat laboris in ea culpa aliquip culpa sit tempor.\r\n",
    "greeting": "Hello, Gay Espinoza"
  },
  
		...
  
  {
    "index": 1000,
    "age": 27,
    "eye_color": "brown",
    "name": "Elena Frazier",
    "gender": "female",
    "company": "EZENT",
    "email": "elenafrazier@ezent.com",
    "phone": "+1 (845) 403-3185",
    "address": "398 Greene Avenue, Waumandee, North Dakota, 5735",
    "about": "Aute in veniam adipisicing fugiat nisi pariatur nulla irure nostrud duis enim eiusmod occaecat. Cupidatat mollit ipsum incididunt consectetur elit magna proident est reprehenderit anim mollit eu. Officia mollit sunt velit mollit incididunt sit commodo laboris tempor sunt. Occaecat amet esse excepteur exercitation labore fugiat irure Lorem veniam quis.\r\n",
    "greeting": "Hello, Elena Frazier"
  }
]
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
