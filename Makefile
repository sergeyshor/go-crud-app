runpostgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

startpostgres:
	docker start cb8f5940d514

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root tweeter

dropdb:
	docker exec -it postgres15 dropdb tweeter

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/tweeter?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5432/tweeter?sslmode=disable" -verbose down

run:
	go run cmd/app/main.go

.PHONY: runpostgres startpostgres createdb dropdb migrateup migratedown run