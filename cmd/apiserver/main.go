package main

import(
	apiserver "awesomeProject/internal/app/apiserver"
	"flag"
	"log"
	"github.com/BurntSushi/toml"
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
	_, err := toml.D

	a := apiserver.New(config)
	err := a.Start
	if err != nil {
		log.Fatal(err)

	}
}