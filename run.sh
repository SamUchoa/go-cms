#!/bin/bash

go mod tidy

export DB_USER=samuel
export DB_PASSWD=root
export DB_NAME=teste_cms

go run main.go
