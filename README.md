# Ordered map Go data structure benchmarks

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

# Results

```
# Insert
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_cznic-b-codegen         30     49.59 ms/op    3195217 B/op     1923 allocs/op
Benchmark_google-btree            20     71.70 ms/op    6775184 B/op   106916 allocs/op
Benchmark_biogo-llrb              20     86.20 ms/op    8000016 B/op   200001 allocs/op
Benchmark_cznic-b                 10    101.83 ms/op    6824199 B/op   201925 allocs/op
Benchmark_glennbrown-skiplist     10    107.68 ms/op   12016763 B/op   400018 allocs/op
Benchmark_huandu-skiplist         10    155.92 ms/op    9908171 B/op   400003 allocs/op
Benchmark_petar-GoLLRB            10    187.61 ms/op    8000016 B/op   200001 allocs/op
Benchmark_ryszard-skiplist         5    266.40 ms/op   66400347 B/op   500006 allocs/op
Benchmark_sortedslice              1   3870.83 ms/op   14598160 B/op       32 allocs/op


# Iterate
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           10000      210.75 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b-codegen        1000     1406.31 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree            500     2609.12 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b                 500     3708.31 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb              300     4228.86 μs/op        0 B/op   0 allocs/op
Benchmark_huandu-skiplist         300     4689.53 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     200     6513.89 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            200     9491.05 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        100    16183.94 μs/op       48 B/op   1 allocs/op


# SortedInsert
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    15.31 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                  50    26.02 ms/op    2681010 B/op     1614 allocs/op
Benchmark_google-btree_ReplaceOrInsert     30    51.39 ms/op    8385296 B/op   109902 allocs/op
Benchmark_glennbrown-skiplist              20    63.47 ms/op   12016763 B/op   400018 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       20    64.16 ms/op    8000016 B/op   200001 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       20    66.92 ms/op    8000016 B/op   200001 allocs/op
Benchmark_cznic-b                          20    68.80 ms/op    6112231 B/op   201616 allocs/op
Benchmark_huandu-skiplist                  20    73.25 ms/op    9906618 B/op   400003 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     10   125.87 ms/op    8000016 B/op   200001 allocs/op
Benchmark_ryszard-skiplist                 10   153.62 ms/op   66400358 B/op   500006 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   189.04 ms/op    8000016 B/op   200001 allocs/op
```
