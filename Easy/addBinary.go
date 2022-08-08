package main

import (
	"fmt"
	"strconv"
)

// Link: https://leetcode.com/problems/add-binary/

func addBinary(a, b string) string {
	// create a variable to store the result
	var result string
	// create a variable to store the carry
	carry := 0

	// loop through the string from the end to the beginning
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		// create a variable to store the sum that equals to the carry
		sum := carry

		// if the index of a is greater than or equal to 0
		if i >= 0 {
			// convert the character at the index to an integer
			sum += int(a[i] - '0')
		}

		// if the index of b is greater than or equal to 0
		if j >= 0 {
			// convert the character at the index to an integer
			sum += int(b[j] - '0')
		}

		// set the carry to the sum divided by 2 for the next iteration
		carry = sum / 2

		// append the sum modulus 2 to the result
		result = strconv.Itoa(sum%2) + result
	}

	// if the carry is greater than 0
	if carry > 0 {
		// append 1 to the result for the carry
		result = "1" + result
	}

	return result
}

func main() {
	// test cases
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("1010", "1011")) // 10101
}
