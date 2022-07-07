package main

import (
	"fmt"
	"strconv"
)

// Link: https://leetcode.com/problems/fizz-buzz/

func fizzBuzz(n int) []string {
	// create a slice to store the result
	result := make([]string, n+1)

	// loop through 1 to n
	for i := 1; i <= n; i++ {
		// if i is divisible by 3 and 5, then insert "FizzBuzz"
		if i%3 == 0 && i%5 == 0 {
			result[i] = "FizzBuzz"
		} else if i%3 == 0 {
			result[i] = "Fizz"
		} else if i%5 == 0 {
			result[i] = "Buzz"
		} else {
			result[i] = strconv.Itoa(i)
		}
	}
	// remove the first element of the slice
	result = result[1:]

	return result
}

func main() {
	// test cases
	fmt.Println(fizzBuzz(9))  // ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "FizzBuzz"]
	fmt.Println(fizzBuzz(15)) // ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "FizzBuzz", "11", "Fizz", "13", "14", "FizzBuzz"]
	fmt.Println(fizzBuzz(25)) // ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "FizzBuzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz"]
}
