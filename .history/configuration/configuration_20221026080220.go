package configuration

var (
	DBTypeDefault = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault = "localhost:8181"
)

type ServiceConfig struct {
	Databasetype dblayer.DBTYPE `json:"databasetype"`
}