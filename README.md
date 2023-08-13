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
go test -bench=. -benchtime=10s
goos: darwin
goarch: arm64
pkg: github.com/juev/pgscan-bench
BenchmarkRandallmlough-8   	      52	 200721934 ns/op
BenchmarkScany-8           	      81	 139823962 ns/op
PASS
ok  	github.com/juev/pgscan-bench	22.308s
```

## Note

The example of a function call works only for specific test data and your call
to this function will fail. The first run of the file, then the correction of
the output string and the second run will be successful.

This example was necessary only to make sure that the selected implementation of
the function returns the desired result.
