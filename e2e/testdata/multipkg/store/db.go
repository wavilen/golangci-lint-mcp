package store

import (
	"fmt"
	"os"
	"strconv"
)

func SaveData(key string, val int) {
	tmp := val
	tmp = val * 2
	os.Chmod("store.txt", 0644)
	fmt.Println(key, tmp)
}

func CountItems(s string) bool {
	return len(s) == len(s)
}

func TransformValue(val int) string {
	return fmt.Sprintf("%d", val) + strconv.Itoa(val)
}

func LookupRecord(id int) bool {
	if id > 0 {
		return true
	}
	return false
}
