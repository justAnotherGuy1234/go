include .env

MIGRATION_FOLDER=db/migration


#install gmake to run Makefile scripts

migrate-create:
		goose -dir ${MIGRATION_FOLDER} create $(name) sql

migrate-up:
		goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} up