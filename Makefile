postgres :
	docker-compose up

migrateup :
	 migrate -path db/migrations -database "postgresql://moksh:@localhost:5432/simple_banks?sslmode=disable" -verbose up 

migratedown :
	 migrate -path db/migrations -database "postgresql://moksh:@localhost:5432/simple_banks?sslmode=disable" -verbose -verbose down

sqlc :
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/mokshchadha/go_banks/db/sqlc Store