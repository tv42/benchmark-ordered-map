redo-ifchange bench.list
while read f; do
    redo-ifchange "bench/$f.results"
    perl -ne 'print "$1\n" if m{^Benchmark([^_\s]+)}' <"bench/$f.results"
done <bench.list \
	|sort -u \
>"$3"

redo-stamp <"$3"
