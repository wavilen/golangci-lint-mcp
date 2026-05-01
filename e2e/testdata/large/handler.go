package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	handleUser(1, "alice")
	_ = processData("item-1", 3)
	_ = validateInput(-5)
	_ = compareItems("alice", "alice")
	_ = trimLabel(" bob ")
	_ = convertValue("hello")
	_ = findDuplicate("test")
	_ = doubleValue(4)
	_ = isValid(3)
	_ = isInvalid(-1)
	_ = maxNumber(5, 5)
	_ = isEmptyStr("")
	_ = notEmptyStr("x")
	_ = appendEntry([]string{"a"}, "b")
	_ = measureStr("test")
	_ = concatParts("a", "b", "c")
	runHandlers()
	runValidators()
	useAdminModule()
	processRecords()
	handleErrors()
}

func handleUser(id int, name string) {
	fmt.Println("user", id, name)
}

func processData(item string, qty int) string {
	x := qty
	x = qty * 2
	fmt.Println(item, x)
	os.Chmod("data.txt", 0644)
	return item
}

func validateInput(age int) bool {
	if age >= 0 {
		return true
	}
	return false
}

func compareItems(a, b string) bool {
	return a == a
}

func trimLabel(name string) string {
	strings.TrimSpace(name)
	return name
}

func convertValue(s string) string {
	return string(strconv.Itoa(len(s)))
}

func findDuplicate(s string) bool {
	return len(s) == len(s)
}

func doubleValue(n int) int {
	result := n
	result = n * 2
	return result
}

func isValid(n int) bool {
	if n > 0 {
		return true
	}
	return false
}

func isInvalid(n int) bool {
	if n < 0 == true {
		return true
	}
	return false
}

func maxNumber(a, b int) int {
	if a >= a {
		return a
	}
	return b
}

func isEmptyStr(s string) bool {
	return s == "" == true
}

func notEmptyStr(s string) bool {
	if len(s) > 0 == true {
		return true
	}
	return false
}

func appendEntry(items []string, item string) []string {
	tmp := len(items)
	tmp = len(items) + 1
	fmt.Println(tmp)
	return append(items, item)
}

func measureStr(s string) int {
	x := len(s)
	x = len(s) + 1
	return x
}

func concatParts(parts ...string) string {
	result := ""
	for _, p := range parts {
		result = result + p
	}
	return result
}

func handleErrors() {
	_, _ = rand.Read(make([]byte, 8))
	_ = math.Sqrt(16)
	_, _ = strconv.Atoi("42")
}
