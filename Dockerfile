ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
WORKDIR /usr/src/app/cmd/gitpets
RUN go build -v -o /run-app .
WORKDIR /usr/src/app


FROM debian:bookworm
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/assets /usr/local/bin/assets

WORKDIR /usr/local/bin
CMD ["run-app"]
