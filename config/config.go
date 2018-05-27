package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type configS struct {
	Prefix    string
	NickLimit int
	Token     string
}

var GConfig configS

func SaveConfig(con configS) {
	bytes, err := json.Marshal(GConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("config.json", bytes, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func OpenConfig() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			SaveConfig(configS{"%", 5, "'Write Token Here!'"})
		}
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &GConfig)
}
