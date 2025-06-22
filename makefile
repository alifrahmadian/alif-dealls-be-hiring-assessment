APP_NAME=alif-dealls-be-hiring-assessment
MAIN=./cmd

.PHONY: run build clean tidy lint test

run:
	go run $(MAIN)