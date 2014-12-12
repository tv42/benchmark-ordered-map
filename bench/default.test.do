# handling Go dependencies is too hard, screw it, let "go test -c" run
# every time
redo-always

go test -i -c -o "$3" "./${1%.test}"

# don't trigger re-running the benchmark if the binary doesn't change
redo-stamp <"$3"
