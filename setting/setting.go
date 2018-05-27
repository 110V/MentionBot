package setting

const (
	Prefix    = "%"
	NickLimit = 5
)

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
