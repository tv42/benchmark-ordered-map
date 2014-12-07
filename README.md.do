redo-ifchange README.edit.md
cat README.edit.md

printf '\n# Results\n'
redo-ifchange workloads.list
while read w; do
    printf '\n## %s\n```\n' "$w"
    redo-ifchange "$w.workload.pretty"
    cat "$w.workload.pretty"
    printf '```\n'
done <workloads.list
