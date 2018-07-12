package commandtype

import (
	"github.com/110V/MentionBot/users"
	"github.com/bwmarrin/discordgo"
)

//command&Args type
type Command int8

const (
	ADMIN = iota + 1
	NICK
	CHANNEL
	ADD
	REMOVE
	RESET
	HELP
	HERE
	LIST
	ON
	OFF
	RELOAD
	STATUS
)

var CommandMap = map[string]Command{
	"admin":   ADMIN,
	"a":       ADMIN,
	"channel": CHANNEL,
	"c":       CHANNEL,
	"add":     ADD,
	"remove":  REMOVE,
	"reset":   RESET,
	"nick":    NICK,
	"n":       NICK,
	"help":    HELP,
	"here":    HERE,
	"list":    LIST,
	"on":      ON,
	"off":     OFF,
	"reload":  RELOAD,
	"status":  STATUS,
	"st":      STATUS,
}

//command list

type TSCommand struct {
	Command Command
	Use     string
	Help    string
	Run     func(s *discordgo.Session, m *discordgo.MessageCreate, user users.User, args []string)
}
