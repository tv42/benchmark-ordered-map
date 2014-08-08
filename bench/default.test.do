# go test -c does not support specifying output file, we just use two
# directories: bench/X/X.test is managed by `go test -c`, bench/X.test
# is managed by `redo`.
#
# https://code.google.com/p/go/issues/detail?id=7724

# handling Go dependencies is too hard, screw it, let "go test -c" run
# every time
redo-always

(
    cd -- "${1%.test}"
    go test -i -c
)
cp --reflink=auto -- "${1%.test}/${1##*/}" "$3"

# don't trigger re-running the benchmark if the binary doesn't change
redo-stamp <"$3"
