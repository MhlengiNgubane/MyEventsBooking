package dblayer

import (

)

type DBTYPE string

const {
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
}

func NewPersistenceLayer(option DBTYPE, Co)