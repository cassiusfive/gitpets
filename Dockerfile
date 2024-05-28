ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo -o /run-app ./cmd/gitpets

FROM scratch

COPY --from=builder /run-app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY assets ./assets
CMD ["./run-app"]
