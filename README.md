# simple-todo-gin

https://github.com/pallat/fundamentals-demo

https://github.com/pallat/todoapi

https://github.com/pallat/golang-simple-login-api/blob/master/main.go

https://github.com/pallat/todowasm

https://github.com/pallat/todohexagonal

```shell
go mod init github.com/panutat-p/simeple-todo-gin

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/joho/godotenv

go get -u github.com/go-playground/validator/v10

go get -u github.com/golang-jwt/jwt/v4
```

```shell
docker image build -t simple .

docker container run -p 8080:8080 --name simple simple

docker container run -p:8080:8080 --env-file ./local.env --name simple simple
```

non production use debug
production use release
GIN_MODE=debug
