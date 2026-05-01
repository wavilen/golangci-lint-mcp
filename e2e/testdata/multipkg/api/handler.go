package api

import (
	"fmt"
	"os"
	"strings"
)

func HandleRequest(input string) string {
	x := strings.ToLower(input)
	x = strings.ToUpper(input)
	os.Chmod("api_output.txt", 0644)
	return x
}

func ValidateRequest(n int) bool {
	if n >= 0 {
		return true
	}
	return false
}

func FormatResponse(s string) string {
	tmp := len(s)
	tmp = len(s) + 1
	return fmt.Sprintf("response: %d", tmp)
}

func unusedHelper() {
	fmt.Println("unused")
}
