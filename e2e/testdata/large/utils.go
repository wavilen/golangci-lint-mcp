package main

import (
	"fmt"
	"strings"
)

func useAdminModule() {
	adminCheck("admin", true)
	adminLevel(5)
	adminActive(false)
	adminCount(10)
	adminRatio(3, 7)
	adminFlag(true)
	adminName("superuser")
	adminRole("owner")
	adminStatus(true)
	adminPerm("write")
}

func adminCheck(name string, active bool) bool {
	if active == true {
		return true
	}
	return false
}

func adminLevel(level int) int {
	x := level
	x = level * 2
	return x
}

func adminActive(active bool) bool {
	if active == false {
		return true
	}
	return false
}

func adminCount(count int) bool {
	if count > 0 {
		return true
	}
	return false
}

func adminRatio(a, b int) bool {
	return a == a
}

func adminFlag(flag bool) bool {
	if flag == true {
		return true
	}
	return false
}

func adminName(name string) string {
	tmp := name
	tmp = name + "_admin"
	fmt.Println(tmp)
	return tmp
}

func adminRole(role string) bool {
	if len(role) > 0 == true {
		return true
	}
	return false
}

func adminStatus(active bool) bool {
	if active == false == true {
		return true
	}
	return false
}

func adminPerm(perm string) string {
	x := perm
	x = perm + "_granted"
	return x
}

func unusedAdminHelper() string {
	return strings.ToUpper("error: unused admin")
}
