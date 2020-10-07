package main

import "testing"

func TestSum(t *testing.T) {

	//first we make a struct to pass the values and expected result in
	type tableTest struct {
		values []int
		result int
	}
	//then we make a slice of tests via that instruct
	test1 := []tableTest{
		tableTest{[]int{1, 2, 3}, 6},
		tableTest{[]int{3, 4}, 7},
		tableTest{[]int{10, 11, 10}, 31},
		tableTest{[]int{10, 0}, 10},
		tableTest{[]int{-1, 11}, 10},
	}
	for _, v := range test1 {
		x := sum(v.values...)
		if x != v.result {
			//t.Erro is one of the outputs that we can have after test
			//here is the rest of the outputs:
			//https://golang.org/pkg/testing/#T
			t.Error("Expected ", v.result, " got", x)
		}
	}

}
