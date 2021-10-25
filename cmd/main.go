package main

import (
	"exercise_db/internal/configure"
	"exercise_db/internal/network"
	"exercise_db/internal/service"
	"flag"
)

func main() {
	var configPath string
	flag.StringVar(&configPath,
		"configPath",
		"/Users/yimo/go/src/practise_db/cmd/configure.yml",
		"configPath")
	flag.Parse()
	config := configure.NewConfig(configPath)
	db := service.New(
		config,
		network.New(config),
	)
	go db.Start()
	<-make(chan int, 0)
}
