# Health
![health](https://camo.githubusercontent.com/49287a63a0e1c52818c4321650b3f8cf2348d5f50108aed820cd6441fbb2574d/68747470733a2f2f63646e2d696d616765732d312e6d656469756d2e636f6d2f6d61782f3830302f312a6746457a416b7674666e51575665463265644b7767512e706e67)
- [Checker](https://github.com/core-go/health/blob/main/checker.go)
- [Health](https://github.com/core-go/health/blob/main/health.go) model
- [Check](https://github.com/core-go/health/blob/main/check.go) to build Health model from HealthChecker

### Server
- [Serve](https://github.com/core-go/health/blob/main/server/serve.go) to start a server, which is usually used for batch job (for example, message queue consumer)

### HTTP handler
- [handler](https://github.com/core-go/health/blob/main/handler.go)
- [gin](https://github.com/gin-gonic/gin) [handler](https://github.com/core-go/health/blob/main/gin/handler.go)
- [echo v4 handler](https://github.com/core-go/health/blob/main/echo/handler.go) and [echo v3 handler](https://github.com/core-go/health/blob/main/echo_v3/handler.go)

## Providers
### Common
- [http client](https://github.com/core-go/health/blob/main/http/health_checker.go) health checker

### Database  
- redis: support [go-redis/redis](https://github.com/core-go/health/blob/main/redis/health_checker.go), [go-redis/redis v8](https://github.com/core-go/health/blob/main/redis/v8/health_checker.go), and [garyburd/redigo](https://github.com/core-go/health/blob/main/redigo/health_checker.go)
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
