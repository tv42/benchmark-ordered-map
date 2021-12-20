# Ordered map Go data structure benchmarks

These results are stale.
Please see https://gitlab.com/cznic/benchmark-ordered-map for an updated version.

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

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_cznic-b-codegen         30     49.37 ms/op    3195224 B/op     1924 allocs/op
Benchmark_google-btree            20     73.34 ms/op    6812128 B/op   106917 allocs/op
Benchmark_biogo-llrb              20     87.55 ms/op    8000016 B/op   200001 allocs/op
Benchmark_kellydunn-art           20     89.73 ms/op   29546968 B/op   800745 allocs/op
Benchmark_cznic-b                 10    100.12 ms/op    6824199 B/op   201925 allocs/op
Benchmark_glennbrown-skiplist     10    108.02 ms/op   12016764 B/op   400018 allocs/op
Benchmark_3xian-elephantlist      10    145.34 ms/op    9630173 B/op   227178 allocs/op
Benchmark_huandu-skiplist         10    163.81 ms/op    9908171 B/op   400003 allocs/op
Benchmark_petar-GoLLRB            10    199.64 ms/op    8000016 B/op   200001 allocs/op
Benchmark_ryszard-skiplist         5    265.26 ms/op   66400347 B/op   500006 allocs/op
Benchmark_sortedslice              1   3884.09 ms/op   14598160 B/op       32 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           10000      207.71 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b-codegen        1000     1405.33 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree            500     2663.96 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b                 500     3709.30 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb              300     4697.60 μs/op        0 B/op   0 allocs/op
Benchmark_huandu-skiplist         300     4764.20 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     200     7196.53 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            200     9874.52 μs/op        0 B/op   0 allocs/op
Benchmark_kellydunn-art           100    13299.66 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        100    16260.66 μs/op       48 B/op   1 allocs/op
```

## SortedInsert
```
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    15.25 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                  50    25.92 ms/op    2681014 B/op     1614 allocs/op
Benchmark_google-btree_ReplaceOrInsert     30    51.05 ms/op    8437184 B/op   109903 allocs/op
Benchmark_glennbrown-skiplist              20    60.87 ms/op   12016763 B/op   400018 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       20    66.43 ms/op    8000016 B/op   200001 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       20    66.74 ms/op    8000016 B/op   200001 allocs/op
Benchmark_cznic-b                          20    69.56 ms/op    6112232 B/op   201616 allocs/op
Benchmark_kellydunn-art                    20    70.75 ms/op   29537728 B/op   799622 allocs/op
Benchmark_huandu-skiplist                  20    73.83 ms/op    9906618 B/op   400003 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     10   118.81 ms/op    8000016 B/op   200001 allocs/op
Benchmark_3xian-elephantlist               10   142.84 ms/op   12563802 B/op   237510 allocs/op
Benchmark_ryszard-skiplist                 10   151.35 ms/op   66400356 B/op   500006 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   193.65 ms/op    8000016 B/op   200001 allocs/op
```
