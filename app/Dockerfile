FROM golang:1.21-alpine as builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin ./...

FROM alpine:3.12

COPY --from=builder /usr/local/bin/teste /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]

