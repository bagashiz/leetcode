package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"0P", false},
		{" ", true},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.IsPalindrome(tc.s)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}
