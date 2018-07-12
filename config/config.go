package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

var (
	mu   sync.RWMutex
	conf Config
)

type Config struct {
	Prefix    string
	NickLimit int
	Token     string
	AdminID   string
}

func Get() Config {
	mu.RLock()
	defer mu.RUnlock()

	return conf
}

func Update(new Config) error {
	mu.Lock()
	defer mu.Unlock()

	conf = new
	return save()
}

func Open() error {
	mu.Lock()
	defer mu.Unlock()

	buf, err := ioutil.ReadFile("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			conf = Config{
				Prefix:    "%",
				NickLimit: 5,
				Token:     "'Write Token Here!'",
				AdminID:   "AdminID",
			}
			err = save()
		}
		return err
	}

	err = json.Unmarshal(buf, &conf)
	return err
}

func save() error {
	buf, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.json", buf, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CheckAdmin(ID string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return ID == conf.AdminID
}
