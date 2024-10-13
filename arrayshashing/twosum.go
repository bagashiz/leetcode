package arrayshashing

// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, num := range nums {
		diff := target - num
		if j, ok := m[diff]; ok && j != i {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil
}
