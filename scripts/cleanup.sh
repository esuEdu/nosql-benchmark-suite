#!/usr/bin/env bash
set -e

echo "=== Cleaning MongoDB ==="
docker exec mongo mongosh --eval 'db.getSiblingDB("benchdb").dropDatabase()'

echo "=== Cleaning Redis ==="
docker exec redis redis-cli FLUSHALL

echo "=== Cleaning Cassandra ==="
docker exec cassandra cqlsh -e "DROP KEYSPACE IF EXISTS bench;"

echo "=== Cleaning DynamoDB (Local) ==="
rm -rf ~/.dynamodb

echo "=== Cleanup completed ==="
