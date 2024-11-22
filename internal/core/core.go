package core

import (
	"lk_back/internal/applayer"
	"lk_back/internal/config"
	"lk_back/internal/httplayer"
	"log"
)

type Core struct {
	applayer  *applayer.App
	httplayer *httplayer.HttpServer
}

func Init() (*Core, error) {
	core := Core{}
	conf, err := config.LoadConfig("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка инициализации ConfigLoader: \n %v", err)
		return &core, err
	}
	server, err := httplayer.New(conf)
	if err != nil {
		log.Fatalf("Ошибка инициализации HttpServer: \n %v", err)
		return &core, err
	}
	core.httplayer = server
	return &core, nil
}
