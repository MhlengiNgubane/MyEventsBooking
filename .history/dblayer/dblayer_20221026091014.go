package dblayer

import (

)

type DBTYPE string

const {
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
}

func Neew