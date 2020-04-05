package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/yurovskiymixed/http-restapi/internal/app/apiserver"
	"log"
)
var(
	configPath string
)

func intit() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")

}

func main() {

	flag.Parse()



	config := apiserver.NewConfig()
	_, err := toml.DecodeFile()

	a := apiserver.New(config)
	err := a.Start
	if err != nil {
		log.Fatal(err)

	}
}