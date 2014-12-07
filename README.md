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
Benchmark_cznic-b-codegen         30     49.52 ms/op    3195239 B/op     1924 allocs/op
Benchmark_google-btree            20     70.37 ms/op    6775184 B/op   106916 allocs/op
Benchmark_cznic-b                 20    100.70 ms/op    6824222 B/op   201926 allocs/op
Benchmark_glennbrown-skiplist     10    107.12 ms/op   12016763 B/op   400018 allocs/op
Benchmark_huandu-skiplist         10    158.35 ms/op    9908171 B/op   400003 allocs/op
Benchmark_petar-GoLLRB            10    194.30 ms/op    8000016 B/op   200001 allocs/op
Benchmark_ryszard-skiplist         5    267.60 ms/op   66400340 B/op   500006 allocs/op
Benchmark_sortedslice              1   3937.34 ms/op   14598160 B/op       32 allocs/op


# Iterate
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           10000      211.18 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b-codegen        1000     1388.93 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree            500     2662.57 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b                 500     3692.48 μs/op        0 B/op   0 allocs/op
Benchmark_huandu-skiplist         300     4662.79 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     200     6861.29 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            200     9948.53 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        100    16716.96 μs/op       48 B/op   1 allocs/op


# SortedInsert
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    15.61 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                  50    26.25 ms/op    2681010 B/op     1614 allocs/op
Benchmark_google-btree_ReplaceOrInsert     30    51.30 ms/op    8385296 B/op   109902 allocs/op
Benchmark_glennbrown-skiplist              20    63.32 ms/op   12016764 B/op   400018 allocs/op
Benchmark_cznic-b                          20    68.31 ms/op    6112231 B/op   201616 allocs/op
Benchmark_huandu-skiplist                  20    74.63 ms/op    9906620 B/op   400003 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     10   126.87 ms/op    8000016 B/op   200001 allocs/op
Benchmark_ryszard-skiplist                 10   156.13 ms/op   66400353 B/op   500006 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   193.42 ms/op    8000016 B/op   200001 allocs/op
```
