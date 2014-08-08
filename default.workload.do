workload="${2##*/}"

redo-ifchange bench.list
while read f; do
    redo-ifchange "bench/$f.results"

    perl -e 'while (<STDIN>) { next unless m{^Benchmark(?<workload>[^_\s]+)((?<alternative>\S+))?(?<rest>\s.*$)}; print "Benchmark_$ARGV[0]$+{alternative}$+{rest}\n" if $+{workload} eq $ARGV[1] }' <"bench/$f.results" "$f" "$workload"
done <bench.list \
    | sort -n --field-separator="$(printf '\t')" -k 3
