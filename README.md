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
Benchmark_cznic-b-codegen         50     44.08 ms/op    3211956 B/op     2133 allocs/op
Benchmark_cznic-b                 20     87.78 ms/op    6045636 B/op   102194 allocs/op
Benchmark_glennbrown-skiplist     20     98.92 ms/op   11582330 B/op   301299 allocs/op
Benchmark_huandu-skiplist         10    153.61 ms/op    9462859 B/op   301053 allocs/op
Benchmark_google-btree            10    154.11 ms/op   11692828 B/op   333167 allocs/op
Benchmark_petar-GoLLRB            10    162.50 ms/op    8140944 B/op   200588 allocs/op
Benchmark_ryszard-skiplist         5    277.18 ms/op   66141742 B/op   404862 allocs/op
Benchmark_sortedslice              1   3800.25 ms/op   14598160 B/op       32 allocs/op


# Iterate
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           10000      206.62 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b-codegen        2000     1318.74 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b                 500     3686.69 μs/op        0 B/op   0 allocs/op
Benchmark_huandu-skiplist         500     4663.82 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     500     6121.29 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree            500     6969.98 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            100    10587.08 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        100    16117.89 μs/op       48 B/op   1 allocs/op


# SortedInsert
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    11.35 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                 100    22.24 ms/op    2695054 B/op     1790 allocs/op
Benchmark_glennbrown-skiplist              50    55.44 ms/op   11600053 B/op   301358 allocs/op
Benchmark_huandu-skiplist                  50    68.09 ms/op    9475684 B/op   301095 allocs/op
Benchmark_cznic-b                          20    68.38 ms/op    5330160 B/op   101841 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     20   104.14 ms/op    8140968 B/op   200588 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   165.83 ms/op    8140992 B/op   200588 allocs/op
Benchmark_google-btree_ReplaceOrInsert     10   179.43 ms/op   22216948 B/op   602078 allocs/op
Benchmark_ryszard-skiplist                 10   183.79 ms/op   66152510 B/op   404975 allocs/op
```
