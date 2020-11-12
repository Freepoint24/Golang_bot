package main

type Update struct {
	Update  int     `json:"update_id"`
	Message Message `json:"massage"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int			`json:"chat_id"`
	Text string			`json:"text"`
}