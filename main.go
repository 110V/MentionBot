package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/nicks"
	"github.com/bwmarrin/discordgo"
)

func main() {
	if !config.OpenConfig() {

		return
	}
	if !nicks.OpenNickList() {

		return
	}
	//ServerSetting
	discord, err := discordgo.New("Bot " + config.GConfig.Token)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	fmt.Println("당신의 토큰은", discord.Token+"입니다")

	//handles
	discord.AddHandler(NewMessageCreate)

	//Open
	err = discord.Open()
	if err != nil {
		fmt.Println(err.Error)
		return
	}

	//sc
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
