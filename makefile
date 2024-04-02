postgress:
	sudo docker run --name=postgres-16 -e POSTGRES_PASSWORD=123456 -p 8000:5432 -d postgres:16-alpine
createdb:
	sudo docker exec -it postgres-16 createdb -U postgres simple-bank
dropdb:
	sudo docker exec -it postgres-16 dropdb -U postgres  simple-bank
migrateup:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:8000/simple-bank?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://postgres:123456@localhost:8000/simple-bank?sslmode=disable" --verbose down