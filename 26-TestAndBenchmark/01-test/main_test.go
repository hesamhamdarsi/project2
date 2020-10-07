//1- you need to create a file with name of "...._test.go" in the same directory as your package is
//2-run one/more func with a signature of "func TestXxx("t *testing.T") for any function you want to test
//		t *testing.T means a pointer to type T from package testing
//3-run "go .test" and go .test -v for more detail
package main

import "testing"

func TestSum(t *testing.T) {
	x := sum(2, 3)
	if x != 5 {
		//t.Erro is one of the outputs that we can have after test
		//here is the rest of the outputs:
		//https://golang.org/pkg/testing/#T
		t.Error("Expected 5, got", x)
	}
}
