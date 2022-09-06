# syntax = docker/dockerfile:1

FROM golang:1.18-alpine3.15 AS base

COPY go.mod go.sum /go/src/backend/
WORKDIR /go/src/backend
RUN go mod download

RUN --mount=type=cache,target=/var/.cache/apk \
    apk update && apk add --no-cache bash build-base

COPY . /go/src/backend/

FROM base as builder
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build

FROM base as tests
COPY --from=builder /go/src/backend/build /usr/bin/
COPY database/migrations/sources/* /database/migrations/sources/
COPY --from=builder /go/src/backend/.env .
ENTRYPOINT ["/usr/bin/build"]

FROM builder as local
COPY --from=builder /go/src/backend/build /usr/bin/
COPY database/migrations/sources/* /database/migrations/sources/
COPY --from=builder /go/src/backend/.env .
ENTRYPOINT ["/usr/bin/build"]

FROM alpine:3.15 as stage

RUN apk update && apk add --no-cache bash npm
RUN npm install -g graphql-schema-utilities
COPY --from=builder /go/src/backend/build /usr/bin/
COPY database/migrations/sources/* /go/database/migrations/sources/
# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR $GOPATH

COPY --from=builder /go/src/backend/.env .

ENTRYPOINT ["/usr/bin/build"]
