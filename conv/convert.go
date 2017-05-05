package conv

import (
	"github.com/TakuSemba/camembert/proto"
	"github.com/TakuSemba/camembert/models"
)

func ToGetPrefecturesResponse(prefectures []models.Prefecture) (res proto.GetPrefecturesResponse) {
	prefecturesProto := ToPrefectures(prefectures)
	res = proto.GetPrefecturesResponse{
		Prefectures: prefecturesProto,
	}
	return
}

func ToPrefectures(prefectures []models.Prefecture) (res []*proto.Prefecture) {
	for _, pref := range prefectures {
		chProto := ToPrefecture(&pref)
		res = append(res, &chProto)
	}
	return
}

func ToPrefecture(pref *models.Prefecture) (chProto proto.Prefecture) {
	chProto = proto.Prefecture{
		Id:        int64(pref.Id),
		Name:      pref.Name,
		Romaji:    pref.Romaji,
	}
	return
}
