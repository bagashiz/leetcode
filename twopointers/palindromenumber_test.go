package twopointers_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/twopointers"
)

var testCases = []struct {
	x    int
	want bool
}{
	{x: 121, want: true},
	{x: -121, want: false},
	{x: 10, want: false},
}

func TestIsPalindromeNumber(t *testing.T) {
	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := twopointers.IsPalindromeNumber(tc.x)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}

func TestIsPalindromeNumber2(t *testing.T) {
	testCases := []struct {
		x    int
		want bool
	}{
		{x: 121, want: true},
		{x: -121, want: false},
		{x: 10, want: false},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := twopointers.IsPalindromeNumber2(tc.x)
			if got != tc.want {
				t.Errorf("[case: %s] got %v, want %v", index, got, tc.want)
				return
			}
		})
	}
}
