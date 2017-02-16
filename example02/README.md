# Write command line tool using urfave/cli

## installation

```sh
go get -u -d github.com/urfave/cli
```

## build command

```
go build -o main 
```

## build docker image

```
make docker_image
```

Run

```
docker run -e DOCKER_USERNAME=appleboy cli/main
```

## Cross Compile in Golang

```
make build_cross
```
