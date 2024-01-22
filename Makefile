search:
	go run cmd/search/main.go

searchproto:
	protoc --go_out=./pkg/grpc/ --go-grpc_out=./pkg/grpc/  ./pkg/grpc/search.proto

tests:
	go test ./internal/...

coverage:
	go test -coverprofile=c.out -coverpkg=./internal/... ./internal/...
	go tool cover -html=c.out

.PHONY: search
.PHONY: searchproto
.PHONY: tests
.PHONY: coverage