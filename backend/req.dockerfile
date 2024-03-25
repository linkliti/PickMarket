# Builder
FROM golang:alpine AS builder
ENV GOPATH=/build

# Dependencies
COPY ./pmutils /build/pmutils
WORKDIR /build/pmutils
RUN go mod download

COPY ./protos /build/protos
WORKDIR /build/protos
RUN go mod download

# Compilation
COPY ./requestHandler /build/src
WORKDIR /build/src
RUN go mod download
RUN go build -o /build/app /build/src/main.go

# App
FROM alpine
WORKDIR /application
COPY --from=builder /build/app /application/app
RUN touch ./requestHandler.log
CMD ["/application/app"]