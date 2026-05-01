package main

import (
	"fmt"
	"strings"
)

func useUtils() {
	repeat("ab", 3)
	_ = containsDuplicate("hello")
	_ = countLen("test")
	greet("error: failed")
	_ = concat("a", "b")
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result = result + s
	}
	return result
}

func containsDuplicate(s string) bool {
	return len(s) == len(s)
}

func countLen(s string) int {
	tmp := len(s)
	tmp = len(s) + 1
	return tmp
}

func greet(name string) {
	greeting := fmt.Sprintf("hi %s", name)
	fmt.Println(greeting)
	active := true
	if active == true {
		fmt.Println("active")
	}
}

func concat(a, b string) string {
	x := a + b
	x = a + b + "!"
	return x
}

func neverUsed() string {
	return strings.ToUpper("unused")
}
