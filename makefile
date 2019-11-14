DBSERVER?=172.17.0.2
DBPORT?=3306
MKFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURDIR := $(dir $(MKFILE))

run:
	DBMS=mysql DBSERVER=${DBSERVER} DBPORT=${DBPORT} DBNAME=information_schema DBLOGIN=root DBPASSWORD=webteam@2019 go run server.go
test:
	go test
