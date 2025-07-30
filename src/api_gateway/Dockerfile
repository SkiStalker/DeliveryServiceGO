FROM golang:alpine AS builder

WORKDIR /usr/src/app

RUN apk update --no-cache
RUN apk add --no-cache git curl unzip protoc protobuf protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ADD api_gateway/clients ./clients
ADD api_gateway/handlers ./handlers
ADD api_gateway/model ./model
ADD api_gateway/router ./router
ADD api_gateway/util ./util
ADD api_gateway/main.go ./
ADD api_gateway/go.mod ./

ADD proto/user.proto ./
RUN mkdir -p ./grpc_build/user
RUN protoc --go_out=./grpc_build/user --go_opt=paths=source_relative --go-grpc_out=./grpc_build/user --go-grpc_opt=paths=source_relative   user.proto



RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api-gateway .

EXPOSE 8080

FROM alpine

WORKDIR /usr/src/app/

RUN apk update --no-cache

RUN apk add --no-cache curl

COPY --from=builder /usr/src/app/api-gateway /usr/src/app/api-gateway

CMD ["/usr/src/app/api-gateway"]