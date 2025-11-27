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
	ops := 1000
	if v := os.Getenv("OPS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			ops = n
		}
	}

	hostsEnv := getEnv("CASSANDRA_HOSTS", "127.0.0.1")
	hosts := strings.Split(hostsEnv, ",")
	keyspace := getEnv("CASSANDRA_KEYSPACE", "bench")

	c, err := db.NewCassandra(hosts, keyspace)
	if err != nil {
		fmt.Println("cassandra connect error:", err)
		return
	}
	resultsDir := getEnv("RESULTS_DIR", "results/local")
	res, err := benchmark.RunAndSave(c, ops, resultsDir)
	if err != nil {
		fmt.Println("benchmark error:", err)
		return
	}
	fmt.Printf("Result: %+v\n", res)
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
