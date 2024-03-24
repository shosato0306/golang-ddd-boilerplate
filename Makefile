lint: 
	golangci-lint run ./...

fmt:
	go fmt ./...

api-server:
	air -c .air.toml

up:
	docker compose up -d

down:
	docker compose down

mysql:
	docker compose exec db mysql -u root -p myapp
