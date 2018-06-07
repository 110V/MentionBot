package userconfig

import (
	"errors"
	"fmt"

	"github.com/110V/MentionBot/config"
	"github.com/110V/MentionBot/consts"

	"github.com/110V/MentionBot/utils"
)

func FindUser(users *TUserConfig, ID string) *TUser {
	/*for _, user := range users.Users {
		if user.Id == ID {
			return &user
		}
	}*/
	for i := range users.Users {
		if users.Users[i].Id == ID {
			return &users.Users[i]
		}
	}
	return nil
}
func AddNick(ID string, Anicks []string) (err error) {
	if len(Anicks) == 0 {
		err = errors.New(consts.ArgsError)
		return
	}

	user := FindUser(&GUserConfig, ID)

	if user != nil {
		for _, nick := range Anicks {
			if config.GConfig.NickLimit != 0 && config.GConfig.NickLimit <= len(user.Nicklist) {
				return
			}
			if utils.IndexOfString(user.Nicklist, nick) == -1 {
				user.Nicklist = append(user.Nicklist, nick)
			}
		}
		SaveConfig(GUserConfig)
		return
	} else {
		GUserConfig.Users = append(GUserConfig.Users, TUser{ID, nil, true})
		fmt.Println("유저없어서새로만듬")
		AddNick(ID, Anicks)
		return
	}
}

func RemoveNick(ID string, Anicks []string) (err error) {
	if len(Anicks) == 0 {
		err = errors.New("Anicks is empty")
		return
	}
	user := FindUser(&GUserConfig, ID)
	if user != nil {
		tempNicks := make([]string, 0)
		for _, nick := range user.Nicklist {
			if utils.IndexOfString(Anicks, nick) == -1 {
				tempNicks = append(tempNicks, nick)
			}
		}
		user.Nicklist = tempNicks
		SaveConfig(GUserConfig)
		return
	}
	err = errors.New("Cannot find userID")
	return
}
func ResetNick(ID string) {
	user := FindUser(&GUserConfig, ID)
	if user == nil || len(user.Nicklist) == 0 {
		return
	}
	user.Nicklist = nil
	SaveConfig(GUserConfig)
}
