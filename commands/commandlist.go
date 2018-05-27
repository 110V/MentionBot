package commands

//commands
type Command int8

const (
	ADMIN = iota
	SET
	ADD
	CHANNEL
	NICK
	PERMISSION
	BLACKLIST
	ALL
	NOONE
)

var CommandMap = map[string]Command{
	"set":     SET,
	"add":     ADD,
	"channel": CHANNEL,
	"nick":    NICK,
	"noone":   NOONE}

/*
	admin
	 -remove name nick
	 -add name nick
	 -channel add #~~,#~~
	 -channel add here
	 -permission use name
	 -permission admin name

	nick add ~~
	nick remove ~~
	nick reset








*/
