FROM golang:1.17.2-alpine

COPY . /chatbox-grpc

WORKDIR /chatbox-grpc/gateway

CMD [ "go", "run", "gateway.go" ]