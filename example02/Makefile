EXECUTABLE := main
DEPLOY_ACCOUNT := cli
DEPLOY_IMAGE := $(EXECUTABLE)

LDFLAGS ?= -X 'main.Version=$(VERSION)'

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

docker_build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags "$(EXTLDFLAGS)-s -w $(LDFLAGS)" -o $(EXECUTABLE)

docker_image: docker_build
	docker build -t $(DEPLOY_ACCOUNT)/$(DEPLOY_IMAGE) -f Dockerfile .

build_cross:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-linux
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE).exe
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(EXECUTABLE)-darwin
