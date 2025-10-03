package main

import "github.com/PauloPimentel-github/go-expert/APIs/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
