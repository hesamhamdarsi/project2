package calculate

import (
	"fmt"
	"testing"
)

//test
func TestSum(t *testing.T) {
	s := Sum(3, 4, 2)
	if s != 9 {
		t.Error("calculated: ", s, " expected: 9")
	}
}

//example
func ExampleSum() {
	fmt.Println(Sum(2, 3, 4))
	//Output:
	//9
}

//benchmark
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 4)
	}
}

/*
go test
PASS
ok      github.com/hesamhamdarsi/project2/26-TestAndBenchmark/Benchmark/calculate       0.002s
-------------------------------
godoc -http=:8088
-------------------------------
go test -bench .
OR
go test -bench Sum
goos: linux
goarch: amd64
pkg: github.com/hesamhamdarsi/project2/26-TestAndBenchmark/Benchmark/calculate
BenchmarkSum-4          251614671                4.75 ns/op
PASS
ok      github.com/hesamhamdarsi/project2/26-TestAndBenchmark/Benchmark/calculate       1.681s
*/
