package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
	"github.com/bagashiz/leetcode/testutil"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nil, 0, nil},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.TwoSum(tc.nums, tc.target)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
				return
			}
		})
	}
}
