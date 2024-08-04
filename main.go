package main

import (
	"fmt"
	"net/http"
	"github.com/GokulAnithaNandakumar/discord-ping/bot"
	"github.com/GokulAnithaNandakumar/discord-ping/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bot is running!")
	bot.Start()
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
