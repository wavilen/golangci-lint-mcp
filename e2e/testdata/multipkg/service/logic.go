package service

import (
	"fmt"
	"os"
	"strings"
)

func ProcessItem(name string) string {
	active := true
	if active == true {
		fmt.Println(name)
	}
	os.Chmod("service.txt", 0644)
	return strings.ToUpper(name) + name
}

func CompareValues(a, b int) bool {
	return a == a
}

func ValidateOutput(s string) bool {
	if len(s) > 0 == true {
		return true
	}
	return false
}

func unusedService() string {
	return "never called"
}
