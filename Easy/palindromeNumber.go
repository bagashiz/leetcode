package main

// Link: https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	// if the value is negative, it is not a palindrome
	if x < 0 {
		return false
	}

	// if the value is 0, it is a palindrome
	if x == 0 {
		return true
	}

	// if the value is divisible by 10, it is not a palindrome
	if x%10 == 0 {
		return false
	}

	// create a variable to store the reverse of the value
	var reversed int

	// loop through the value
	for x > reversed {
		// add the last digit to the reversed value
		reversed = reversed*10 + x%10
		// remove the last digit from the value
		x /= 10
	}
	// return true if the value is equal to the reversed value
	// or the value is equal to the reversed value divided by 10
	return x == reversed || x == reversed/10
}

func main() {
	// test cases
	println(isPalindrome(121))  // true
	println(isPalindrome(-121)) // false
	println(isPalindrome(10))   // false
	println(isPalindrome(0))    // true
	println(isPalindrome(1))    // true
}
