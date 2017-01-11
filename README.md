# echo-server

## Run
```
go run server.go
```

## Docker
Build:

```
docker build -t my-golang-app .
```

Run:

```
docker run --publish 6060:8080 --name my-running-app my-golang-app
```
