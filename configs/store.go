package configs

import (
	"fmt"

	. "github.com/iamtraining/url-shortener/app"
	"github.com/iamtraining/url-shortener/store"
)

func BootStore() {
	var err error
	url := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable",
		Cfg.Driver,
		Cfg.User,
		Cfg.Passsword,
		Cfg.Host,
		Cfg.Port,
		Cfg.DBName)
	Store, err = store.NewStore(url)
	if err != nil {
		panic(err)
	}
}

//"postgres://postgres:1111@localhost/shortener?sslmode=disable"
