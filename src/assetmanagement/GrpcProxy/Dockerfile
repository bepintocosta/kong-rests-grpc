FROM golang:1.20 AS base

RUN apt update
RUN apt upgrade -y
RUN apt install unzip 
RUN apt install -y protobuf-compiler


WORKDIR /app
ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GIT_TERMINAL_PROMPT=1
ENV GRPC_HOST=:5000


COPY /src/assetmanagement/GrpcProxy/tools/tools.go /app/tools/

WORKDIR /app/tools

RUN go mod init tools
RUN go mod tidy
RUN go generate tools.go


#Dependecies
WORKDIR /app

COPY src/assetmanagement/GrpcProxy/go.mod ./ 
COPY src/assetmanagement/GrpcProxy/go.sum ./
#Proto&HttpRules
COPY Protos /app/protos

COPY src/assetmanagement/GrpcProxy/httpRules.yaml ./

RUN mkdir -p ./gen

COPY src/assetmanagement/GrpcProxy/makefile /app

COPY src/assetmanagement/GrpcProxy/main.go /app

RUN make generate

RUN make client


EXPOSE 8090
ENTRYPOINT "/app/grpc_gateway"