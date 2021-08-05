# PersistenceSDK Swagger API Documentation

### Pre-requirements
- get the [swag](https://github.com/swaggo/swag) go binary or install from source 
```go
go get -u github.com/swaggo/swag/cmd/swag
```

### Generate Swagger 
```bash
cd swagger
swag init  --parseDependency true --parseInternal true --parseVendor true
```
>Note: about cmd will take much longer time depends on system.
- generate `/docs` folder.
- import `/docs` into `main.go` file to use swagger.
```go
    _ "github.com/persistenceOne/persistenceSDK/swagger/docs"

```

To know more about swaggo use [this](https://github.com/swaggo/swag)