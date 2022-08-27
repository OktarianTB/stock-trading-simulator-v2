DB_URL=postgresql://root:secret@localhost:5432/stock_trading_simulator?sslmode=disable

network:
	docker network create stock-trading-simulator-network

postgres:
	docker run --name postgres --network stock-trading-simulator-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.4-alpine

postgreslogin:
	docker exec -it postgres psql -U root -d stock_trading_simulator

createdb:
	docker exec -it postgres createdb --username=root --owner=root stock_trading_simulator

dropdb:
	docker exec -it postgres dropdb stock_trading_simulator

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc Store

dockerserver:
	docker run --name stock-trading-simulator --network stock-trading-simulator-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres14:5432/stock_trading_simulator?sslmode=disable" stock-trading-simulator:latest

.PHONY: network postgres postgreslogin createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock dockerserver