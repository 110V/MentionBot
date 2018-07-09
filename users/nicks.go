package users

import (
	"errors"

	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/consts"

	"github.com/110V/MentionBot/utils"
)

func (user *User) AddNick(nicks []string) error {
	if len(nicks) == 0 {
		return errors.New(consts.ArgsError)
	}

	tempNicks := make([]string, 0, len(user.Nicklist)+len(nicks))
	tempNicks = append(tempNicks, user.Nicklist...)

	for _, nick := range nicks {
		if utils.IndexOfString(tempNicks, nick) == -1 {
			tempNicks = append(tempNicks, nick)
		}
	}

	if config.Get().NickLimit < len(tempNicks) && config.Get().NickLimit != -1 {
		return errors.New(consts.ExceedLimit)
	}

	user.Nicklist = tempNicks

	return nil
}

func (user *User) RemoveNick(nicks []string) error {
	if len(nicks) == 0 {
		return errors.New(consts.ArgsError)
	}

	for i := range user.Nicklist {
		for j := range nicks {
			if nicks[i] == nicks[j] {
				user.Nicklist = append(user.Nicklist[:i], user.Nicklist[i+1:]...)
				i--
				nicks = append(nicks[:i], nicks[i+1:]...)
				j--
			}
		}
	}

	return nil
}

func (user *User) ResetNick() {
	user.Nicklist = nil
}
