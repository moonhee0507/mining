package main

import (
	"flag"

	"mining/app"
	"mining/config"
)

var (
	configFlag = flag.String("environment", "./config/environment.toml", "environment toml file not found")
	difficulty = flag.Int("difficulty", 12, "difficulty err")
)

func main() {
	flag.Parse()
	c := config.NewConfig(*configFlag)

	app.NewApp(c)
}