package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	handler(1)
	process("hello")
	_ = validate(-1)
	_ = compare(3, 3)
	_ = trimIt(" hello ")
	logMsg("error: failed")
	useTypes()
	useUtils()
}

func handler(id int) {
	fmt.Println("handling", id)
}

func process(input string) {
	x := strings.ToLower(input)
	x = strings.ToUpper(input)
	fmt.Println(x)
	os.Chmod("output.txt", 0644)
}

func validate(n int) bool {
	if n >= 0 {
		return true
	}
	return false
}

func compare(a, b int) bool {
	return a == a
}

func trimIt(s string) string {
	strings.TrimSpace(s)
	return s
}

func logMsg(msg string) {
	flag := true
	if flag == true {
		fmt.Println(msg)
	}
}
