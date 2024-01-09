# Hugeman

## 1. pre-commit install
```bash
pre-commit install --install-hooks
```

## 2. dependency check
```bash
brew install dependency-check
```

## 3. config dependency check in cmd/
```bash
echo '<?xml version="1.0" encoding="UTF-8"?>\n<suppressions xmlns="https://jeremylong.github.io/DependencyCheck/dependency-suppression.1.3.xsd">\n</suppressions>' > dependency-suppressions.xml
```

## 4. 1st scan
```bash
dependency-check --out cmd/ --scan cmd/ --project $(basename $PWD)/cmd --enableExperimental --failOnCVSS 0 --suppression cmd/dependency-suppressions.xml
```

## 5. run
```bash
docker-compose up -d && docker-compose logs --follow
```

## 6. gen mock (optional)
```bash
make -f Makefile gen-mock-all
```

## 7. test
```bash
go test -cover ./...

?   	hugeman/cmd/api	[no test files]
?   	hugeman/configs	[no test files]
?   	hugeman/docs	[no test files]
?   	hugeman/internal/core/domain	[no test files]
?   	hugeman/internal/core/ports	[no test files]
?   	hugeman/internal/repositories	[no test files]
?   	hugeman/pkg/database_driver/gorm	[no test files]
?   	hugeman/pkg/testings/mock_repository	[no test files]
?   	hugeman/pkg/testings/mock_service	[no test files]
?   	hugeman/pkg/validator	[no test files]
?   	hugeman/protocal	[no test files]
ok  	hugeman/internal/handlers	0.157s	coverage: 83.8% of statements
ok  	hugeman/internal/services	0.496s	coverage: 82.8% of statements
```

## 8. layout
```bash
├── build
├── cmd
│   └── api
│       └── main.go
├── configs
│   └── config.go
├── docs
│   └── docs.go
├── internal
│   ├── core
│   │   ├── domain
│   │   │   ├── gorm.go
│   │   │   ├── hook.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   └── ports
│   │       ├── repository.go
│   │       └── service.go
│   ├── handlers
│   │   ├── handler.go
│   │   └── handler_test.go
│   ├── repositories
│   │   └── postgres.go
│   └── services
│       ├── service.go
│       └── service_test.go
├── pkg
│   ├── database_driver
│   │   └── gorm
│   │       └── postgres.go
│   ├── testings
│   │   ├── mock_repository
│   │   │   └── mockrepository.go
│   │   └── mock_service
│   │       └── mockservice.go
│   └── validator
│       └── validator.go
├── protocal
│   └── http.go
└── tmp
    └── app
```

### 9. Swagger

http://localhost:9089/swagger/index.html

```bash
swag init -g cmd/api/main.go --output docs
```