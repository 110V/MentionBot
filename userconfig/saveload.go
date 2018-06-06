package userconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TUser struct {
	Id       string
	Nicklist []string
	Running  bool
}
type TUserConfig struct {
	Users []TUser
}

var GUserConfig TUserConfig

func SaveConfig(nick TUserConfig) {
	bytes, err := json.Marshal(nick)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("userConfig.json", bytes, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("리스트를 저장했습니다")
}

func OpenConfig() bool {
	bytes, err := ioutil.ReadFile("userConfig.json")
	if err != nil {
		if os.IsNotExist(err) {
			SaveConfig(GUserConfig)
		}
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(bytes, &GUserConfig)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
