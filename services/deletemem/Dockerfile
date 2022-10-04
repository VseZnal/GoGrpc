FROM golang:1.12.6

WORKDIR /app

RUN mkdir -p /app/src
COPY . /app

RUN apt-get update
RUN apt install -y protobuf-compiler

RUN GO111MODULE=on \
        go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

RUN go mod download
CMD go run main.go

EXPOSE 8080