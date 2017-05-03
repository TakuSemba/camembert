FROM mongo:3.0

ADD . $GOPATH/src/github.com/TakuSemba/camembert/mongo

WORKDIR $GOPATH/src/github.com/TakuSemba/camembert/mongo

#CMD mongo 172.17.0.2:27017/sun_moon data/wiki.js
CMD mongorestore --host 172.17.0.2 --port 27017 -d prefectures --drop ./data/prefectures

EXPOSE 27017