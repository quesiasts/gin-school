# Gin-School

## 1. Installation

Please, install it if you don't have it yet.

Docker: https://docs.docker.com/get-docker/

Docker Compose: https://docs.docker.com/compose/install/

Go to project root and install the dependencies:
```shell
go mod tidy
```

Before starting make sure to stop running containers:
```shell
docker stop $(docker ps -a -q)
```
Now build the application:
```shell
docker-compose up -d --build
```

Once it's done, you can check logs by running:
```shell
docker-compose logs -f
```
Wait until server is up and running. You will see this message:
```
database_1         | LOG:  listening on IPv4 address "0.0.0.0", port 5432
database_1         | LOG:  listening on IPv6 address "::", port 5432
database_1         | LOG:  database system is ready to accept connections
pgadmin-compose_1  | [INFO] Starting gunicorn 20.1.0
pgadmin-compose_1  | [INFO] Listening at: http://[::]:80
```

Now we need to up another shell and running local application:
```shell
go run main.go
```

You can access: http://127.0.0.1:8080/<your-name>

## 2. Tests

Go to project root and run the tests:
```shell
go test -v ./tests
```

## 3. Route List
* /index
* /:name
* /students
* /students/:id
* /students/document/:document