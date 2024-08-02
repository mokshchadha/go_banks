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