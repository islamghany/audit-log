package main

import (
	"context"
	"database/sql"
	"db-writer/api"
	"db-writer/broker"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port string // port number e.g. 8080
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
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

	// because the servise is write-intensive, 75(in-use and idle) connection is a good number
	flag.IntVar(&conf.db.maxOpenConns, "db-max-open-conns", 75, "PostgreSQL max open connections")
	flag.IntVar(&conf.db.maxIdleConns, "db-max-idle-conns", 50, "PostgreSQL max idle connections")
	flag.StringVar(&conf.db.maxIdleTime, "db-max-idle-time", "10m", "PostgreSQL max connection idle time")
	flag.Parse()

	// 2- connect to the database.

	broker, err := Connect("Message Broker", 10, 1*time.Second, func() (*broker.RedisBroker, error) {
		return broker.NewRedisClient(conf.redis.host, conf.redis.port, conf.redis.username, conf.redis.password)
	})

	if err != nil {
		log.Fatal(err)
	}

	defer broker.Client.Close()

	db, err := Connect("PostgreSQL", 10, 1*time.Second, func() (*sql.DB, error) {
		return openDB(conf)
	})
	log.Printf("Connected to server on port %s \n", conf.port)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server := api.NewServer(db)

	log.Fatal(broker.Subscriber(ctx, "audit-log", func(s string) {
		server.CreateLogHandler(s)
	}))

}

func openDB(conf config) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf.db.dsn)

	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool. Note that
	// passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(conf.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(conf.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(conf.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	// Set the maximum idle timeout.
	db.SetConnMaxIdleTime(duration)

	// Create a context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be
	// established successfully within the 5 second deadline, then this will return an
	// error.

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}

func loadEnvVars(conf *config) {
	conf.db.dsn = os.Getenv("LOGS_DB_DSN")

	conf.port = optionalString(os.Getenv("PORT"), "80")
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
