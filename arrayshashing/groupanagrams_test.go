package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
	"github.com/bagashiz/leetcode/testutil"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{
				{"hat"},
				{"act", "cat"},
				{"stop", "pots", "tops"},
			},
		},
		{
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.GroupAnagrams(tc.strs)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
				return
			}
		})
	}
}
