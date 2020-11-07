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
	ChatId int `json:"chat"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}
