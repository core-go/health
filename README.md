# health
- [HealthChecker](https://github.com/core-go/health/blob/main/health_checker.go)
- [Health](https://github.com/core-go/health/blob/main/health.go) model
- [Check](https://github.com/core-go/health/blob/main/check.go) to build Health model from HealthChecker
###server
- [Serve](https://github.com/core-go/health/blob/main/server/serve.go) to start a server, which is usually used for batch job (for example, message queue consumer)

### http handler
- [http handler](https://github.com/core-go/health/blob/main/health_handler.go)
- [gin](https://github.com/gin-gonic/gin) [handler](https://github.com/core-go/health/blob/main/gin/health_handler.go)
- [echo](https://github.com/labstack/echo) [handler](https://github.com/core-go/health/blob/main/echo/health_handler.go)

## Providers
### Common
- http: http client health checker

### Database  
- redis: support [go-redis/redis](https://github.com/core-go/health/blob/main/redis/health_checker.go) and [garyburd/redigo](https://github.com/core-go/health/blob/main/redigo/health_checker.go)
- [sql](https://github.com/core-go/health/blob/main/sql/health_checker.go)
- [mongo](https://github.com/core-go/health/blob/main/mongo/health_checker.go)
- [dynamodb](https://github.com/core-go/health/blob/main/dynamodb/health_checker.go)
- [firestore](https://github.com/core-go/health/blob/main/firestore/health_checker.go)
- [elasticsearch](https://github.com/core-go/health/blob/main/elasticsearch/health_checker.go) and [elasticsearch/v7](https://github.com/core-go/health/blob/main/elasticsearch/v7/health_checker.go)

### Message Queue
- Amazon Simple Queue Service ([SQS](https://github.com/core-go/health/blob/main/sqs/health_checker.go))
- Google Cloud [Pub/Sub](https://github.com/core-go/health/blob/main/pubsub/health_checker.go)
- Kafka: support [segmentio/kafka-go](https://github.com/core-go/health/blob/main/kafka/health_checker.go) and [Shopify/sarama](https://github.com/core-go/health/blob/main/sarama/health_checker.go)
- [NATS](https://github.com/core-go/health/blob/main/nats/health_checker.go)
- [Active MQ](https://github.com/core-go/health/blob/main/amq/health_checker.go)
- [RabbitMQ](https://github.com/core-go/health/blob/main/rabbitmq/health_checker.go)
- [IBM MQ](https://github.com/core-go/health/blob/main/ibm-mq/health_checker.go)

## Installation
Please make sure to initialize a Go module before installing core-go/health:

```shell
go get -u github.com/core-go/health
```

Import:
```go
import "github.com/core-go/health"
```