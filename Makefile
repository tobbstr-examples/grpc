# Name of Docker image. It may be whatever.
IMAGE=github.com/tobbstr-examples/grpc:v1.0.0

# Build proto-generation docker image.
proto_image:
	DOCKER_BUILDKIT=1 docker build -f build/Dockerfile -t $(IMAGE) .

# Run go generate in docker.
gen: proto_image
	 docker run --rm -w /wd/in -v $(CURDIR):/wd/in $(IMAGE) go generate ./...

run:
	go run cmd/server/server.go
