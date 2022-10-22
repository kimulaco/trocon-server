# bin/bash

go test -cover ./... -coverpkg ./... -race -coverprofile=coverage.out -covermode=atomic
