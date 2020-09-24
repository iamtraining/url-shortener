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
}

var Cfg *Config

func LoadCfg() {
	/*
		u := os.Getenv("DB_USER")
		p := os.Getenv("DB_PASSWORD")
		h := os.Getenv("DB_HOST")
		po := os.Getenv("DB_PORT")
		n := os.Getenv("DB_NAME")
		d := os.Getenv("DB_DRIVER")
		if u == "" || p == "" || h == "" || po == "" || n == "" || d == "" {
			fmt.Println("xdxd")
		}
	*/
	Cfg = &Config{
		User:      os.Getenv("DB_USER"),
		Passsword: os.Getenv("DB_PASSWORD"),
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		Driver:    os.Getenv("DB_DRIVER"),
	}
}
