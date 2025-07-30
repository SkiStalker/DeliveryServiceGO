FROM golang:alpine AS builder

WORKDIR /usr/src/app

RUN apk update --no-cache
RUN apk add --no-cache git curl unzip protoc protobuf protobuf-dev
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.39/grpc_health_probe-linux-amd64 ./grpc-health-probe
RUN chmod +x ./grpc-health-probe

ADD services/user/model ./model
ADD services/user/repository ./repository
ADD services/user/util ./util
ADD services/user/main.go ./
ADD services/user/user.go ./




ADD proto/user.proto ./
RUN mkdir -p ./grpc_build/user
RUN protoc --go_out=./grpc_build/user --go_opt=paths=source_relative --go-grpc_out=./grpc_build/user --go-grpc_opt=paths=source_relative   user.proto

ADD services/user/go.mod ./
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o user_service .

FROM alpine

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/user_service /usr/src/app/user_service
COPY --from=builder /usr/src/app/grpc-health-probe /usr/src/app/grpc-health-probe

EXPOSE ${USER_SERVICE_PORT}

CMD ["./user_service"]

