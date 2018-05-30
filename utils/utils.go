package utils

import (
	"github.com/110V/MentionBot/config"
)

func CheckAdmin(ID string) bool {
	return ID == config.GConfig.AdminID
}

func IndexOfString(arr []string, content string) int {
	for i := range arr {
		//fmt.Println(arr[i], content)
		if arr[i] == content {
			return i
		}
	}
	return -1
}
