lint: 
	golangci-lint run ./...

fmt:
	go fmt ./...

build:
	docker build -t myapp .

api-server:
	 docker run --rm -p 8080:8080 myapp api-server
