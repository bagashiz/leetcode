package main

// Link: https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	// power of two means only one bit of n is 1 and all other bits are 0
	// example: 2 = 10, 4 = 100, 8 = 1000, ...
	// so we can use AND operation to check if n is power of two
	// if n & (n-1) equals 0, then n is power of two
	return n > 0 && n&(n-1) == 0
	// example:
	// 16 & 15 = 10000 & 1111 = 00000, therefore 16 is power of two
	// 14 & 13 = 1110 & 1101 = 1100, therefore 14 is NOT power of two
}

func main() {
	// test cases
	println(isPowerOfTwo(1))   // true
	println(isPowerOfTwo(16))  // true
	println(isPowerOfTwo(0))   // false
	println(isPowerOfTwo(3))   // false
	println(isPowerOfTwo(218)) // false
}
