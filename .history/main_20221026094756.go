package main

import "honnef.co/go/tools/config"

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the cofiguration json file")
	flag.Parse()
	//extract configuration
	config, _ := Configuration.Ex
}