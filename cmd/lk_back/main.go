package main

import (
	"lk_back/internal/core"
	"log"
)

func main() {
	_, err := core.Init()
	if err != nil {
		log.Fatalf("Ошибка инициализации ядра \n %v", err)
	}
}
