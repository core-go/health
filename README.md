# Health
- HealthService (no dependency)
- Health model
- Build Health model from HealthService 
- HttpHealthService
- HealthController for http response

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
- v1.1.1: HealthController for [gin](https://github.com/gin-gonic/gin) 
- v1.1.3: HealthController for [echo v3](https://github.com/labstack/echo)
- v1.1.4: HealthController for [echo v4](https://github.com/labstack/echo)
- **v1.0.0: HealthController** (latest version)