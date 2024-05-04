postgres :
	docker-compose up

migrateup :
	 migrate -path db/migrations -database "postgresql://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable" -verbose up 

migratedown :
	 migrate -path db/migrations -database "postgresql://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable" -verbose down

sqlc :
	sqlc generate