package main

import (
	"log"

	"github.com/stg35/InternTestQSOFT/api"
)

const serverAddress = "0.0.0.0:8081"

func main() {
	server := api.NewServer()
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}