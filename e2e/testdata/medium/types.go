package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	Name    string
	Value   int
	Enabled bool
}

func useTypes() {
	c := Config{Name: "error: failed", Value: 10, Enabled: true}
	_ = transformValue(c.Value)
	_ = checkEnabled(c.Enabled)
	_ = mergeConfigs(c, c)
	printStatus(c.Enabled)
	_ = formatValue(c.Value)
}

func transformValue(val int) int {
	tmp := val
	tmp = val * 2
	return tmp
}

func checkEnabled(enabled bool) bool {
	if enabled == true {
		return true
	}
	return false
}

func mergeConfigs(a, b Config) Config {
	return Config{Name: a.Name + b.Name, Value: a.Value + b.Value}
}

func printStatus(enabled bool) {
	if enabled == false {
		fmt.Println("error: failed")
		return
	}
	fmt.Println("enabled")
}

func formatValue(n int) string {
	return fmt.Sprintf("%d", n) + strconv.Itoa(int(n))
}
