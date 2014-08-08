redo-ifchange "$2.workload"
{
    cat "$2.workload"
    echo "ok fake"
} \
| prettybench -no-passthrough
