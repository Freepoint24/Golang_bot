package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Точка входа программы
func main() {
	botToken := "1454820573:AAH0prOuE7Wq_anwAi5YPmx_g1XHN_iMsbw"
	//https://api.telegram.org/bot<token>/METHOD_NAME
	botApi := "https://api.telegram.org/bot"
	botUrl := botApi + botToken
	for {
		updates, err := getUpdates(botUrl)
		if err != nil {
			log.Println("Smt went wrong: ", err.Error())
		}
		fmt.Println(updates)
	}
}

// Запрос обновлений
func main(botUr string) ([]Update, error)  {
	resp, err := http.Get(botUrl + "/getUpdates")
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
func respond() {

}
