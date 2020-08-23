package models

import "time"

//easyjson:json
type Chat struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Users   []string  `json:"users"`
	Created time.Time `json:"created_at"`
}

//easyjson:json
type ChatId struct {
	ChatId string `json:"chat"`
}

//easyjson:json
type Chats []Chat
