ENV := $(PWD)/.env

include $(ENV)

postgresinit:
	docker run --name postgres15 -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it postgres15 dropdb $(DB_NAME)

.PHONY: postgresinit postgres createdb dropdb