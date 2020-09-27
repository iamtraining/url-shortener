package configs

import (
	"os"
)

type Config struct {
	User      string
	Passsword string
	Host      string
	Port      string
	DBName    string
	Driver    string
	SrvPort   string
	SrvAddr   string
}

var Cfg *Config

func LoadCfg() {
	Cfg = &Config{
		User:      os.Getenv("DB_USER"),
		Passsword: os.Getenv("DB_PASSWORD"),
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		Driver:    os.Getenv("DB_DRIVER"),
		SrvPort:   os.Getenv("SRV_PORT"),
		SrvAddr:   os.Getenv("SRV_ADDR"),
	}
}
