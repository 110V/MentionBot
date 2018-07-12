package channels

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/110V/MentionBot/utils"
)

var channelList []string
var mu sync.RWMutex

func Get() []string {
	mu.RLock()
	defer mu.RUnlock()

	return channelList
}

func Open() error {
	mu.Lock()
	defer mu.Unlock()

	buf, err := ioutil.ReadFile("channels.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	err = json.Unmarshal(buf, &channelList)
	return err

}

func Save() error {
	mu.Lock()
	defer mu.Unlock()

	buf, err := json.Marshal(channelList)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("channels.json", buf, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddChannel(channelID string) error {
	mu.Lock()

	if utils.IndexOfString(channelList, channelID) == -1 {
		channelList = append(channelList, channelID)
	}
	mu.Unlock()

	err := Save()
	if err != nil {
		return err
	}
	return nil
}
func RemoveChannel(channelID string) error {
	mu.Lock()

	index := utils.IndexOfString(channelList, channelID)
	if index == -1 {
		mu.Unlock()
		return nil
	}
	channelList = append(channelList[:index], channelList[index+1:]...)
	mu.Unlock()
	err := Save()
	if err != nil {
		return err
	}
	return nil
}
func Reset() error {
	mu.Lock()
	channelList = nil
	mu.Unlock()
	return Save()
}
