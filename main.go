package main

import (
	_ "github.com/lib/pq"
	"go-gorm/internal/delivery/http"
	"log"
)

func main() {
	router := http.InitHandler()
	err := router.Start(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
