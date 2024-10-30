postgresinit:
	docker run --name postgres17 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Pauline@1079 -d postgres:17-bookworm
postgres:
	docker exec -it postgres17 psql
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root chatdb
dropdb:
	docker exect -it dropdb chatdb
migrateup:
	migrate -path db/migrations -database "postgres://root:Pauline@1079@localhost:5433/chatdb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgres://root:Pauline@1079@localhost:5433/chatdb?sslmode=disable" -verbose down
.PHONY:postgresinit postgres createdb dropdb migrateup migratedown
