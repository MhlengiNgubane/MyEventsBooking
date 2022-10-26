package main

import "honnef.co/go/tools/config"

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the cofiguration json file")
	flag.Parse()
	//extract configuration
	config, _ := Configuration.ExtractConfiguration(*confPath)
	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.Res))
}