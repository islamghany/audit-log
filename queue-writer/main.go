package main

import (
	"fmt"
	"log"
	"os"
	"queue-writer/api"
	"queue-writer/broker"
	"time"
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

func main() {

	// 1- load the env vars
	var conf config

	loadEnvVars(&conf)

	// 2- connect to the database.

	broker, err := Connect("Message Broker", 10, 1*time.Second, func() (*broker.RedisBroker, error) {
		return broker.NewRedisClient(conf.redis.host, conf.redis.port, conf.redis.username, conf.redis.password)
	})

	if err != nil {
		log.Fatal(err)
	}

	defer broker.Client.Close()

	s := api.NewServer(broker)
	log.Printf("Connected to server on port %s \n", conf.port)
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

func Connect[T any](connectName string, counts int64, backOff time.Duration, fn func() (*T, error)) (*T, error) {
	var connection *T

	for {
		c, err := fn()
		if err == nil {
			log.Println("connected to: ", connectName)
			connection = c
			break
		}

		log.Printf("%s not yet read", connectName)
		counts--
		if counts == 0 {
			return nil, fmt.Errorf("can not connect to the %s", connectName)
		}
		backOff = backOff + (time.Second * 2)

		log.Println("Backing off.....")
		time.Sleep(backOff)
		continue

	}
	return connection, nil
}
