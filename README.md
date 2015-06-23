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

## Insert
```
benchmark                       iter       time/iter     bytes alloc             allocs
---------                       ----       ---------     -----------             ------
Benchmark_cznic-b-codegen         30     52.53 ms/op    3195224 B/op     1924 allocs/op
Benchmark_google-btree            20     74.62 ms/op    6812128 B/op   106917 allocs/op
Benchmark_biogo-llrb              20     84.64 ms/op    8000016 B/op   200001 allocs/op
Benchmark_glennbrown-skiplist     10    109.70 ms/op   12016764 B/op   400018 allocs/op
Benchmark_cznic-b                 10    110.88 ms/op    6824199 B/op   201925 allocs/op
Benchmark_3xian-elephantlist      10    161.44 ms/op    9630175 B/op   227178 allocs/op
Benchmark_petar-GoLLRB            10    193.66 ms/op    8000016 B/op   200001 allocs/op
Benchmark_huandu-skiplist          5    201.57 ms/op    9907820 B/op   400003 allocs/op
Benchmark_ryszard-skiplist         5    270.58 ms/op   66400347 B/op   500006 allocs/op
Benchmark_sortedslice              1   3847.87 ms/op   14598160 B/op       32 allocs/op
```

## Iterate
```
benchmark                        iter         time/iter   bytes alloc        allocs
---------                        ----         ---------   -----------        ------
Benchmark_sortedslice           10000      215.94 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b-codegen        1000     1484.51 μs/op        0 B/op   0 allocs/op
Benchmark_google-btree            500     2634.05 μs/op        0 B/op   0 allocs/op
Benchmark_biogo-llrb              300     4348.20 μs/op        0 B/op   0 allocs/op
Benchmark_cznic-b                 300     4662.93 μs/op        0 B/op   0 allocs/op
Benchmark_glennbrown-skiplist     200     6729.84 μs/op        0 B/op   0 allocs/op
Benchmark_huandu-skiplist         200     9453.53 μs/op        0 B/op   0 allocs/op
Benchmark_petar-GoLLRB            200     9795.90 μs/op        0 B/op   0 allocs/op
Benchmark_ryszard-skiplist        100    16593.15 μs/op       48 B/op   1 allocs/op
```

## SortedInsert
```
benchmark                                iter      time/iter     bytes alloc             allocs
---------                                ----      ---------     -----------             ------
Benchmark_sortedslice                     100    15.88 ms/op   14598160 B/op       32 allocs/op
Benchmark_cznic-b-codegen                  50    27.12 ms/op    2681014 B/op     1614 allocs/op
Benchmark_google-btree_ReplaceOrInsert     30    52.10 ms/op    8437184 B/op   109903 allocs/op
Benchmark_glennbrown-skiplist              20    63.59 ms/op   12016763 B/op   400018 allocs/op
Benchmark_biogo-llrb_ReplaceOrInsert       20    65.78 ms/op    8000016 B/op   200001 allocs/op
Benchmark_biogo-llrb_InsertNoReplace       20    65.92 ms/op    8000016 B/op   200001 allocs/op
Benchmark_cznic-b                          20    72.21 ms/op    6112231 B/op   201616 allocs/op
Benchmark_huandu-skiplist                  20    74.78 ms/op    9907169 B/op   400003 allocs/op
Benchmark_petar-GoLLRB_InsertNoReplace     10   123.03 ms/op    8000016 B/op   200001 allocs/op
Benchmark_3xian-elephantlist               10   146.51 ms/op   12563802 B/op   237510 allocs/op
Benchmark_ryszard-skiplist                 10   155.58 ms/op   66400356 B/op   500006 allocs/op
Benchmark_petar-GoLLRB_ReplaceOrInsert     10   199.05 ms/op    8000016 B/op   200001 allocs/op
```
