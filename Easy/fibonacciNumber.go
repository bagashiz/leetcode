package main

// Link: https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
	// if n == 0, return 0
	if n == 0 {
		return 0
	}
	// if n == 1, return 1
	if n == 1 {
		return 1
	}
	// if n > 1, return fib(n-1) + fib(n-2)
	return fib(n-1) + fib(n-2)
}

func main() {
	// test cases
	println(fib(0))  // 0
	println(fib(1))  // 1
	println(fib(5))  // 5
	println(fib(10)) // 55
}
