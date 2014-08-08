# Ordered map Go data structure benchmarks

## Workloads

### InsertIterate

Inserts the test data set in the order it is given (pseudorandom),
replacing existing items.

Iterates all values in key order, accessing each key and value.

TODO: split this into two

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

# Results

```
# InsertIterate
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_cznic-b-codegen         50     46.23 ms/op    3211980 B/op     2133 allocs/op
Benchmark_cznic-b                 20     92.77 ms/op    6045628 B/op   102194 allocs/op
Benchmark_glennbrown-skiplist     20    106.01 ms/op   11582088 B/op   301298 allocs/op
Benchmark_google-btree            10    163.08 ms/op   11694564 B/op   333173 allocs/op
Benchmark_huandu-skiplist         10    163.58 ms/op    9444142 B/op   300999 allocs/op
Benchmark_petar-GoLLRB            10    171.60 ms/op    8140944 B/op   200588 allocs/op
Benchmark_ryszard-skiplist         5    357.94 ms/op   66041931 B/op   403822 allocs/op
Benchmark_sortedslice              1   3901.88 ms/op   14598160 B/op       32 allocs/op


# SortedInsert
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    11.24 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                 100    22.35 ms/op    2695056 B/op     1790 allocs/op
Benchmark_glennbrown-skiplist              50    56.28 ms/op   11595396 B/op   301342 allocs/op
Benchmark_huandu-skiplist                  50    68.28 ms/op    9479384 B/op   301106 allocs/op
Benchmark_cznic-b                          50    68.68 ms/op    5330160 B/op   101841 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     20   102.73 ms/op    8140968 B/op   200588 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   168.19 ms/op    8140896 B/op   200588 allocs/op
Benchmark_google-btree_ReplaceOrInsert     10   178.27 ms/op   22272955 B/op   602254 allocs/op
Benchmark_ryszard-skiplist                 10   204.17 ms/op   66146107 B/op   404908 allocs/op
```
