migrate:
	migrate create -ext sql -dir ../migrations/ -seq users

migrate up:
	migrate -path ../migrations/ -database "postgresql://postgres:postgres@localhost:5432/golang_first_project?sslmode=disable" -verbose up

migrate down:
	migrate -path ../migrations/ -database "postgresql://postgres:postgres@localhost:5432/golang_first_project?sslmode=disable" -verbose down