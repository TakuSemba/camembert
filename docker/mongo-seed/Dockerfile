FROM mongo:3.0

ADD . $GOPATH/src/github.com/TakuSemba/camembert/docker/mongo-seed

WORKDIR $GOPATH/src/github.com/TakuSemba/camembert/docker/mongo-seed

CMD mongo 172.17.0.3:27017/camembert data.js