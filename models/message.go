package models

import "time"

//easyjson:json
type Message struct {
	Id int `json:"id"`
	Chat int `json:"chat"`
	Author int `json:"author"`
	Text string `json:"text"`
	Created time.Time `json:"created_at"`
}

