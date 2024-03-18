FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o proxy ./cmd/main.go

FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/proxy .

EXPOSE 8008 8443

CMD ["./sushi-proxy"]