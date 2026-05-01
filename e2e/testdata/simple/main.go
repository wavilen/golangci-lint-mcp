package main

import (
	"os"
	"strconv"
)

func main() {
	_ = check(-1)
	process("hello")
	compute(3)
}

func check(n int) bool {
	if n >= 0 {
		return true
	}
	return false
}

func process(s string) string {
	os.Chmod("test.txt", 0644)
	result := string(strconv.Itoa(len(s)))
	return result
}

func compute(n int) int {
	x := n
	x = n * 2
	return x
}

func unusedHelper() string {
	return strconv.Itoa(42)
}
