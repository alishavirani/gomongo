package gomongo

import (
	"errors"
)

//Our Database types types
const (
	MYSQL   = "mysql"
	SQLITE  = "sqlite"
	MONGODB = "mongodb"
)

var (
	MongoErrorNotFound = errors.New("not found") //especiall for mongo not found error | that's why "n" is in small letters | dont change it
	ErrorNotFound      = errors.New("Data not found")
	ErrorInvalidDBType = errors.New("Invalid database type")
	ErrorInvalidDriver = errors.New("Unsupported storage driver!")
)
