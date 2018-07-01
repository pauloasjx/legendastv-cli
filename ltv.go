package main

import (
	"./legendastv"
	"fmt"
	"os"

	"github.com/howeyc/gopass"
)

func main() {
	var login string

	fmt.Printf("Login: ")
	fmt.Scanf("%s", &login)

	fmt.Printf("Password: ")
	password, _ := gopass.GetPasswdMasked()
	client := legendastv.Login(login, string(password))

	subtitles := client.Search(os.Args[1])

	fmt.Printf(`
 ___      _______  _______  _______  __    _  ______   _______  _______  _______  __   __         _______  ___      ___
|   |    |       ||       ||       ||  |  | ||      | |   _   ||       ||       ||  | |  |       |       ||   |    |   |
|   |    |    ___||    ___||    ___||   |_| ||  _    ||  |_|  ||  _____||_     _||  |_|  | ____  |       ||   |    |   |
|   |    |   |___ |   | __ |   |___ |       || | |   ||       || |_____   |   |  |       ||____| |       ||   |    |   |
|   |___ |    ___||   ||  ||    ___||  _    || |_|   ||       ||_____  |  |   |  |       |       |      _||   |___ |   |
|       ||   |___ |   |_| ||   |___ | | |   ||       ||   _   | _____| |  |   |   |     |        |     |_ |       ||   |
|_______||_______||_______||_______||_|  |__||______| |__| |__||_______|  |___|    |___|         |_______||_______||___|

`)

	for i, subtitle := range subtitles {
		fmt.Println(i+1, subtitle.Title)
	}
	fmt.Printf("\nChoose a subtitle: ")

	var a int
	fmt.Scanf("%d", &a)
	client.Download(subtitles[a-1])

}
