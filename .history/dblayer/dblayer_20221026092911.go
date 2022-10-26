package dblayer

import (

)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options 
}