//coverage means how much of your wrriten code is covered by test
/*
commands you can use:

go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/hesamhamdarsi/project2/26-TestAndBenchmark/05-Benchmark/calculate    0.002s
---------------------------
go test -coverprofile Testoutput.out
PASS
coverage: 83.3% of statements
ok      github.com/hesamhamdarsi/project2/26-TestAndBenchmark/05-Benchmark/calculate    0.002s
---------------------------
after we create that file, we can convert it to html to see where is not tested yet:
go tool cover -html=Testoutput.out
*/