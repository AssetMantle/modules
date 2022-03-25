# PersistenceSDK Swagger API Documentation

### Pre-requirements

- get the [swag](https://github.com/swaggo/swag) go binary or install from source

```
go get -u github.com/swaggo/swag/cmd/swag
```

### Generate Swagger

```bash
go mod vendor
cd swagger
swag init  --parseDependency true --parseInternal true --parseVendor true
```

> Note: about cmd will take much longer time depends on system.

- generate `/docs` folder.
- import `/docs` into `main.go` file to use swagger.

```
import _"github.com/persistenceOne/persistenceSDK/swagger/docs"

```

To know more about swaggo use [this](https://github.com/swaggo/swag)

### Start AssetMantle server in unsafe mode
`assetNode start`
`assetClient rest-server --chain-id test --unsafe-cors`

```yml
# BaseRequest

GasPrices     sdk.DecCoins `json:"gas_prices" swaggerignore:"true"`
```