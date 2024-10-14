package stack_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/stack"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"(])", false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := stack.IsValid(tc.s)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}
