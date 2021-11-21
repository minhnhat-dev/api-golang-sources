postgresql:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine
createdb: 
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
dropdb: 
	docker exec -it postgres14 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost/simple_bank?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
	
.PHONY: postgresql createdb dropdb migrateup migratedown test
