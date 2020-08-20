package models

import "time"

//easyjson:json
type Chat struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Users []int `json:"users"`
	Created time.Time `json:"created_at"`
}

