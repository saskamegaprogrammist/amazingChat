package models

//easyjson:json
type IdModel struct {
	Id int `json:"id"`
}

func CreateId(id int) IdModel {
	return IdModel{Id:id}
}

