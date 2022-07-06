package main

func romanToInt(s string) int {
	// create a map of roman numerals and their values
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// create a variable to store the total value
	total := 0

	// loop through the string
	for i := 0; i < len(s); i++ {
		// get the value of the current character
		value := roman[string(s[i])]

		// if the next character is greater than the current character, subtract the value of the next character
		if i+1 < len(s) && roman[string(s[i+1])] > value {
			value = -value
		}

		// add the value to the total
		total += value
	}

	// return the total
	return total
}

func main() {
	// test cases
	println(romanToInt("III"))     // 3
	println(romanToInt("IV"))      // 4
	println(romanToInt("IX"))      // 9
	println(romanToInt("LVIII"))   // 58
	println(romanToInt("MCMXCIV")) // 1994
}
