package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	_ = checkStatus(-1)
	processInput("hello")
	fmt.Println(convertValue(42))
}

func checkStatus(code int) bool {
	if code >= 0 {
		return true
	}
	return false
}

func processInput(s string) string {
	x := len(s)
	x = len(s) + 1
	os.Chmod("output.txt", 0644)
	return string(strconv.Itoa(x))
}

func convertValue(n int) string {
	return fmt.Sprintf("%d", n) + strconv.Itoa(n)
}
