# grpc

An example of how to generate a gRPC server and client from Protobuf files, and how to run the server.

## Prerequisites

- A docker runtime such as ex. Docker Desktop
- Go installed

## Usage

- To run the gRPC server enter `make run`
- To generate Go code from the Protobuf file in `./api/grpc/v1` run `make gen`.

## What's in this repo?

`./api/grpc/v1`

Contains the protobuf file that defines the gRPC server and its request and response
models.

`./build`

Contains the Dockerfile that specifies an image that can generate Go code from the protobuf
file.

`./cmd/server`

Contains the gRPC server implementation.

`./internal/pb/v1`

Contains the generated protobuf Go code.

`./pkg/client`

Contains the client that consumers of this gRPC server should use to connect to the server
and make requests.

`./buf.work.yaml` and `./buf.gen.yaml`

Files needed by [Buf](https://buf.build/) to generate Go code from the protobuf file.
