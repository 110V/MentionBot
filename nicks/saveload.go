package nicks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TUser struct {
	Id       string
	Nicklist []string
}
type TNick struct {
	Users []TUser
}

var GNick TNick

func SaveNickList(nick TNick) {
	bytes, err := json.Marshal(nick)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("nicks.json", bytes, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("리스트를 저장했습니다")
}

func OpenNickList() bool {
	bytes, err := ioutil.ReadFile("nicks.json")
	if err != nil {
		if os.IsNotExist(err) {
			SaveNickList(GNick)
		}
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(bytes, &GNick)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
