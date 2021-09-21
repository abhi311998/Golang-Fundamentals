package main

import "testing"

func Test_mySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{2, 3, 4}, 9},
		test{[]int{2, 3, 4, 5, 6}, 20},
		test{[]int{-2, -3}, -5},
		test{[]int{1, 0, -1}, 0},
	}

	for _, v := range tests {
		if mySum(v.data...) != v.answer {
			t.Error("Expected ", v.answer, " , Got", v.answer)
		}
	}
}
