# pgscan-bench

There are two libraries to simplify working with database queries:

- [randallmlough/pgxscan](github.com/randallmlough/pgxscan)
- [georgysavva/scany](github.com/georgysavva/scany/pgxscan)

I decided to compare their performance in order to understand which
implementation is preferable to use in my projects.

To do this, I implemented a test bench using docker. The database is populated
with random data.

I wrote test functions to measure the performance of a typical operation for a
database query.

## Results

```sh
$ make run
go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/juev/pgscan-bench
BenchmarkRandallmlough-8          	       5	 209839208 ns/op
BenchmarkRandallmloughScanOne-8   	    3798	    284995 ns/op
BenchmarkScany-8                  	       7	 151589583 ns/op
BenchmarkScanyScanOne-8           	    4476	    247141 ns/op
BenchmarkManual-8                 	       8	 137495333 ns/op
BenchmarkManualScanOne-8          	    4929	    208781 ns/op
PASS
ok  	github.com/juev/pgscan-bench	8.558s
```

## Note

The example of a function call works only for specific test data and your call
to this function will fail. The first run of the file, then the correction of
the output string and the second run will be successful.

This example was necessary only to make sure that the selected implementation of
the function returns the desired result.

## Update

2023/08/13 added manual data parsing.
