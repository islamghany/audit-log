package main

import (
	"context"
	"log"
	"os"
	"queue-writer/api"
	"queue-writer/broker"
)

type config struct {
	port string // port number e.g. 8080

	redis struct {
		host     string
		port     string
		password string
		username string
	}
}

var ctx = context.Background()

func main() {

	// 1- load the env vars
	var conf config

	loadEnvVars(&conf)

	// 2- connect to the database.

	broker, err := broker.NewRedisClient(conf.redis.host, conf.redis.port, conf.redis.username, conf.redis.password)
	if err != nil {
		log.Fatal(err)
	}

	defer broker.Client.Close()

	s := api.NewServer(broker)

	log.Fatal(s.Serve(conf.port))

}

func loadEnvVars(conf *config) {

	conf.port = optionalString(os.Getenv("PORT"), "8001")
	conf.redis.host = optionalString(os.Getenv("REDIS_HOST"), "localhost")
	conf.redis.port = optionalString(os.Getenv("REDIS_PORT"), "6739")
}

func optionalString(s, p string) string {
	if s == "" {
		return p
	}
	return s
}
