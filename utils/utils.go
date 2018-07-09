package utils

func IndexOfString(arr []string, content string) int {
	for i := range arr {
		if arr[i] == content {
			return i
		}
	}
	return -1
}
