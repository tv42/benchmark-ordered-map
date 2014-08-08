redo-ifchange workloads.list
while read f; do
    redo-ifchange "$f.workload.pretty"
    echo 1>&2
    echo "# $f" 1>&2
    cat "$f.workload.pretty" 1>&2
    echo 1>&2
done <workloads.list
