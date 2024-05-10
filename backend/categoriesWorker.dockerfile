# Builder
FROM golang:alpine AS builder
ENV GOPATH=/build
ARG SERVICE_NAME

# Dependencies
COPY ./pmutils /build/pmutils
WORKDIR /build/pmutils
RUN go mod download

COPY ./protos /build/protos
WORKDIR /build/protos
RUN go mod download

# Compilation
COPY ./$SERVICE_NAME /build/src
WORKDIR /build/src
RUN go mod download
RUN go build -o /build/app /build/src/main.go

# App
FROM alpine
COPY ./bin/grpc_health_probe-linux-amd64 /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe
WORKDIR /application
COPY --from=builder /build/app /application/app
RUN touch ./$SERVICE_NAME.log
CMD ["/application/app"]