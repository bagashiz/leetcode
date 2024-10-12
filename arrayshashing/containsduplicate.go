package arrayshashing

// https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}

	return false
}
