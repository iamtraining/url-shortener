package main

import (
	. "github.com/iamtraining/url-shortener/app"
	"github.com/iamtraining/url-shortener/configs"
	"github.com/iamtraining/url-shortener/routes"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("couldnt load .env")
	}
}

func main() {
	configs.BootApp()
	configs.LoadCfg()
	configs.BootStore()
	routes.Load()
	err := App.Listen(configs.Cfg.SrvPort)
	if err != nil {
		panic(err)
	}
}
