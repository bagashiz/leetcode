package main

// Link: https://leetcode.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
	// create a map to store the characters of jewels string and pair it with a value of bool
	jMap := make(map[rune]bool)

	// loop through the jewels string
	for _, j := range jewels {
		// add the character to the map with a value of true
		jMap[j] = true
	}

	// create a variable to store the number of jewels in the stones string
	count := 0

	// loop through the stones string
	for _, s := range stones {
		// check if the character of the stones string is in the map of jewels
		if jMap[s] {
			// if it is, increment the count
			count++
		}
	}

	return count
}

func main() {
	// test cases
	println(numJewelsInStones("aA", "aAAbbbb")) // 3
	println(numJewelsInStones("z", "ZZ"))       // 0
}
