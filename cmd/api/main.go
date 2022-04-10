package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/itolog/go_api/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs :", err)
	}

	server := api.New(config)

	// if err := server.Start(); err !=nil {
	// 	log.Fatal(err)
	// }
	log.Fatal(server.Start())
}
