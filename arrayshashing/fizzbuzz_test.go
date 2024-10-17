package arrayshashing_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/leetcode/arrayshashing"
	"github.com/bagashiz/leetcode/testutil"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for i, tc := range testCases {
		index := fmt.Sprint(i)
		t.Run(index, func(t *testing.T) {
			got := arrayshashing.FizzBuzz(tc.n)
			if diff := testutil.Diff(got, tc.want); diff != "" {
				t.Errorf("[case: %s] %s %s", index, testutil.Callers(), diff)
				return
			}
		})
	}
}
