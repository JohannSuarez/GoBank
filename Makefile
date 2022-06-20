postgres:
	docker run --name postgres-instance -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=esth3r -d postgres

createdb:
	docker exec -it postgres-instance createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-instance dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:esth3r@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:esth3r@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
