package main

import "github.com/tiagoncardoso/golang-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}
