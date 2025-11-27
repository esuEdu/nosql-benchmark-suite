#!/usr/bin/env bash
set -e

export RESULTS_DIR="results/ec2"
mkdir -p $RESULTS_DIR

export OPS=${OPS:-5000}

echo "=== Running Mongo benchmark ==="
go run cmd/mongo/main.go

echo "=== Running Redis benchmark ==="
go run cmd/redis/main.go

echo "=== Running DynamoDB benchmark ==="
export DYNAMO_TABLE="BenchTable"
go run cmd/dynamo/main.go || echo "DynamoDB skipped"

echo "=== Running Cassandra benchmark ==="
export CASSANDRA_HOSTS="127.0.0.1"
export CASSANDRA_KEYSPACE="bench"
go run cmd/cassandra/main.go

echo "=== Benchmarks finalizados ==="
