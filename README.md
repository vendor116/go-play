# play-go

```shell
curl -X GET \
  "localhost:8080/v1/info"
```

## Гексагональная архитектура
```
user-service/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── migration/
│       └── main.go
├── internal/
│   ├── core/
│   │   ├── domain/
│   │   ├── ports/
│   │   └── services/
│   ├── application/
│   │   ├── usecases/
│   │   └── dto/
│   └── adapters/
│       ├── primary/
│       │   ├── http/
│       │   ├── grpc/
│       │   └── cli/
│       ├── secondary/
│       │   ├── persistence/
│       │   ├── cache/
│       │   ├── messagebroker/
│       │   └── external/
│       └── common/
│           ├── logger/
│           ├── telemetry/
│           └── health/
├── pkg/
│   ├── config/
│   ├── database/
│   └── utils/
├── api/
│   ├── proto/
│   └── openapi/
├── deployments/
│   ├── docker/
│   └── k8s/
├── scripts/
└── Makefile
```
