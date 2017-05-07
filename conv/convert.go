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
		prefProto := ToPrefecture(&pref)
		res = append(res, &prefProto)
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


func ToGetMembersResponse(members []models.Member) (res proto.GetMembersResponse) {
	membersProto := ToMembers(members)
	res = proto.GetMembersResponse{
		Members: membersProto,
	}
	return
}

func ToMembers(members []models.Member) (res []*proto.Member) {
	for _, member := range members {
		memberProto := ToMember(&member)
		res = append(res, &memberProto)
	}
	return
}

func ToMember(member *models.Member) (memberProto proto.Member) {
	memberProto = proto.Member{
		Id:        	int64(member.Id),
		Age: 		int64(member.Age),
		EyeColor:	member.EyeColor,
		Name:		member.Name,
		Gender:		member.Gender,
		Company:	member.Company,
		Email:		member.Email,
		Phone:		member.Phone,
		Address:	member.Address,
		About:		member.About,
		Greeting:	member.Greeting,

	}
	return
}