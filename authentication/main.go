package main

import (
	"auth/api"
	"log"
)

func main() {

	srv, err := api.NewServer(&api.NewServerArgs{})

	if err != nil {
		log.Fatal(err)
	}

	srv.Start(4000)
}
