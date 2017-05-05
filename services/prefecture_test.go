package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	proto1 "github.com/golang/protobuf/proto"
	"github.com/k0kubun/pp"
	"github.com/TakuSemba/camembert/proto"
)

const (
	reqUrl = "http://127.0.0.1:8001"
)

func TestSum(t *testing.T) {
	actual := 10 + 20
	expected := 30
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGetTeamChannels(t *testing.T) {

	url := reqUrl + "/v2/prefectures"
	log.Println("[TEST] ", url)

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log("error", err)
	}

	log.Println("Response:", resp.StatusCode)
	if resp.StatusCode == 200 {
		resProto := proto.GetPrefecturesResponse{}
		if err = proto1.Unmarshal(resData, &resProto); err != nil {
			log.Println("unmarshaling error :", err)
		}
		pp.Println(resProto)
	} else {
		resProto := proto.ErrorResult{}
		if err = proto1.Unmarshal(resData, &resProto); err != nil {
			log.Println("unmarshaling error :", err)
		}
		pp.Println(resProto)
	}
}
