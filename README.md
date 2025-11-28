# NoSQL Benchmark Suite (Go)

Um conjunto de benchmarks escritos em **Go** para comparar a performance
dos principais bancos NoSQL:

-   **MongoDB** (Document Store)
-   **Redis** (Key-Value in-memory)
-   **DynamoDB** (Key-Value distribuído cloud-native)
-   **Cassandra** (Wide-Column Store)

Os testes medem:

-   **Latência de escrita**
-   **Latência de leitura**
-   **Tempo total de execução**
-   **Throughput por segundo (ops/s)**

Tudo de forma **simples**, **reproduzível** e **100% gratuita**, ideal
para artigos, estudos, testes práticos e pesquisa acadêmica.

## 📁 Estrutura do Repositório

    nosql-benchmark-suite/
    │
    ├── cmd/
    │   ├── mongo/
    │   │   └── main.go
    │   ├── redis/
    │   │   └── main.go
    │   ├── dynamo/
    │   │   └── main.go
    │   ├── cassandra/
    │   │   └── main.go
    │   └── runner/
    │       └── main.go
    │
    ├── internal/
    │   ├── db/
    │   │   ├── mongo.go
    │   │   ├── redis.go
    │   │   ├── dynamo.go
    │   │   ├── cassandra.go
    │   │   └── interfaces.go
    │   └── benchmark/
    │       └── runner.go
    │
    ├── configs/
    │   ├── mongo.yaml
    │   ├── redis.yaml
    │   ├── dynamo.yaml
    │   └── cassandra.yaml
    │
    ├── scripts/
    │   ├── setup_ec2.sh
    │   ├── run_all.sh
    │   └── docker-compose.yaml
    │
    ├── results/
    │   ├── local/
    │   ├── ec2/
    │   └── charts/
    │
    ├── docs/
    │   ├── article.md
    │   ├── architecture.md
    │   └── results.md
    │
    ├── go.mod
    ├── go.sum
    └── README.md

## 🚀 Como Rodar os Benchmarks

### 1. Pré-requisitos

-   Go 1.21+
-   Docker (opcional)

### Rodando os testes individualmente

MongoDB:

```bash
go run cmd/mongo/main.go
```

Redis:

```bash
go run cmd/redis/main.go
```

DynamoDB:

```bash
go run cmd/dynamo/main.go
```

Cassandra:

```bash
go run cmd/cassandra/main.go
```

## 📝 Licença

MIT License
