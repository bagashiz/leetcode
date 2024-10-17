package arrayshashing

import (
	"maps"
	"slices"
	"strconv"
)

// https://leetcode.com/problems/fizz-buzz/
func FizzBuzz(n int) []string {
	ans := make([]string, n)

	wordMap := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	keys := maps.Keys(wordMap)
	divisors := slices.Collect(keys)
	slices.Sort(divisors)

	for i := 1; i <= n; i++ {
		var word string

		for _, divisor := range divisors {
			if i%divisor == 0 {
				word += wordMap[divisor]
			}
		}

		if word == "" {
			ans[i-1] = strconv.Itoa(i)
		} else {
			ans[i-1] = word
		}
	}

	return ans
}
