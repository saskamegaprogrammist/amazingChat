package models

import "time"

//easyjson:json
type User struct {
	Id       string    `json:"id"`
	UserName string    `json:"username"`
	Created  time.Time `json:"created_at"`
}

//easyjson:json
type UserId struct {
	UserId string `json:"user"`
}
