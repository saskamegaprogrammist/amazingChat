package models

import "time"

//easyjson:json
type User struct {
	Id int `json:"id"`
	UserName string `json:"username"`
	Created time.Time `json:"created_at"`
}
