package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
)

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.IsAnagram(tc.input1, tc.input2)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}
