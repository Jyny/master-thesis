# syntax=docker/dockerfile:experimental

# set builder environment
FROM golang:alpine as builder

RUN mkdir /workdir
WORKDIR /workdir

COPY . .

# build server
RUN cd server && go build -o server ./src/main.go

## build estimat-shift
RUN cd estimat-shift && go build -o estimat-shift ./src/main.go

# set deployment environment
FROM alpine as base

RUN mkdir /workdir
WORKDIR /workdir

RUN apk --no-cache add ca-certificates

# packaging deployment image
FROM base

COPY --from=builder /workdir/server/server /workdir/server
COPY --from=builder /workdir/estimat-shift/estimat-shift /workdir/estimat-shift