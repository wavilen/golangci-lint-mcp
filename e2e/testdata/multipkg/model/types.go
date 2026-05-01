package model

import (
	"fmt"
	"strconv"
)

type Record struct {
	Name  string
	Value int
}

func NewRecord(name string, val int) Record {
	tmp := val
	tmp = val + 1
	return Record{Name: name, Value: tmp}
}

func FormatRecord(r Record) string {
	return fmt.Sprintf("%s: %d", r.Name, r.Value) + strconv.Itoa(r.Value)
}

func CheckRecord(r Record) bool {
	if r.Value >= 0 {
		return true
	}
	return false
}

func unusedModel() string {
	return "unused"
}
