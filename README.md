# Health Check in Microservices
### Health Check Definition
A health check in microservices is a mechanism that ensures each service is functioning correctly and is available. It typically involves periodically checking the status of various components of a service and reporting their health.

![health](https://cdn-images-1.medium.com/max/800/1*wiWnkgzUoSgJT9QUXfzI8A.png)

### Use Cases of Health Check
#### Service Availability Monitoring:
- <b>Scenario</b>: Ensuring that each microservice is up and running.
- <b>Benefit</b>: Helps in quickly identifying and addressing service outages.
#### Dependency Checking
- <b>Scenario</b>: Verifying that all dependencies of a service are available and functioning.
- <b>Benefit</b>: Ensures the entire application stack is healthy and operational.
#### Deployment Validation
- <b>Scenario</b>: Checking the health of services post-deployment to ensure they are functioning as expected.
- <b>Benefit</b>: Detects deployment issues early, preventing faulty services from affecting the system.
#### Load Balancing
- <b>Scenario</b>: Directing traffic only to healthy instances of a service.
- <b>Benefit</b>: Ensures reliable service delivery by avoiding unhealthy instances.
#### Auto-scaling
- <b>Scenario</b>: Scaling up or down based on the health and load of the services.
- <b>Benefit</b>: Optimizes resource usage and cost efficiency.

### Implementation
#### Tools
- Spring Boot Actuator, AWS Elastic Load Balancer, Kubernetes liveness and readiness probes
#### Endpoints:
- Health check endpoints (e.g., /health, /status) that return the health status of the service.
#### API Design:
- Request: GET /health
- Response:
  ```json
  {
    "status": "DOWN",
    "details": {
      "sql": {
        "status": "DOWN",
        "data": {
          "error": "pq: database 'demo' does not exist"
        }
      },
      "firestore": {
        "status": "UP"
      },
      "kafka": {
        "status": "UP"
      }
    }
  }
  ```
#### Health Status
- UP: Indicates that the application is functioning normally and all health checks have passed.
- DOWN: Indicates that the application is experiencing issues, and one or more health checks have failed.
- OUT_OF_SERVICE: Indicates that the application is temporarily out of service, typically during startup or shutdown.
- UNKNOWN: Indicates that the health status is unknown or could not be determined.

## Implementation of [core-go/health](https://github.com/core-go/health)
#### Core Library
- <b>Purpose</b>: Provides basic health check functionalities
- <b>Features:
  - Define standard health check interfaces.
    - Model [Health](https://github.com/core-go/health/blob/main/health.go)
      ```json
      package health

      type Health struct {
        Status  string                 `json:"status,omitempty"`
        Data    map[string]interface{} `json:"data,omitempty"`
        Details map[string]Health      `json:"details,omitempty"`
      }
      ```
  - Allow custom health checks with this standard interface [Checker](https://github.com/core-go/health/blob/main/checker.go):
    ```json
    package health

    import "context"

    type Checker interface {
      Name() string
      Check(ctx context.Context) (map[string]interface{}, error)
      Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{}
    }
    ```
  - Build the response JSON from many custom health checks by this GO function [Check](https://github.com/core-go/health/blob/main/check.go)
    - This function can be called by http handler ([gin](https://github.com/gin-gonic/gin), [echo](https://github.com/labstack/echo), [mux](https://github.com/gorilla/mux), [go-chi](https://github.com/go-chi/chi)) 
  - Implement basic checks
    - CPU, memory, disk space: not yet implemented.
    - Cache (Redis, Memcached)
    - Databases: [sql](https://github.com/core-go/health/blob/main/sql/health_checker.go), [mongo](https://github.com/core-go/health/blob/main/mongo/health_checker.go), [dynamodb](https://github.com/core-go/health/blob/main/dynamodb/health_checker.go), [firestore](https://github.com/core-go/health/blob/main/firestore/health_checker.go), [elasticsearch](https://github.com/core-go/health/blob/main/elasticsearch/v8/health_checker.go), [cassandra](https://github.com/core-go/health/blob/main/cassandra/health_checker.go), [hive](https://github.com/core-go/health/blob/main/hive/health_checker.go)
    - Message Queue
    - External Service Health
  - Integration with Existing Systems, by supporting these Go libraries: [gin](https://github.com/gin-gonic/gin), [echo](https://github.com/labstack/echo), [mux](https://github.com/gorilla/mux), [go-chi](https://github.com/go-chi/chi)

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
    - Redis: support [go-redis/redis](https://github.com/core-go/health/blob/main/redis/v9/health_checker.go) and [garyburd/redigo](https://github.com/core-go/health/blob/main/redigo/health_checker.go). The sample is at [go-admin](https://github.com/project-samples/go-admin). 
  - Validate cache hit/miss ratio and performance metrics.

#### Database Health Check Library
- <b>Purpose</b>: Monitors the health of database connections
- <b>Features</b>:
  - Check connectivity and response time for various databases (SQL, NoSQL).
    - [sql](https://github.com/core-go/health/blob/main/sql/health_checker.go). The sample is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample).
    - [mongo](https://github.com/core-go/health/blob/main/mongo/health_checker.go). The sample is at [go-mongo-sample](https://github.com/go-tutorials/go-mongo-sample).
    - [dynamodb](https://github.com/core-go/health/blob/main/dynamodb/health_checker.go). The sample is at [go-dynamodb-tutorial](https://github.com/go-tutorials/go-dynamodb-tutorial).
    - [firestore](https://github.com/core-go/health/blob/main/firestore/health_checker.go). The sample is at [go-firestore-sample](https://github.com/go-tutorials/go-firestore-sample).
    - [elasticsearch](https://github.com/core-go/health/blob/main/elasticsearch/v8/health_checker.go). The sample is at [go-elasticsearch-sample](https://github.com/go-tutorials/go-elasticsearch-sample).
    - [cassandra](https://github.com/core-go/health/blob/main/cassandra/health_checker.go). The sample is at [go-cassandra-sample](https://github.com/go-tutorials/go-cassandra-sample).
    - [hive](https://github.com/core-go/health/blob/main/hive/health_checker.go). The sample is at [go-hive-sample](https://github.com/go-tutorials/go-hive-sample).
  - Provide detailed status messages and error handling.

#### Message Queue Health Check Library
- <b>Purpose</b>: Ensures message queues are operational.
- <b>Features</b>:
  - Check connectivity and queue depth for different message brokers.
    - [Amazon SQS](https://github.com/core-go/health/blob/main/sqs/health_checker.go). The sample is at [go-amazon-sqs-sample](https://github.com/project-samples/go-amazon-sqs-sample).
    - [Google Pub/Sub](https://github.com/core-go/health/blob/main/pubsub/health_checker.go). The sample is at [go-pubsub-sample](https://github.com/project-samples/go-pubsub-sample).
    - [Kafka](https://github.com/core-go/health/blob/main/kafka/health_checker.go). The sample is at [go-kafka-sample](https://github.com/project-samples/go-kafka-sample).
    - [NATS](https://github.com/core-go/health/blob/main/nats/health_checker.go). The sample is at [go-nats-sample](https://github.com/project-samples/go-nats-sample).
    - [Active MQ](https://github.com/core-go/health/blob/main/activemq/health_checker.go). The sample is at [go-active-mq-sample](https://github.com/project-samples/go-active-mq-sample).
    - [RabbitMQ](https://github.com/core-go/health/blob/main/rabbitmq/health_checker.go). The sample is at [go-rabbit-mq-sample](https://github.com/project-samples/go-rabbit-mq-sample).
    - [IBM MQ](https://github.com/core-go/health/blob/main/ibmmq/health_checker.go). The sample is at [go-ibm-mq-sample](https://github.com/project-samples/go-ibm-mq-sample).
  - Monitor message lag and processing time (Not yet implemented)

### Future Libraries to develop
#### Cluster Health Check Library
- <b>Purpose</b>: Ensures the health of the microservices cluster.
- <b>Features</b>:
  - Check node status, CPU, and memory usage across the cluster.
  - Integrate with orchestration platforms like Kubernetes for liveness and readiness probes.
#### Metrics and Monitoring Integration Library
- <b>Purpose</b>: Integrates health checks with monitoring tools.
- <b>Features</b>:
  - Export health check results to monitoring systems (Prometheus, Grafana, ELK stack).
  - Provide detailed dashboards and alerting mechanisms.
#### Notification and Alerting Library
- <b>Purpose</b>: Sends alerts based on health check results.
- <b>Features</b>:
  - Integrate with notification systems (Slack, PagerDuty, email).
  - Provide configurable thresholds and alerting rules.

### Integration with Existing Systems
- Designed to integrate seamlessly with existing Go libraries: [Gorilla mux](https://github.com/gorilla/mux), [Go-chi](https://github.com/go-chi/chi), [Echo](https://github.com/labstack/echo) and [Gin](https://github.com/gin-gonic/gin).
  - [handler](https://github.com/core-go/health/blob/main/handler.go), to support [Gorilla mux](https://github.com/gorilla/mux) and [Go-chi](https://github.com/go-chi/chi). The sample is at [go-sql-sample](https://github.com/go-tutorials/go-sql-sample).
  - [echo handler](https://github.com/core-go/health/blob/main/echo/handler.go) to support [Echo](https://github.com/labstack/echo). The sample is at [go-sql-echo-sample](https://github.com/go-tutorials/go-sql-echo-sample).
  - [gin handler](https://github.com/core-go/health/blob/main/gin/handler.go) to support [Gin](https://github.com/gin-gonic/gin). The sample is at  is at [go-sql-gin-sample](https://github.com/go-tutorials/go-sql-gin-sample).

## Installation
Please make sure to initialize a Go module before installing core-go/health:

```shell
go get -u github.com/core-go/health
```

Import:
```go
import "github.com/core-go/health"
```
