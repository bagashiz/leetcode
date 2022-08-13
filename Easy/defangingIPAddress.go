package main

import (
	"fmt"
	"strings"
)

// Link: https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	// replace all "." with "[.]" using strings.ReplaceAll
	return strings.ReplaceAll(address, ".", "[.]")
}

func main() {
	// test cases
	fmt.Println(defangIPaddr("1.1.1.1"))      // 1[.]1[.]1[.]1
	fmt.Println(defangIPaddr("255.100.50.0")) // 255[.]100[.]50[.]0
}
