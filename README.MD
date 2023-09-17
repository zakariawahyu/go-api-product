# Api Product Services ✨

## Technologies - Libraries
* [labstack/echo](https://go.dev/dl/) - as http framework
* [go-gorm/gorm](https://github.com/go-gorm/gorm) - as ORM
* [spf13/viper](https://github.com/spf13/viper) - as config configuration
* [uber-go/zap](https://github.com/uber-go/zap) - as logger
* [go-playground/validator](https://github.com/go-playground/validator) - as validator
* [swaggo/swag](https://github.com/swaggo/swag) - as api documentation

## Run with docker
Run docker-compose.yaml
```
$ make docker-compose
```
or
```
$ docker docker-compose up -d --build
```
- After run this, don't forget to check all container are successfully started
- There are 2 container, api_service and database

## Run locally
1. Create database container
```
$ make local

```
&ensp; or
```
docker-compose -f docker-compose.local.yaml up -d
```
2. Run Service
```
$ make run
```
&ensp; or
```
$ go run cmd/main.go
```
3. If use other container database, please set config file in config/config.yaml
```
Mysql:
  Host: localhost
  Port: 3306
  User: root
  Password: masukdb
  DbName: api_product
  MaxIdleConnection: 10
  MaxOpenConnection: 100
```
## Swagger Docs
```
http://localhost:7070/swagger/index.html
```
