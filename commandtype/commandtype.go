package commandtype

import "github.com/bwmarrin/discordgo"

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
)

var CommandMap = map[string]Command{
	"admin":   ADMIN,
	"channel": CHANNEL,
	"add":     ADD,
	"remove":  REMOVE,
	"reset":   RESET,
	"nick":    NICK,
	"help":    HELP}

//command list

type TSCommand struct {
	Command Command
	Help    string
	Run     func(s discordgo.Session, m discordgo.MessageCreate, args []string)
}
