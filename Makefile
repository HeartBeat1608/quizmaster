.PHONY: run-server build-server run-client build-client run build

run-server:
	cd server && go run cmd/main.go
	
run-client:
	cd client && pnpm dev

run: run-server run-client

build-server:
	cd server && go build -o ./bin/server cmd/main.go
	
build-client:
	cd client && pnpm build
	
build: build-server build-client
	
db:
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:alpine
	
migrate:
	cd server/migrations && go run migrate.go 
	
