package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {

			got := twoSum(test.nums, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", test.nums, test.target, got, test.want)
			}
		})
	}
}