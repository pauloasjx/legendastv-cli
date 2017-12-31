package main

import (
	"./legendastv"
	"encoding/json"
	"io/ioutil"
)

type (
	Config struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

func main() {
	var config Config

	config_file, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(config_file, &config)

	client := legendastv.Login(config.Login, config.Password)
	subtitles := legendastv.Search(client, "mr robot s03e03")
	legendastv.Download(client, subtitles[0])

}
