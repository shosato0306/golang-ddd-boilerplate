lint: 
	golangci-lint run ./...

fmt:
	go fmt ./...

api-server:
	air -c .air.toml
