package arrayshashing

// https://leetcode.com/problems/contains-duplicate/
func ContainsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	m := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}

	return false
}
