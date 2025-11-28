package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/esuEdu/nosql-benchmark-suite/intenal/benchmark"
	"github.com/esuEdu/nosql-benchmark-suite/intenal/db"
)

func main() {
	ops := 100000
	if v := os.Getenv("OPS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			ops = n
		}
	}

	table := getEnv("DYNAMO_TABLE", "BenchTable")
	d, err := db.NewDynamo(table)
	if err != nil {
		fmt.Println("dynamo connect error:", err)
		return
	}
	resultsDir := getEnv("RESULTS_DIR", "results/local")
	res, err := benchmark.RunAndSave(d, ops, resultsDir)
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
