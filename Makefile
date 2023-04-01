lint:
	golangci-lint run
dockerup:
	docker-compose up --build

doc:
	swag init -g ../../app/main.go