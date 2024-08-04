package main

import (
	"fmt"
	"net/http"
	"github.com/GokulAnithaNandakumar/discord-ping/bot"
	"github.com/GokulAnithaNandakumar/discord-ping/config"
)

func main(){
	err:=config.ReadConfig()

	if err!=nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}