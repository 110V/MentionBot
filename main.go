package mentionbot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	//ServerSetting
	discord, err := discordgo.New("Bot NDE2MjAzMTQyOTk3MjEzMTk0.Der7Cw.nRtFOT9Ti1yoKo4ncWESDv-zNiU")
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
