# Health Check in Microservices

> A lightweight, framework-independent health check framework for Go microservices.

Most health endpoints only answer one question:

> **"Is my HTTP server running?"**

Production systems need a much more important answer:

- Is the database reachable?
- Is Kafka connected?
- Is RabbitMQ available?
- Is Elasticsearch healthy?
- Is the external HTTP service responding?
- Can this service actually perform its job?

[**core-go/health**](https://www.linkedin.com/pulse/microservice-health-check-go-nodejs-duc-nguyen-qunvc) provides a simple, extensible, and production-ready way to monitor the health of your application's critical dependencies.

- You can refer to [Microservice Health Check](https://www.linkedin.com/pulse/microservice-health-check-go-nodejs-duc-nguyen-qunvc) at my [Linked In](https://vn.linkedin.com/in/duc-nguyen-437240239?trk=article-ssr-frontend-pulse_publisher-author-card) for more details.

![health](https://cdn-images-1.medium.com/max/800/1*NreJfea6tHobxMpiq96PPQ.png)


---

# Why?

Many Go applications expose a health endpoint like this:

```go
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
})
```

This only indicates that the HTTP server is alive.

Imagine a microservice that depends on:

- Oracle
- Kafka
- Redis
- Elasticsearch
- External HTTP APIs

If Kafka becomes unavailable:

- the HTTP server is still running
- `/health` still returns **200 OK**
- Kubernetes considers the pod healthy
- users begin experiencing failures
- engineers spend valuable time finding the root cause

With **core-go/health**, the health endpoint immediately reports that **Kafka is DOWN**, allowing your monitoring platform to notify the correct team before users report problems.

---

# Features

- Lightweight
- Framework independent
- Production-ready
- Kubernetes ready
- Built-in providers for common infrastructure
- Per-provider configurable timeout
- Ready-to-use HTTP handlers
- Pluggable provider architecture
- Easy to extend with custom providers
- Zero dependency on any web framework

---

# Supported Providers

Built-in providers include:

### Databases

- SQL Database
- PostgreSQL
- MySQL
- Oracle
- SQL Server
- MongoDB
- Firestore

### Search

- Elasticsearch

### Cache

- Redis

### Message Brokers

- Kafka
- RabbitMQ
- IBM MQ
- ActiveMQ
- NATS

### HTTP

- HTTP Client

### Custom

- Custom Providers

---

# Integration with Existing Systems
Designed to integrate seamlessly with existing Go libraries: [Gorilla mux](https://github.com/gorilla/mux), [Go-chi](https://github.com/go-chi/chi), [Echo](https://github.com/labstack/echo) and [Gin](https://github.com/gin-gonic/gin).

| Framework | Supported |
|-----------|:---------:|
| net/http | ✅ |
| Gin | ✅ |
| Echo | ✅ |

Samples:
  - [handler](https://github.com/core-go/health/blob/main/handler.go), to support [Gorilla mux](https://github.com/gorilla/mux) and [Go-chi](https://github.com/go-chi/chi). The sample is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample).
  - [echo handler](https://github.com/core-go/health/blob/main/echo/handler.go) to support [Echo](https://github.com/labstack/echo). The sample is at [go-sql-echo-sample](https://github.com/go-tutorials/go-sql-echo-sample).
  - [gin handler](https://github.com/core-go/health/blob/main/gin/handler.go) to support [Gin](https://github.com/gin-gonic/gin). The sample is at [go-sql-gin-sample](https://github.com/go-tutorials/go-sql-gin-sample).

### net/http

```go
http.Handle("/health", health.NewHandler(checker))
```

### Gin

```go
router.GET("/health", healthgin.NewHandler(checker))
```

### Echo

```go
e.GET("/health", healthecho.NewHandler(checker))
```

Using another framework?

Simply invoke the health checker and write the response using your preferred framework.

---

# Configurable Timeout

A slow dependency should never block your health endpoint.

Since Kubernetes typically invokes health endpoints every few seconds, every provider supports its own timeout.

```go
health.Add(
    kafka.NewChecker(
        producer,
        4*time.Second,
    ),
)
```

or

```go
health.Add(
    http.NewChecker(
        "User Service",
        "https://user-service/health",
        4*time.Second,
    ),
)
```

Benefits:

- Prevent hanging health requests
- Detect slow or unreachable services
- Keep Kubernetes probes responsive
- Configure different timeout values for different providers

---

# Example

```go
checker := health.New()

checker.Add(sql.New(database))
checker.Add(redis.New(redisClient))
checker.Add(kafka.New(producer))
checker.Add(elasticsearch.New(client))
checker.Add(http.New(
    "User Service",
    "https://user-service/health",
))

http.Handle("/health", health.NewHandler(checker))
```

---

# Example Response

Everything is healthy:

```json
{
    "status": "UP",
    "services": [
        {
            "name": "database",
            "status": "UP"
        },
        {
            "name": "redis",
            "status": "UP"
        },
        {
            "name": "kafka",
            "status": "UP"
        }
    ]
}
```

Kafka is unavailable:

```json
{
    "status": "DOWN",
    "services": [
        {
            "name": "database",
            "status": "UP"
        },
        {
            "name": "redis",
            "status": "UP"
        },
        {
            "name": "kafka",
            "status": "DOWN",
            "error": "connection refused"
        }
    ]
}
```

---

# Architecture

```
                    +-----------------------+
                    |     Health Checker    |
                    +-----------------------+
                               |
        -----------------------------------------------------
        |         |         |         |         |            |
       SQL      Redis     Kafka    RabbitMQ   HTTP   Elasticsearch
        |         |         |         |         |            |
        -----------------------------------------------------
                               |
                        Health Response
```

Every provider implements the same interface.

Adding support for a new system only requires implementing a new provider.

---

# Production Use Cases

Perfect for:

- Kubernetes Readiness Probe
- Kubernetes Liveness Probe
- Docker Health Check
- Amazon ECS
- Google Cloud Run
- Background Workers
- Scheduled Jobs
- REST APIs
- Microservices

---

# Why core-go/health?

Unlike a simple `/health` endpoint that only checks whether the HTTP server is running, **core-go/health** verifies the health of the infrastructure your application actually depends on.

This enables you to:

- Detect dependency failures immediately
- Reduce Mean Time To Detect (MTTD)
- Improve production observability
- Simplify troubleshooting
- Standardize health checks across services
- Reuse the same health framework throughout your organization

---

# Real World Example

A typical microservice may depend on multiple infrastructure components.

```
Order Service
    │
    ├── Oracle
    ├── Kafka
    ├── Redis
    ├── Elasticsearch
    └── User Service (HTTP)
```

When Kafka becomes unavailable:

Without **core-go/health**

```
✓ HTTP Server

/health -> 200 OK
```

Everything appears healthy.

With **core-go/health**

```
✓ Oracle
✓ Redis
✗ Kafka
✓ Elasticsearch
✓ User Service

/health -> DOWN
```

Your monitoring system immediately identifies the failed dependency, allowing engineers to investigate the correct component instead of spending hours searching for the root cause.

---

# Design Principles

- **Framework Independent** — No dependency on Gin, Echo, or any other web framework.
- **Provider Based** — Every dependency is implemented as an independent provider.
- **Extensible** — Easily add custom providers.
- **Production Ready** — Designed for Kubernetes and cloud-native applications.
- **Minimal API** — Simple to learn and easy to integrate.

---

# Philosophy

A healthy service is not simply one whose HTTP server is running.

A healthy service is one that can successfully communicate with all of its critical dependencies and continue serving requests.

---

# Examples:


#### External Service Health Check Library
- <b>Purpose</b>: Monitors the availability of external services.
- <b>Features</b>:
  - Check HTTP/HTTPS endpoints for expected responses.
    - [http client](https://github.com/core-go/health/blob/main/http/health_checker.go). The sample is at [go-sql-hexagonal-architecture-sample](https://github.com/go-tutorials/go-sql-hexagonal-architecture-sample).
  - Measure response time and reliability.

#### Cache Health Check Library
- <b>Purpose</b>: Verifies the status of cache services.
- <b>Features</b>:
  - Check connectivity to cache servers (Redis, Memcached).
    - Redis:
      - [go-redis/redis](https://github.com/core-go/health/blob/main/redis/v9/health_checker.go) to support [redis/go-redis](https://github.com/redis/go-redis). The sample is at [go-admin](https://github.com/project-samples/go-admin).
      - [garyburd/redigo](https://github.com/core-go/health/blob/main/redigo/health_checker.go) to support [gomodule/redigo](https://github.com/gomodule/redigo).
  - Validate cache hit/miss ratio and performance metrics.

#### Database Health Check Libraries
- <b>Purpose</b>: Monitors the health of database connections
- <b>Features</b>:
  - Check connectivity and response time for various databases (SQL, NoSQL).
    - [sql](https://github.com/core-go/health/blob/main/sql/health_checker.go). The sample is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample).
    - [mongo](https://github.com/core-go/health/blob/main/mongo/health_checker.go) to support [mongo-driver/mongo](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo). The sample is at [go-mongo-sample](https://github.com/go-tutorials/go-mongo-sample).
    - [dynamodb](https://github.com/core-go/health/blob/main/dynamodb/health_checker.go) to support [aws/aws-sdk-go/service/dynamodb](https://github.com/aws/aws-sdk-go/tree/main/service/dynamodb). The sample is at [go-dynamodb-tutorial](https://github.com/go-tutorials/go-dynamodb-tutorial).
    - [firestore](https://github.com/core-go/health/blob/main/firestore/health_checker.go) to support [go/firestore](https://pkg.go.dev/cloud.google.com/go/firestore). The sample is at [go-firestore-sample](https://github.com/go-tutorials/go-firestore-sample).
    - [elasticsearch](https://github.com/core-go/health/blob/main/elasticsearch/v8/health_checker.go) to support [elastic/go-elasticsearch](https://github.com/elastic/go-elasticsearch). The sample is at [go-elasticsearch-sample](https://github.com/go-tutorials/go-elasticsearch-sample).
    - [cassandra](https://github.com/core-go/health/blob/main/cassandra/health_checker.go) to support [apache/cassandra-gocql-driver](https://github.com/apache/cassandra-gocql-driver). The sample is at [go-cassandra-sample](https://github.com/go-tutorials/go-cassandra-sample).
    - [hive](https://github.com/core-go/health/blob/main/hive/health_checker.go) to support [beltran/gohive](https://github.com/beltran/gohive). The sample is at [go-hive-sample](https://github.com/go-tutorials/go-hive-sample).
  - Provide detailed status messages and error handling.

#### Message Queue Health Check Libraries
- <b>Purpose</b>: Ensures message queues are operational.
- <b>Features</b>:
  - Check connectivity and queue depth for different message brokers.
    - [Amazon SQS](https://github.com/core-go/health/blob/main/sqs/health_checker.go): support [aws-sdk-go/service/sqs](https://github.com/aws/aws-sdk-go/tree/main/service/sqs). The sample is at [go-amazon-sqs-sample](https://github.com/project-samples/go-amazon-sqs-sample).
    - [Google Pub/Sub](https://github.com/core-go/health/blob/main/pubsub/health_checker.go): support [go/pubsub](https://pkg.go.dev/cloud.google.com/go/pubsub). The sample is at [go-pubsub-sample](https://github.com/project-samples/go-pubsub-sample).
    - [Kafka](https://github.com/core-go/health/blob/main/kafka/health_checker.go): support 3 GO libraries: [segmentio/kafka-go](https://github.com/segmentio/kafka-go), [IBM/sarama](https://github.com/IBM/sarama) and [confluent](https://github.com/confluentinc/confluent-kafka-go). The sample is at [go-kafka-sample](https://github.com/project-samples/go-kafka-sample).
    - [NATS](https://github.com/core-go/health/blob/main/nats/health_checker.go): support [nats.go](https://github.com/nats-io/nats.go). The sample is at [go-nats-sample](https://github.com/project-samples/go-nats-sample).
    - [Active MQ](https://github.com/core-go/health/blob/main/activemq/health_checker.go): support [go-stomp](https://github.com/go-stomp/stomp). The sample is at [go-active-mq-sample](https://github.com/project-samples/go-active-mq-sample).
    - [RabbitMQ](https://github.com/core-go/health/blob/main/rabbitmq/health_checker.go): support [rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go). The sample is at [go-rabbit-mq-sample](https://github.com/project-samples/go-rabbit-mq-sample).
    - [IBM MQ](https://github.com/core-go/health/blob/main/ibmmq/health_checker.go): support [ibmmq](https://github.com/ibm-messaging/mq-golang). The sample is at [go-ibm-mq-sample](https://github.com/project-samples/go-ibm-mq-sample).

---

# License

MIT
