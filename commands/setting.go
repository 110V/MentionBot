package commands

//commands
type Command int8

const (
	ADMIN = Command(0)
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
	"set":        SET,
	"add":        ADD,
	"channel":    CHANNEL,
	"nick":       NICK,
	"permission": PERMISSION,
	"blacklist":  BLACKLIST,
	"all":        ALL,
	"noone":      NOONE}
