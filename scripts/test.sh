# bin/bash

go test -cover ./... -coverpkg ./... -race -coverprofile=coverage.out -covermode=atomic

go tool cover -html coverage.out -o coverage.html
