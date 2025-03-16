include .env

# make run - запустить сервер
.PHONY: run
run:
	go run cmd/diana-quiz/main.go

# make build - собрать сервер (executable)
.PHONY: build
build:
	go build -o build/ cmd/diana-quiz/main.go

# make docker-up - запустить Docker конейнеры
.PHONY: docker-up
docker-up:
	docker compose up -d

# make docker-down - остановить и удалить Docker контейнеры
.PHONY: docker-down
docker-down:
	docker compose down

# make migrate-generate name=example - создать файлы миграции с названием example
.PHONY: migrate-generate
migrate-generate:
	migrate create -ext sql -dir migrations -seq $(name)

# make migrate-up - запустить миграции БД (выполнение скриптов в файлах _up.sql)
.PHONY: migrate-up
migrate-up:
	migrate -database ${DATABASE_URL} -path migrations up

# make migrate-down - запустить миграции БД (выполнение скриптов в файлах _down.sql)
.PHONY: migrate-down
migrate-down:
	migrate -database ${DATABASE_URL} -path migrations down