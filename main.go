package main

import (
	"./legendastv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	subtitles := client.Search(os.Args[1])

	for i, subtitle := range subtitles {
		fmt.Println(i+1, subtitle.Title)
	}
	fmt.Printf("\nChoose a subtitle: ")

	var a int
	fmt.Scanf("%d", &a)
	client.Download(subtitles[a-1])

}
