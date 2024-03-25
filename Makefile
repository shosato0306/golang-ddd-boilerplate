export PATH := $(PWD)/bin:$(PATH)

install:
	$(eval BIN:=$(abspath bin))
	mkdir -p ./bin
	go mod download golang.org/x/tools
	GOBIN=$(BIN) go install github.com/volatiletech/sqlboiler/v4@latest
	GOBIN=$(BIN) go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest


lint: 
	golangci-lint run ./...

format:
	$(call format)

api-server:
	air -c .air.toml

up:
	docker compose up -d

down:
	docker compose down

mysql:
	docker compose exec db mysql -u root -p myapp

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database 'mysql://root:password@tcp(localhost:3306)/myapp' -verbose up

migrate-down:
	migrate -path db/migrations -database 'mysql://root:password@tcp(localhost:3306)/myapp' -verbose down 1

sqlboiler:
	./bin/sqlboiler mysql --config ./db/sqlboiler.toml

define format
	go fmt ./... && goimports -w ./ && go mod tidy
endef
