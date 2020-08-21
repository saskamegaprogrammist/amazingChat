package models

//easyjson:json
type IdModel struct {
	Id string `json:"id"`
}

func CreateId(id string) IdModel {
	return IdModel{Id:id}
}

