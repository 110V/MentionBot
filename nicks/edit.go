package nicks

import (
	"fmt"

	"github.com/110V/MentionBot/utils"
)

func FindUser(users *TNick, ID string) *TUser {
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
func AddNick(ID string, Anicks []string) bool {
	if len(Anicks) == 0 {
		return false
	}
	user := FindUser(&GNick, ID)
	if user != nil {
		for _, nick := range Anicks {
			if utils.IndexOfString(user.Nicklist, nick) == -1 {

				user.Nicklist = append(user.Nicklist, nick)
			}
		}
		SaveNickList(GNick)
		return true
	} else {
		GNick.Users = append(GNick.Users, TUser{ID, nil})
		fmt.Println("유저없어서새로만듬")
		AddNick(ID, Anicks)
		return true
	}
}

func RemoveNick(ID string, Anicks []string) bool {
	if len(Anicks) == 0 {
		return false
	}
	user := FindUser(&GNick, ID)
	if user != nil {
		tempNicks := make([]string, 0)
		for _, nick := range user.Nicklist {
			if utils.IndexOfString(Anicks, nick) == -1 {

				tempNicks = append(tempNicks, nick)
			}
		}
		user.Nicklist = tempNicks
		SaveNickList(GNick)
		return true
	}
	return false
}
func ResetNick(ID string) {
	user := FindUser(&GNick, ID)
	if len(user.Nicklist) != 0 {
		user.Nicklist = []string{}
		SaveNickList(GNick)
	}
}
