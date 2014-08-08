redo-always

# use $3 to make sure output is created, even if empty
find ./bench/ -mindepth 1 -maxdepth 1 \
    \! -name '.*' \! -name '_*' \! -name '*~' \
    -type d \
    -printf '%P\n' \
    | sort \
    >"$3"

redo-stamp <"$3"
