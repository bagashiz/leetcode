package main

// Link: https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

func subtractProductAndSum(n int) int {
	// create a variable to store the product with initial value 1
	product := 1

	// create a variable to store the sum with initial value 0
	sum := 0

	// loop through the digits of the number
	for n > 0 {
		// get the last digit of the number using modulus operator
		digit := n % 10

		// multiply the product by the digit
		product *= digit

		// add the digit to the sum
		sum += digit

		// remove the last digit from the number
		n /= 10
	}

	// return the difference between the product and the sum
	return product - sum
}

func main() {
	// test cases
	println(subtractProductAndSum(234))  // 15
	println(subtractProductAndSum(4421)) // 21
}
