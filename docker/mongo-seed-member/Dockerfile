FROM mongo:3.4.4

ADD . $GOPATH/src/github.com/TakuSemba/camembert/docker/mongo-seed

WORKDIR $GOPATH/src/github.com/TakuSemba/camembert/docker/mongo-seed

CMD mongoimport --host 172.17.0.2:27017 --type json --jsonArray --db camembert --collection members --file member.json
