package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Tconfig struct {
	Prefix      string
	NickLimit   int
	Token       string
	AdminID     string
	ChannelList []string
}

var GConfig Tconfig

func SaveConfig(con Tconfig) {
	bytes, err := json.Marshal(con)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("config.json", bytes, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("콘피그를 새로 저장했습니다.")
}

func OpenConfig() bool {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			SaveConfig(Tconfig{"%", 5, "'Write Token Here!'", "AdminID", []string{}})
		}
		fmt.Println(err)
		return false
	}
	err = json.Unmarshal(bytes, &GConfig)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
