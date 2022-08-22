postgres:
	docker run --name postgres --network bank-network  -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d postgres

createdb: 
	docker exec -it postgres createdb --username=postgres --owner=postgres receita_na_mao

dropdb: 
	docker exec -it postgres dropdb --username=postgres receita_na_mao

migrationsup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable" -verbose up

migrationsup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable" -verbose up 1

migrationsdown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable" -verbose down

migrationsdown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/receita_na_mao?sslmode=disable" -verbose down 1

sqlcgenerate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...	

server: 
	go run main.go

watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

mockgenerate:
	mockgen -package mockdb -destination db/mock/store.go github.com/vinicius-gregorio/receita_na_mao/db/sqlc Store 

.PHONY: postgres createdb dropdb migrationsup migrationsdown sqlcgenerate test server mockgenerate migrationsdown1 migrationsup1