package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	"strconv"
)

// Точка входа программы
func main() {
	botToken := "1454820573:AAH0prOuE7Wq_anwAi5YPmx_g1XHN_iMsbw"
	//https://api.telegram.org/bot<token>/METHOD_NAME
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	offset := 0
	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Smt went wrong: ", err.Error())
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)
	}
}

// Запрос обновлений
func getUpdates(botUrl string, offset int) ([]Update, error)  {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err := json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

// Ответ на обновления
func respond(botUrl string, update Update) (error) {
	var BotMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
			return err
	}
	_, err = http.Post(botUrl + "/sendMessaje", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}