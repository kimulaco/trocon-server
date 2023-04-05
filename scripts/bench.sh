# bin/bash

go test -benchmem -run=^Benchmark -bench ^Benchmark ./... -benchmem
