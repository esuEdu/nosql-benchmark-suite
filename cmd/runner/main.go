package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/esuEdu/nosql-benchmark-suite/intenal/benchmark"
	"github.com/esuEdu/nosql-benchmark-suite/intenal/db"
)

func main() {
	ops := 50000
	if v := os.Getenv("OPS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			ops = n
		}
	}
	resultsDir := getEnv("RESULTS_DIR", "results/test")

	// Mongo
	if true {
		mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
		mongoDB, err := db.NewMongo(mongoURI, getEnv("MONGO_DB", "benchdb"), getEnv("MONGO_COLL", "tests"))
		if err != nil {
			fmt.Println("mongo err:", err)
		} else {
			if res, err := benchmark.RunAndSave(mongoDB, ops, resultsDir); err == nil {
				fmt.Println("mongo:", res)
			}
		}
	}

	// Redis
	if true {
		redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
		redisDB, err := db.NewRedis(redisAddr)
		if err != nil {
			fmt.Println("redis err:", err)
		} else {
			if res, err := benchmark.RunAndSave(redisDB, ops, resultsDir); err == nil {
				fmt.Println("redis:", res)
			}
		}
	}

	// Dynamo (only if configured)
	if ddb, err := db.NewDynamo(getEnv("DYNAMO_TABLE", "BenchTable")); err == nil {
		if res, err := benchmark.RunAndSave(ddb, ops, resultsDir); err == nil {
			fmt.Println("dynamo:", res)
		} else {
			fmt.Println("dynamo run err:", err)
		}
	} else {
		fmt.Println("dynamo conn err:", err)
	}

	// Cassandra
	if true {
		hosts := strings.Split(getEnv("CASSANDRA_HOSTS", "127.0.0.1"), ",")
		if c, err := db.NewCassandra(hosts, getEnv("CASSANDRA_KEYSPACE", "bench")); err == nil {
			if res, err := benchmark.RunAndSave(c, ops, resultsDir); err == nil {
				fmt.Println("cassandra:", res)
			}
		} else {
			fmt.Println("cassandra err:", err)
		}
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
