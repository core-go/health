# Health
- HealthChecker (no dependency)
- Health model
- Build Health model from HealthChecker 
- HttpHealthChecker
- SqlHealthChecker (v1.0.4)
- HealthHandler for http response

## Installation

Please make sure to initialize a Go module before installing common-go/health:

```shell
go get -u github.com/common-go/health
```

Import:

```go
import "github.com/common-go/health"
```

You can optimize the import by version:
- v0.0.1: HealthService
- v0.0.2: Health Model
- v0.0.3: Check
- v0.0.8: Http Health Service
- v1.1.1: HealthHandler for [gin](https://github.com/gin-gonic/gin) 
- v1.1.3: HealthHandler for [echo v3](https://github.com/labstack/echo)
- v1.1.4: HealthHandler for [echo v4](https://github.com/labstack/echo)
- **v1.0.3: HealthHandler**
- **v1.0.4: HealthHandler and SqlHealthChecker**
- **v1.0.6: HealthHandler, SqlHealthChecker and Serve(to start http server)**