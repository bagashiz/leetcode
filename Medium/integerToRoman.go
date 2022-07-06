package main

func intToRoman(n int) string {
	// create slice of roman numerals
	roman := []string{
		"M",
		"CM",
		"D",
		"CD",
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}

	// create a slices of value of roman numeral
	value := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	// create a variable to store the roman numeral
	s := ""

	// loop through the slices
	for i := 0; i < len(value); i++ {
		// while the value of the current character is less than the total, add the current character to the roman numeral
		for n >= value[i] {
			s += roman[i]
			n -= value[i]
		}
	}

	// return the string
	return s
}

func main() {
	// test cases
	println(intToRoman(3))    // "III"
	println(intToRoman(4))    // "IV"
	println(intToRoman(9))    // "IX"
	println(intToRoman(58))   // "LVIII"
	println(intToRoman(1994)) // "MCMXCIV"
}
