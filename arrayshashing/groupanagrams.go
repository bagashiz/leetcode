package arrayshashing

// https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	groups := make(map[[26]int][]string, 0)

	for _, str := range strs {
		var index [26]int

		for _, c := range str {
			index[c-'a']++
		}

		groups[index] = append(groups[index], str)
	}

	for _, group := range groups {
		res = append(res, group)
	}

	return res
}
