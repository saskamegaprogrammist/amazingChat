package models

import "time"

//easyjson:json
type Message struct {
	Id string `json:"id"`
	Chat string `json:"chat"`
	Author string `json:"author"`
	Text string `json:"text"`
	Created time.Time `json:"created_at"`
}

//easyjson:json
type Messages []Message
