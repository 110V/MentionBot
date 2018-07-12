package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/110V/MentionBot/channels"
	"github.com/110V/MentionBot/commands"
	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

func main() {
	err := config.Open()
	if err != nil {
		panic(err)
	}

	err = users.Open()
	if err != nil {
		panic(err)
	}

	err = channels.Open()
	if err != nil {
		panic(err)
	}

	commands.RegistCommands()
	//ServerSetting
	discord, err := discordgo.New("Bot " + config.Get().Token)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	fmt.Println("당신의 토큰은", discord.Token+"입니다")

	//handles
	discord.AddHandler(newMessageCreate)
	discord.AddHandler(channelDelete)
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
	discord.Close()
}
