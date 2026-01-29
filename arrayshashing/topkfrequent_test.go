package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
	"github.com/bagashiz/leetcode/testutil"
)

var testCases = []struct {
	nums []int
	k    int
	want []int
}{
	{
		nums: []int{1, 2, 2, 3, 3, 3},
		k:    2,
		want: []int{2, 3},
	},
	{
		nums: []int{7, 7},
		k:    1,
		want: []int{7},
	},
}

func TestTopKFrequent(t *testing.T) {
	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.TopKFrequent(tc.nums, tc.k)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
				return
			}
		})
	}
}

func TestTopKFrequent2(t *testing.T) {
	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.TopKFrequent2(tc.nums, tc.k)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
				return
			}
		})
	}
}
