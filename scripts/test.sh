# bin/bash

go test -cover ./... -race -coverprofile=coverage.out -covermode=atomic
