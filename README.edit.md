# Ordered map Go data structure benchmarks

## How to run

On OSX you'll need gfind, `brew install findutils`, and it in the path as `find`.

`prettybench` in the PATH, `go get github.com/cespare/prettybench`

Run benchmarks with `redo` in the root folder. `redo` can be executed in place in a checkout of `https://github.com/apenwarr/redo`.

## Workloads

### Insert

Inserts the test data set in the order it is given (pseudorandom),
replacing existing items.

### Iterate

Iterates all values in key order, accessing each key and value.

### SortedInsert

Inserts the test data, pre-sorted, in key order.


## Requirements for a library

  - Name the benchmark functions
    `Benchmark<Workload>[_<Alternative>]`.

  - If the API takes key and value separately, use `Key` and `Value`;
    if not, use `Item`.

  - Never store pointers to `Key`, `Value` or `Item`; we can simulate
	that by changing the size of the data types, instead.

  - Remember to loop `b.N` times.

  - When iterating, remember to assign both `Key` and `Value` to `_`,
	after any type asserts if those are needed.

  - TODO: After the loop, `b.StopTimer()` and submit results to `Verify`.
