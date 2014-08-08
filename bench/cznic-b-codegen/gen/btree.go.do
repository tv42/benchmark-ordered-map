DIR="$(go list -f '{{.Dir}}' github.com/cznic/b)"
redo-ifchange "$DIR/btree.go"
make --silent -C "$DIR" generic \
| sed 's/KEY/bench.Key/g; s/VALUE/bench.Value/g; s|^\(package .*\)$|\1; import bench "github.com/tv42/benchmark-ordered-map/fixture"|'
